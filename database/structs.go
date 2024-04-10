package database

// Define a struct to represent your database model
type User struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Email string `gorm:"unique"`
}
