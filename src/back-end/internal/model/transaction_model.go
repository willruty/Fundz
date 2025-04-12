package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"transaction_id"`
	UserID        uuid.UUID `gorm:"type:uuid; not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	TransactionType        string          `gorm:"type:varchar(7);not null" json:"transaction_type"`
	TransactionAmount      decimal.Decimal `gorm:"type:numeric(12,2); not null" json:"transaction_amount"`
	TransactionDescription string          `gorm:"type:varchar(100);not null" json:"transaction_description"`
	TransactionDate        time.Time       `gorm:"type:timestamp(0); not null" json:"transaction_date"`
}
