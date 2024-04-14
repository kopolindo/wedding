package models

import (
	"github.com/google/uuid"
)

// Define a struct to represent your database model
type Guest struct {
	ID                   uint      `gorm:"primary_key;autoIncrement"`
	FirstName            string    `gorm:"not null;type:varchar(30)"`
	LastName             string    `gorm:"not null;type:varchar(30)"`
	UUID                 uuid.UUID `gorm:"not null;type:uuid"`
	NumberOfPartecipants int       `gorm:"not null;type:int"`
	Confirmed            bool      `gorm:"type:bool"`
	Notes                []byte
}
