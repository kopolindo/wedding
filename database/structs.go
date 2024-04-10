package database

import (
	"github.com/google/uuid"
)

// Define a struct to represent your database model
type Guest struct {
	ID                   uint      `gorm:"primary_key;autoIncrement"`
	UUID                 uuid.UUID `gorm:"not null"`
	NumberOfPartecipants int       `gorm:"not null"`
	Notes                []byte
}

var (
	USERPASSWORD string
)

const (
	USERPASSWORDFILE string = "password_db.txt"
	USER             string = "user"    // not a secret
	DBNAME           string = "wedding" // not a secret
	PORT             int    = 3306
	ADDRESS          string = "localhost"
	DSNFORMAT        string = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)
