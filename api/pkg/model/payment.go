package model

import (
	"time"

	"github.com/offblocks/offblocks-common/types"
)

type Payment struct {
	Id types.UUID `gorm:"type:uuid;default:gen_random_uuid()"`

	Amount   types.Decimal `gorm:"not null;"`
	Currency string        `gorm:"not null;"`

	SourceWalletId *string
	Destination    types.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"not null;index;"`
	UpdatedAt time.Time `gorm:"not null;index"`
}
