package event

import (
	"github.com/google/uuid"
)

type OrderBasketItem struct {
	ProductID uuid.UUID
	Quantity  int
}

type OrderCreated struct {
	ID          uuid.UUID
	CustomerID  uuid.UUID
	TotalAmount float64
	Items       []OrderBasketItem
}
