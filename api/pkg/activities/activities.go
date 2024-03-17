package activities

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offblocks/eth-global-london/api/pkg/blockchain"
	"github.com/offblocks/eth-global-london/api/pkg/external"
	"github.com/offblocks/eth-global-london/api/pkg/model"
	"github.com/offblocks/eth-global-london/api/pkg/repository"
	"github.com/offblocks/offblocks-common/types"
	"go.temporal.io/sdk/activity"
)

type Activities struct {
	Repository         *repository.Repository
	BaseBlockchain     *blockchain.BaseBlockchain
	EthereumBlockchain *blockchain.EthereumBlockchain
}

type TransferTokensInput struct {
	FromWalletId string
	Amount       types.Decimal
	Currency     string
}

type TransferTokensResult struct {
	TransactionId string
}

type SwapTokensInput struct {
	Beneficiary    string
	MaxInputAmount types.Decimal
	InputCurrency  string
	OutputAmount   types.Decimal
	OutputCurrency string
}

type SwapTokensResult struct {
	TransactionId string
}

func (s *Activities) SavePayment(ctx context.Context, input *model.Payment) error {
	return s.Repository.PaymentRepository.Save(ctx, input)
}

func (s *Activities) FetchAccount(ctx context.Context, input *types.UUID) (*model.Account, error) {
	return s.Repository.AccountRepository.GetById(ctx, *input)
}

func (s *Activities) TransferTokens(ctx context.Context, input *TransferTokensInput) (*TransferTokensResult, error) {
	transactionId, err := s.BaseBlockchain.TransferToken(ctx, common.HexToAddress(input.FromWalletId), input.Amount, input.Currency)
	if err != nil {
		return nil, err
	}

	return &TransferTokensResult{
		TransactionId: transactionId.Hex(),
	}, nil
}

func (s *Activities) IsTokenTransferFinished(ctx context.Context, input *string) error {
	for {
		time.Sleep(30 * time.Second)

		activity.RecordHeartbeat(ctx)

		resp, err := external.FetchCCIPMessage(*input)
		if err != nil {
			continue
		}

		if len(resp.Data.TransactionHash.Nodes) == 0 || resp.Data.TransactionHash.Nodes[0].DestTransactionHash == nil {
			continue
		} else {
			break
		}
	}

	return nil
}

func (s *Activities) SwapTokens(ctx context.Context, input *SwapTokensInput) (*SwapTokensResult, error) {
	transactionId, err := s.EthereumBlockchain.SwapToken(ctx, common.HexToAddress(input.Beneficiary), input.MaxInputAmount, input.InputCurrency, input.OutputAmount, input.OutputCurrency)
	if err != nil {
		return nil, err
	}

	return &SwapTokensResult{
		TransactionId: transactionId.Hex(),
	}, nil
}
