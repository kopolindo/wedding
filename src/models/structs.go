package models

import (
	"github.com/google/uuid"
)

// Define a struct to represent your database model
type Guest struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	FirstName string    `gorm:"not null;type:varchar(30)"`
	LastName  string    `gorm:"not null;type:varchar(30)"`
	UUID      uuid.UUID `gorm:"not null;type:uuid"`
	Secret    string    `gorm:"not null;type:varchar(100)"`
	Confirmed bool      `gorm:"type:bool"`
	Notes     []byte
}

type Guests []Guest
