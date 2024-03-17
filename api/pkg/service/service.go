package service

import (
	"context"

	"github.com/offblocks/eth-global-london/api/pkg/api"
	"github.com/offblocks/eth-global-london/api/pkg/blockchain"
	"github.com/offblocks/eth-global-london/api/pkg/model"
	"github.com/offblocks/eth-global-london/api/pkg/repository"
	"github.com/offblocks/eth-global-london/api/pkg/workflows"
	"github.com/offblocks/offblocks-common/types"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

const (
	TaskQueue = "pay-api"
)

type Service struct {
	Repository         repository.Repository
	BaseBlockchain     *blockchain.BaseBlockchain
	EthereumBlockchain *blockchain.EthereumBlockchain
	TemporalClient     client.Client
}

func NewApiService(repository repository.Repository, base *blockchain.BaseBlockchain, ethereum *blockchain.EthereumBlockchain, temporal client.Client) *Service {
	return &Service{repository, base, ethereum, temporal}
}

// Retrieve accounts
// (GET /accounts)
func (s *Service) GetAccounts(ctx context.Context, request api.GetAccountsRequestObject) (api.GetAccountsResponseObject, error) {
	accounts, err := s.Repository.AccountRepository.GetAll(ctx)
	if err != nil {
		return api.GetAccounts500Response{}, err
	}

	apiAccounts := make([]api.Account, len(accounts))
	for i, account := range accounts {
		apiAccounts[i] = api.Account{
			Id:             account.Id,
			Name:           account.Name,
			Identifier:     account.Identifier,
			IdentifierType: account.IdentifierType,
			WalletId:       account.WalletId,
			CreatedAt:      types.Time{Time: account.CreatedAt},
			UpdatedAt:      types.Time{Time: account.UpdatedAt},
		}
	}

	return api.GetAccounts200JSONResponse(apiAccounts), nil
}

// Create new account
// (POST /accounts)
func (s *Service) CreateAccount(ctx context.Context, request api.CreateAccountRequestObject) (api.CreateAccountResponseObject, error) {
	account := &model.Account{
		Name:           request.Body.Name,
		Identifier:     request.Body.Identifier,
		IdentifierType: request.Body.IdentifierType,
		WalletId:       request.Body.WalletId,
	}

	err := s.Repository.AccountRepository.Save(ctx, account)
	if err != nil {
		return api.CreateAccount500Response{}, err
	}

	return api.CreateAccount201JSONResponse(api.Account{
		Id:             account.Id,
		Name:           account.Name,
		Identifier:     account.Identifier,
		IdentifierType: account.IdentifierType,
		WalletId:       account.WalletId,
		CreatedAt:      types.Time{Time: account.CreatedAt},
		UpdatedAt:      types.Time{Time: account.UpdatedAt},
	}), nil
}

// Retrieve account
// (GET /accounts/{accountId})
func (s *Service) GetAccount(ctx context.Context, request api.GetAccountRequestObject) (api.GetAccountResponseObject, error) {
	account, err := s.Repository.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return api.GetAccount500Response{}, err
	}

	return api.GetAccount200JSONResponse(api.Account{
		Id:             account.Id,
		Name:           account.Name,
		Identifier:     account.Identifier,
		IdentifierType: account.IdentifierType,
		WalletId:       account.WalletId,
		CreatedAt:      types.Time{Time: account.CreatedAt},
		UpdatedAt:      types.Time{Time: account.UpdatedAt},
	}), nil
}

// Initiate payment
// (POST /payments)
func (s *Service) InitiatePayment(ctx context.Context, request api.InitiatePaymentRequestObject) (api.InitiatePaymentResponseObject, error) {
	// Resolve destination account
	destinationAccount, err := s.Repository.AccountRepository.GetById(ctx, request.Body.Destination)
	if err != nil {
		return api.InitiatePayment404Response{}, err
	}

	usdQuote, err := s.EthereumBlockchain.GetQuote(ctx, "USD", request.Body.Currency, request.Body.Amount)
	if err != nil {
		return api.InitiatePayment500Response{}, err
	}

	paymentId := workflows.NewPaymentWorkflowId()

	input := &workflows.PaymentWorkflowInput{
		Amount:      request.Body.Amount,
		Currency:    request.Body.Currency,
		Destination: destinationAccount.Id,
	}

	if _, err := s.TemporalClient.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			TaskQueue:                                TaskQueue,
			ID:                                       paymentId.String(),
			WorkflowIDReusePolicy:                    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
			WorkflowExecutionErrorWhenAlreadyStarted: true,
		},
		workflows.Payment,
		paymentId,
		input,
	); err != nil {
		return api.InitiatePayment500Response{}, err
	}

	val, err := s.TemporalClient.QueryWorkflow(ctx, paymentId.String(), "", workflows.PaymentQuery)
	if err != nil {
		return api.InitiatePayment500Response{}, err
	}

	var payment *model.Payment
	if err := val.Get(&payment); err != nil {
		return api.InitiatePayment500Response{}, err
	}

	return api.InitiatePayment201JSONResponse(api.Payment{
		Id:              payment.Id,
		Amount:          payment.Amount,
		Currency:        payment.Currency,
		Destination:     payment.Destination,
		DestinationName: destinationAccount.Name,
		GatewayWalletId: s.BaseBlockchain.Gateway.Hex(),
		UsdAmount:       *usdQuote,
		CreatedAt:       types.Time{Time: payment.CreatedAt},
		UpdatedAt:       types.Time{Time: payment.UpdatedAt},
	}), nil
}

// Retrieve payment
// (GET /payments/{paymentId})
func (s *Service) GetPayment(ctx context.Context, request api.GetPaymentRequestObject) (api.GetPaymentResponseObject, error) {
	payment, err := s.Repository.PaymentRepository.GetById(ctx, request.PaymentId)
	if err != nil {
		return api.GetPayment404Response{}, err
	}

	destinationAccount, err := s.Repository.AccountRepository.GetById(ctx, payment.Destination)
	if err != nil {
		return api.GetPayment500Response{}, err
	}

	usdQuote, err := s.EthereumBlockchain.GetQuote(ctx, "USD", payment.Currency, payment.Amount)
	if err != nil {
		return api.GetPayment500Response{}, err
	}

	return api.GetPayment200JSONResponse(api.Payment{
		Id:              payment.Id,
		Amount:          payment.Amount,
		Currency:        payment.Currency,
		Destination:     payment.Destination,
		DestinationName: destinationAccount.Name,
		GatewayWalletId: s.BaseBlockchain.Gateway.Hex(),
		UsdAmount:       *usdQuote,
		CreatedAt:       types.Time{Time: payment.CreatedAt},
		UpdatedAt:       types.Time{Time: payment.UpdatedAt},
	}), nil
}

// Submit payment
// (PATCH /payments/{paymentId}/submit)
func (s *Service) SubmitPayment(ctx context.Context, request api.SubmitPaymentRequestObject) (api.SubmitPaymentResponseObject, error) {
	update, err := s.TemporalClient.UpdateWorkflow(
		ctx,
		request.PaymentId.String(),
		"",
		workflows.PaymentSubmissionUpdate,
		&workflows.PaymentSubmissionInput{
			SourceWalletId: request.Body.PayerWalletId,
			UsdAmount:      request.Body.UsdAmount,
		},
	)
	if err != nil {
		return api.SubmitPayment500Response{}, err
	}

	var payment model.Payment
	if err := update.Get(ctx, &payment); err != nil {
		return api.SubmitPayment500Response{}, err
	}

	return api.SubmitPayment202Response{}, nil
}
