package entity

import (
	"time"

	"github.com/google/uuid"
)

type Basket struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;"`
	CreatedAt  time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
