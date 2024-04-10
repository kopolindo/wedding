package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

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
	db.AutoMigrate(&Guest{})
}

func InsertGuestData(guest Guest) (uint, error) {
	result := db.Create(&guest)
	log.Printf("inserting guest: %v\n", guest)
	index := guest.ID   // returns inserted data's primary key
	err := result.Error // returns error

	if err != nil {
		return 0, err
	}
	return index, err
}
