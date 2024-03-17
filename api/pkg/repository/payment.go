package repository

import (
	"context"

	"github.com/offblocks/eth-global-london/api/pkg/model"
	"github.com/offblocks/offblocks-common/types"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetById(ctx context.Context, id types.UUID) (*model.Payment, error)
	Save(ctx context.Context, account *model.Payment) error
}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) PaymentRepository {
	return &paymentRepository{DB}
}

func (p *paymentRepository) GetById(ctx context.Context, id types.UUID) (*model.Payment, error) {
	var payment model.Payment

	err := p.DB.WithContext(ctx).Where(&model.Payment{Id: id}).First(&payment).Error

	return &payment, err
}

func (p *paymentRepository) Save(ctx context.Context, payment *model.Payment) error {
	return p.DB.WithContext(ctx).Save(payment).Error
}
