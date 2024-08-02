package models

import (
	"github.com/google/uuid"
)

type TypeOfGuest int

const (
	Adult TypeOfGuest = iota
	Child
	NewBorn
)

// Define a struct to represent your database model
type (
	Guest struct {
		ID        uint      `gorm:"primaryKey;autoIncrement" validate:"required,number,min=1,max=200"`
		FirstName string    `gorm:"not null;type:varchar(30)" validate:"required,ascii,min=3,max=20"`
		LastName  string    `gorm:"not null;type:varchar(30)" validate:"required,ascii,min=3,max=20"`
		UUID      uuid.UUID `gorm:"not null;type:uuid"`
		Secret    string    `gorm:"not null;type:varchar(100)"`
		Confirmed bool      `gorm:"type:bool"`
		Notes     string    `gorm:"type:varchar(100)" validate:"omitempty,ascii,max=100"`
		Type      int       `gorm:"type:int" validate:"number,min=0,max=2"`
	}

	Guests []Guest

	Users struct {
		ID       uint   `gorm:"primaryKey;autoIncrement" validate:"required,number,min=1,max=200"`
		Username string `gorm:"not null;type:varchar(30)" validate:"required,ascii,min=3,max=20"`
		Password string `gorm:"not null;type:varchar(100)"`
		Role     string `gorm:"type:string"`
	}
)

const (
	Admin int = iota
	User
	Auditor
)
