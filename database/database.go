package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
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
	db, err := gorm.Open("mysql", DSN)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Automigrate the schema, creating the users table if it doesn't exist
	db.AutoMigrate(&Guest{})
}
