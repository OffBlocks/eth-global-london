package model

import (
	"time"

	"github.com/offblocks/eth-global-london/api/pkg/api"
	"github.com/offblocks/offblocks-common/types"
)

type Account struct {
	Id types.UUID `gorm:"type:uuid;default:gen_random_uuid()"`

	Name string `gorm:"not null;"`

	Identifier     string                    `gorm:"not null;index;"`
	IdentifierType api.AccountIdentifierType `gorm:"not null;"`

	WalletId string `gorm:"not null;index;"`

	CreatedAt time.Time `gorm:"not null;index;"`
	UpdatedAt time.Time `gorm:"not null;index"`
}
