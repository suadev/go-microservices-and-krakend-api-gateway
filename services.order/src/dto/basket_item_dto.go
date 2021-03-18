package dto

import (
	"time"

	"github.com/google/uuid"
)

type BasketItemDto struct {
	ID          uuid.UUID
	BasketID    uuid.UUID
	ProductID   uuid.UUID
	ProductName string
	UnitPrice   float64
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
