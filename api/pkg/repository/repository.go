package repository

import (
	"github.com/offblocks/eth-global-london/api/pkg/db"
)

type Repository struct {
	AccountRepository AccountRepository
	PaymentRepository PaymentRepository
}

func NewRepository(h *db.Handler) Repository {
	return Repository{NewAccountRepository(h.DB), NewPaymentRepository(h.DB)}
}
