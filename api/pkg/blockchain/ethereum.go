package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/offblocks/eth-global-london/api/pkg/blockchain/contracts"
	"github.com/offblocks/eth-global-london/api/pkg/config"
	"github.com/offblocks/offblocks-common/types"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumBlockchain struct {
	client       *ethclient.Client
	currencies   map[string]Token
	quoter       *contracts.IQuoterV2
	transferer   *contracts.ProgrammableTokenTransfer
	swapRouter   *contracts.ISwapRouter02
	adminAddress common.Address
	key          *ecdsa.PrivateKey
}

func NewEthereumBlockchain(c *config.Config) (*EthereumBlockchain, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, c.EthereumEndpointUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum endpoint: %w", err)
	}

	currencies := map[string]Token{
		"USD": {common.HexToAddress(c.EthereumUsdAddress), c.EthereumUsdDecimals},
		"EUR": {common.HexToAddress(c.EthereumEurAddress), c.EthereumEurDecimals},
	}

	quoter, err := contracts.NewIQuoterV2(common.HexToAddress(c.EthereumQuoterAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create quoter contract: %w", err)
	}

	transferer, err := contracts.NewProgrammableTokenTransfer(common.HexToAddress(c.EthereumTransferAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create transferer contract: %w", err)
	}

	swapRouter, err := contracts.NewISwapRouter02(common.HexToAddress(c.EthereumSwapRouterAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create swap router contract: %w", err)
	}

	key, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &EthereumBlockchain{
		client:       client,
		currencies:   currencies,
		quoter:       quoter,
		swapRouter:   swapRouter,
		transferer:   transferer,
		adminAddress: common.HexToAddress(c.AdminAddress),
		key:          key,
	}, nil
}

func (e *EthereumBlockchain) GetQuote(ctx context.Context, source string, destination string, output types.Decimal) (*types.Decimal, error) {
	sourceToken, ok := e.currencies[source]
	if !ok {
		return nil, fmt.Errorf("source currency not supported")
	}

	destinationToken, ok := e.currencies[destination]
	if !ok {
		return nil, fmt.Errorf("destination currency not supported")
	}

	outputAmount := output.Mul(decimal.New(10, int32(destinationToken.Decimals-1))).Round(0)

	out, err := e.quoter.QuoteExactOutputSingleStatic(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, contracts.IQuoterV2QuoteExactOutputSingleParams{
		TokenIn:           sourceToken.Address,
		TokenOut:          destinationToken.Address,
		Amount:            outputAmount.BigInt(),
		Fee:               big.NewInt(500),
		SqrtPriceLimitX96: big.NewInt(0),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %w", err)
	}

	// Add 1% to the input amount to account for slippage and FX rate changes
	input := decimal.NewFromBigInt(out.AmountIn, -int32(sourceToken.Decimals)).Mul(decimal.NewFromFloat(1.01)).Truncate(int32(sourceToken.Decimals))

	return &types.Decimal{Decimal: input}, nil
}

func (e *EthereumBlockchain) SwapToken(ctx context.Context, to common.Address, maxInputAmount types.Decimal, inputCurrency string, outputAmount types.Decimal, outputCurrency string) (common.Hash, error) {
	chainId, err := e.client.ChainID(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get chain ID: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(e.key, chainId)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to create transactor: %w", err)
	}

	nonce, err := e.client.PendingNonceAt(ctx, e.adminAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)               // in wei
	auth.GasLimit = uint64(0)                // in units
	auth.GasTipCap = big.NewInt(2000000000)  // maxPriorityFeePerGas = 2 Gwei
	auth.GasFeeCap = big.NewInt(20000000000) // maxFeePerGas = 20 Gwei
	auth.Context = ctx

	inputToken, ok := e.currencies[inputCurrency]
	if !ok {
		return common.Hash{}, fmt.Errorf("input currency not supported")
	}

	outputToken, ok := e.currencies[outputCurrency]
	if !ok {
		return common.Hash{}, fmt.Errorf("output currency not supported")
	}

	maxInput := maxInputAmount.Mul(decimal.New(10, int32(inputToken.Decimals-1))).Round(0)
	output := outputAmount.Mul(decimal.New(10, int32(outputToken.Decimals-1))).Round(0)

	tx, err := e.swapRouter.ExactOutputSingle(auth, contracts.IV3SwapRouterExactOutputSingleParams{
		TokenIn:           inputToken.Address,
		TokenOut:          outputToken.Address,
		Fee:               big.NewInt(500),
		Recipient:         to,
		AmountOut:         output.BigInt(),
		AmountInMaximum:   maxInput.BigInt(),
		SqrtPriceLimitX96: big.NewInt(0),
	})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to swap token: %w", err)
	}

	return tx.Hash(), nil
}
