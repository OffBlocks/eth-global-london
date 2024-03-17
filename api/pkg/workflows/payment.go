package workflows

import (
	"errors"
	"time"

	"github.com/offblocks/eth-global-london/api/pkg/activities"
	"github.com/offblocks/eth-global-london/api/pkg/model"
	"github.com/offblocks/offblocks-common/types"
	"go.temporal.io/sdk/workflow"
)

const (
	PaymentQuery            = "payment"
	PaymentSubmissionUpdate = "payment-submission"
)

type PaymentWorkflowInput struct {
	Amount      types.Decimal
	Currency    string
	Destination types.UUID
}
type PaymentState = model.Payment
type PaymentWorkflowResult = PaymentState

type PaymentSubmissionInput struct {
	SourceWalletId string
	UsdAmount      types.Decimal
}

// paymentWorkflow represents the state for an account workflow execution.
type paymentWorkflow struct {
	*PaymentState
	Beneficiary *model.Account
	UsdAmount   types.Decimal
	SubmitC     workflow.Channel
}

// newPaymentWorkflow initializes an paymentWorkflow struct.
func newPaymentWorkflow(ctx workflow.Context, state *PaymentState) *paymentWorkflow {
	return &paymentWorkflow{
		PaymentState: state,
		SubmitC:      workflow.NewChannel(ctx),
	}
}

func (w *paymentWorkflow) fetchAccount(ctx workflow.Context, id types.UUID) (*model.Account, error) {
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{
		ScheduleToCloseTimeout: 1 * time.Second,
	})

	var account *model.Account

	err := workflow.ExecuteLocalActivity(ctx, a.FetchAccount, &id).Get(ctx, &account)
	return account, err
}

func (w *paymentWorkflow) savePayment(ctx workflow.Context) error {
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{
		ScheduleToCloseTimeout: time.Second,
	})

	return workflow.ExecuteLocalActivity(ctx, a.SavePayment, w.PaymentState).Get(ctx, nil)
}

func (w *paymentWorkflow) waitForPaymentSubmission(ctx workflow.Context) error {
	isValid := func(ctx workflow.Context, update PaymentSubmissionInput) error {
		if w.PaymentState.SourceWalletId != nil {
			return errors.ErrUnsupported
		}
		return nil
	}

	if err := workflow.SetUpdateHandlerWithOptions(
		ctx,
		PaymentSubmissionUpdate,
		func(ctx workflow.Context, input PaymentSubmissionInput) (*PaymentWorkflowResult, error) {
			w.SourceWalletId = &input.SourceWalletId
			w.UsdAmount = input.UsdAmount

			if err := w.savePayment(ctx); err != nil {
				return w.PaymentState, err
			}

			w.SubmitC.SendAsync(struct{}{})
			return w.PaymentState, nil
		},
		workflow.UpdateHandlerOptions{Validator: isValid},
	); err != nil {
		return err
	}

	return nil
}

func (w *paymentWorkflow) transferToken(ctx workflow.Context) (*activities.TransferTokensResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Minute,
	})

	input := &activities.TransferTokensInput{
		FromWalletId: *w.SourceWalletId,
		Amount:       w.UsdAmount,
		Currency:     "USD",
	}

	var result activities.TransferTokensResult

	err := workflow.ExecuteActivity(ctx, a.TransferTokens, &input).Get(ctx, &result)
	return &result, err
}

func (w *paymentWorkflow) waitForTransferFinished(ctx workflow.Context, transactionId string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Hour,
		HeartbeatTimeout:       time.Minute,
	})

	return workflow.ExecuteActivity(ctx, a.IsTokenTransferFinished, &transactionId).Get(ctx, nil)
}

func (w *paymentWorkflow) swapToken(ctx workflow.Context) (*activities.SwapTokensResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Minute,
	})

	input := &activities.SwapTokensInput{
		Beneficiary:    w.Beneficiary.WalletId,
		MaxInputAmount: w.UsdAmount,
		InputCurrency:  "USD",
		OutputAmount:   w.Amount,
		OutputCurrency: w.Currency,
	}

	var result activities.SwapTokensResult

	err := workflow.ExecuteActivity(ctx, a.SwapTokens, &input).Get(ctx, &result)
	return &result, err
}

// Payment is a Workflow Definition that calls for the execution of a variable set of Activities and Child Workflows.
// This is the main entry point of the application.
func Payment(ctx workflow.Context, paymentId types.UUID, input *PaymentWorkflowInput) (*PaymentWorkflowResult, error) {
	w := newPaymentWorkflow(
		ctx,
		&PaymentState{
			Id:          paymentId,
			Amount:      input.Amount,
			Currency:    input.Currency,
			Destination: input.Destination,
			CreatedAt:   workflow.Now(ctx),
			UpdatedAt:   workflow.Now(ctx),
		},
	)

	// The query returns the state of a account and is used by the API.
	if err := workflow.SetQueryHandler(ctx, PaymentQuery, func() (*PaymentState, error) {
		return w.PaymentState, nil
	}); err != nil {
		return w.PaymentState, err
	}

	account, err := w.fetchAccount(ctx, input.Destination)
	if err != nil {
		return w.PaymentState, err
	}
	w.Beneficiary = account

	if err := w.savePayment(ctx); err != nil {
		return w.PaymentState, err
	}

	if err := w.waitForPaymentSubmission(ctx); err != nil {
		return w.PaymentState, err
	}

	s := workflow.NewSelector(ctx)
	s.AddReceive(w.SubmitC, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, nil)
	})

	for {
		s.Select(ctx)
		if w.SourceWalletId != nil {
			break
		}
	}

	tx, err := w.transferToken(ctx)
	if err != nil {
		return w.PaymentState, err
	}

	if err := w.waitForTransferFinished(ctx, tx.TransactionId); err != nil {
		return w.PaymentState, err
	}

	if _, err := w.swapToken(ctx); err != nil {
		return w.PaymentState, err
	}

	return w.PaymentState, nil
}
