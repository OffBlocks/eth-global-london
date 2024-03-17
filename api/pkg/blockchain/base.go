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

type BaseBlockchain struct {
	client                *ethclient.Client
	currencies            map[string]Token
	transferer            *contracts.ProgrammableTokenTransfer
	ethereumReceiver      common.Address
	ethereumChainSelector uint64
	adminAddress          common.Address
	key                   *ecdsa.PrivateKey
	Gateway               common.Address
}

func NewBaseBlockchain(c *config.Config) (*BaseBlockchain, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, c.BaseEndpointUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum endpoint: %w", err)
	}

	currencies := map[string]Token{
		"USD": {common.HexToAddress(c.BaseUsdAddress), c.BaseUsdDecimals},
	}

	gateway := common.HexToAddress(c.BaseTransferAddress)

	transferer, err := contracts.NewProgrammableTokenTransfer(common.HexToAddress(c.BaseTransferAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create transferer contract: %w", err)
	}

	key, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &BaseBlockchain{
		client:                client,
		currencies:            currencies,
		transferer:            transferer,
		key:                   key,
		ethereumReceiver:      common.HexToAddress(c.EthereumTransferAddress),
		adminAddress:          common.HexToAddress(c.AdminAddress),
		ethereumChainSelector: c.EthereumChainSelector,
		Gateway:               gateway,
	}, nil
}

func (b *BaseBlockchain) TransferToken(ctx context.Context, from common.Address, amount types.Decimal, currency string) (common.Hash, error) {
	chainId, err := b.client.ChainID(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get chain ID: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(b.key, chainId)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to create transactor: %w", err)
	}

	nonce, err := b.client.PendingNonceAt(ctx, b.adminAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)               // in wei
	auth.GasLimit = uint64(0)                // in units
	auth.GasTipCap = big.NewInt(2000000000)  // maxPriorityFeePerGas = 2 Gwei
	auth.GasFeeCap = big.NewInt(20000000000) // maxFeePerGas = 20 Gwei
	auth.Context = ctx

	token, ok := b.currencies[currency]
	if !ok {
		return common.Hash{}, fmt.Errorf("currency not supported")
	}

	inputAmount := amount.Mul(decimal.New(10, int32(token.Decimals-1))).Truncate(0)

	tx, err := b.transferer.TransferToken(
		auth,
		uint64(b.ethereumChainSelector),
		b.ethereumReceiver,
		contracts.ProgrammableTokenTransferTransferData{
			Payer:       from,
			Beneficiary: b.adminAddress,
		},
		token.Address,
		inputAmount.BigInt(),
	)

	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to transfer token: %w", err)
	}

	return tx.Hash(), nil
}
