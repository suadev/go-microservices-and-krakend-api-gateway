package event

import (
	"github.com/google/uuid"
)

type orderBasketItem struct {
	ProductID uuid.UUID
	Quantity  int
}

type OrderCreated struct {
	ID          uuid.UUID
	CustomerID  uuid.UUID
	TotalAmount float64
	Items       []orderBasketItem
}
