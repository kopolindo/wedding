package database

import (
	"errors"
	"fmt"
	"log"
	"wedding/backend"
	"wedding/models"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

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

func init() {
	var err error
	readUserPassword()
	DSN := fmt.Sprintf(
		DSNFORMAT,
		USER,
		USERPASSWORD,
		ADDRESS,
		PORT,
		DBNAME,
	)

	// Establish a connection to the MySQL database
	db, err = gorm.Open(mysql.Open(DSN))
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	// Automigrate the schema, creating the users table if it doesn't exist
	db.AutoMigrate(&models.Guest{})
	for _, g := range backend.GUESTS {
		db.Where(&models.Guest{
			FirstName: g.FirstName,
			LastName:  g.LastName,
		}).FirstOrCreate(&g)
	}
}

// CreateGuest function insert new guests in guest table
// returns index (uint) and error
func CreateGuest(guest models.Guest) (uint, error) {
	result := db.Create(&guest)
	log.Printf("inserting guest: %v\n", guest)
	index := guest.ID   // returns inserted data's primary key
	err := result.Error // returns error

	if err != nil {
		return 0, err
	}
	return index, err
}

// GuestExists function returns true if guest exists, false otherwise
func GuestExists(update models.Guest) bool {
	guest := models.Guest{}
	result := db.First(&guest, &models.Guest{
		FirstName: update.FirstName,
		LastName:  update.LastName,
		UUID:      update.UUID,
	})
	return result.Error == nil
}

// UpdateGuest function update guest data
// returns index (uint) and error
func UpdateGuest(update models.Guest) error {
	if !GuestExists(update) {
		return fmt.Errorf(
			"user %s %s (%s) not found",
			update.FirstName,
			update.LastName,
			update.UUID,
		)
	}
	result := db.Model(&models.Guest{}).
		Where(&models.Guest{UUID: update.UUID}).
		Updates(models.Guest{
			NumberOfPartecipants: update.NumberOfPartecipants,
			Confirmed:            update.Confirmed,
			Notes:                update.Notes,
		})
	err := result.Error // returns error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUUID function returns guest given a UUID
func GetUserByUUID(u uuid.UUID) (models.Guest, error) {
	guest := models.Guest{}
	// debugging!!
	result := db.Debug().First(&guest, &models.Guest{
		UUID: u,
	})
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Guest{}, fmt.Errorf("user not found")
	}
	return guest, nil
}
