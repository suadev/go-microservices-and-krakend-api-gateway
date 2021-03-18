package entity

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus int

const (
	OrderCreated   OrderStatus = 0
	OrderCompleted OrderStatus = 1
	OrderFailed    OrderStatus = 2
)

type Order struct {
	ID          uuid.UUID   `gorm:"type:uuid;primary_key;"`
	CustomerID  uuid.UUID   `gorm:"type:uuid;not null;"`
	TotalAmount float64     `gorm:"not null;"`
	Status      OrderStatus `gorm:"not null;"`
	CreatedAt   time.Time   `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
