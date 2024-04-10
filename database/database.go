package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	// Establish a connection to the MySQL database
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Automigrate the schema, creating the users table if it doesn't exist
	db.AutoMigrate(&User{})
}
