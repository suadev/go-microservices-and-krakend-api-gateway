package entity

import (
	"time"

	"github.com/google/uuid"
)

type BasketItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	BasketID    uuid.UUID `gorm:"type:uuid;not null;"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;"`
	ProductName string    `gorm:"size:255;not null;"`
	UnitPrice   float64   `gorm:"not null;"`
	Quantity    int       `gorm:"not null;"`
	CreatedAt   time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
