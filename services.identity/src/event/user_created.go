package event

import "github.com/google/uuid"

type UserCreated struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
}
