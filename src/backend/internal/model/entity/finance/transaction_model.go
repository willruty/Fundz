package entity

import (
	"fundz/internal/model/entity"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Transaction_id uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"transaction_id" binding:"required"`
	UserID         uuid.UUID   `gorm:"type:uuid; not null" json:"user_id"`
	User           entity.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	Transaction_type        string          `gorm:"type:varchar(250);not null" json:"transaction_type" binding:"required"`
	Transaction_amount      decimal.Decimal `gorm:"type:numeric(20,2);not null" json:"transaction_amount" binding:"required"`
	Transaction_description string          `gorm:"type:varchar(250);not null" json:"transaction_description" binding:"required"`
	Transaction_date        time.Time       `gorm:"type:timestamp(0);not null" json:"transaction_date" binding:"required"`
}
