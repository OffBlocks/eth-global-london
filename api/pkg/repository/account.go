package repository

import (
	"context"

	"github.com/offblocks/eth-global-london/api/pkg/model"
	"github.com/offblocks/offblocks-common/types"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetById(ctx context.Context, id types.UUID) (*model.Account, error)
	GetAll(ctx context.Context) ([]*model.Account, error)
	Save(ctx context.Context, account *model.Account) error
}

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) AccountRepository {
	return &accountRepository{DB}
}

func (p *accountRepository) GetById(ctx context.Context, id types.UUID) (*model.Account, error) {
	var account model.Account

	err := p.DB.WithContext(ctx).Where(&model.Account{Id: id}).First(&account).Error

	return &account, err
}

func (p *accountRepository) GetAll(ctx context.Context) ([]*model.Account, error) {
	var accounts []*model.Account

	err := p.DB.WithContext(ctx).Find(&accounts).Error

	return accounts, err
}

func (p *accountRepository) Save(ctx context.Context, account *model.Account) error {
	return p.DB.WithContext(ctx).Save(account).Error
}
