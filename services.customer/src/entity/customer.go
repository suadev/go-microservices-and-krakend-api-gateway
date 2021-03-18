package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Email       string    `gorm:"size:255;not null;"`
	FirstName   string    `gorm:"size:255;not null;"`
	LastName    string    `gorm:"size:255;not null;"`
	Address     string    `gorm:"size:255;"`
	PhoneNumber string    `gorm:"size:255;"`
	CreatedAt   time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
