package event

import "github.com/google/uuid"

type ProcuctCreated struct {
	ID       uuid.UUID
	Name     string
	Price    float64
	Quantity int
}
