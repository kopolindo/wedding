package database

import (
	"fmt"
	"log"
	"wedding/backend"
	"wedding/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db, err = gorm.Open("mysql", DSN)
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	// Automigrate the schema, creating the users table if it doesn't exist
	db.AutoMigrate(&models.Guest{})
	for _, g := range backend.GUESTS {
		InsertGuestData(g)
	}
}

func InsertGuestData(guest models.Guest) (uint, error) {
	result := db.Create(&guest)
	log.Printf("inserting guest: %v\n", guest)
	index := guest.ID   // returns inserted data's primary key
	err := result.Error // returns error

	if err != nil {
		return 0, err
	}
	return index, err
}
