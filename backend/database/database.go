package database

import (
	"errors"
	"fmt"
	"wedding/log"
	"wedding/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	USERPASSWORD string
)

const (
	USERPASSWORDFILEDOCKER string = "/run/secrets/password_db"
	USERPASSWORDFILE       string = "../password_db.txt"
	USER                   string = "user"    // not a secret
	DBNAME                 string = "wedding" // not a secret
	PORT                   int    = 3306
	ADDRESS                string = "localhost"
	DSNFORMAT              string = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	var dbHost string = ADDRESS
	if runningInDocker() {
		dbHost = "mariadb"
	}
	var err error
	readUserPassword()
	USERPASSWORD = urlencode(USERPASSWORD)
	DSN := fmt.Sprintf(
		DSNFORMAT,
		USER,
		USERPASSWORD,
		dbHost,
		PORT,
		DBNAME,
	)

	// Establish a connection to the MySQL database
	db, err = gorm.Open(mysql.Open(DSN))
	if err != nil {
		log.Errorf("failed to connect database: %s\n", err.Error())
	}
	//defer db.Close()

	if isTableEmpty("guests") {
		log.Infof("table guests is empty: initializing it")
		log.Infof("initiating db from guests.csv")
		createGuestList()
		// Automigrate the schema, creating the users table if it doesn't exist
		err = db.AutoMigrate(&models.Guest{})
		if err != nil {
			log.Errorf("error during db initialization: %s\n", err.Error())
		}
		err = db.AutoMigrate(&models.Users{})
		if err != nil {
			log.Errorf("error during db initialization: %s\n", err.Error())
		}
		for _, g := range GUESTS {
			db.Where(&models.Guest{
				FirstName: g.FirstName,
				LastName:  g.LastName,
			}).FirstOrCreate(&g)
		}
	} else {
		log.Infof("table guests already initialized")
	}
}

// CreateGuest function insert new guests in guest table
// returns index (uint) and error
func CreateGuest(guest models.Guest) (uint, error) {
	n := CountGuests(guest.UUID)
	if n < 5 {
		result := db.Create(&guest)
		index := guest.ID   // returns inserted data's primary key
		err := result.Error // returns error

		if err != nil {
			return 0, err
		}
		return index, err
	}
	return 0, fmt.Errorf("too many guests")
}

// GuestExists function returns true if guest exists, false otherwise
func GuestExists(id uint, u uuid.UUID) bool {
	guest := models.Guest{}
	result := db.First(&guest, "id = ? AND uuid = ?", id, u)
	return result.Error == nil
}

// GuestExistsByUUID function returns true if guest exists, false otherwise
func GuestExistsByUUID(u uuid.UUID) bool {
	guest := models.Guest{}
	result := db.First(&guest, &models.Guest{
		UUID: u,
	})
	return result.Error == nil
}

// CountGuests function returns the number of guests given the UUID
func CountGuests(u uuid.UUID) (total int64) {
	db.Model(&models.Guest{}).
		Where("uuid = ?", u).
		Count(&total)
	return total
}

// UpdateGuest function update guest data
// returns index (uint) and error
func UpdateGuest(update models.Guest) error {
	if !GuestExists(update.ID, update.UUID) {
		return fmt.Errorf(
			"user %s %s (%s) not found",
			update.FirstName,
			update.LastName,
			update.UUID,
		)
	}
	// add select to update empty fields: https://gorm.io/docs/update.html#Update-Selected-Fields
	result := db.Model(&models.Guest{}).
		Select("FirstName", "LastName", "Confirmed", "Notes", "Type").
		Where(&models.Guest{ID: update.ID, UUID: update.UUID}).
		Updates(models.Guest{
			FirstName: update.FirstName,
			LastName:  update.LastName,
			Confirmed: true,
			Notes:     update.Notes,
			Type:      update.Type,
		})
	err := result.Error // returns error
	if err != nil {
		return err
	}
	return nil
}

// GetUsersByUUID function returns guests slice given a UUID
func GetUsersByUUID(u uuid.UUID) ([]models.Guest, error) {
	var guests []models.Guest
	result := db.Find(&guests, &models.Guest{
		UUID: u,
	})
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return []models.Guest{}, fmt.Errorf("user not found")
	}
	return guests, nil
}

// GetMainUserByUUID function returns first and last name of main user given the UUID
func GetMainUserByUUID(u uuid.UUID) (firstName string, lastName string, err error) {
	var guest models.Guest
	result := db.First(&guest, &models.Guest{
		UUID: u,
	})
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("user not found")
		return "", "", err
	}
	if guest.FirstName != "" {
		firstName = guest.FirstName
	}
	if guest.LastName != "" {
		lastName = guest.LastName
	}
	return firstName, lastName, nil
}

// GetUserByID function returns guests slice given a UUID
func GetUserByID(id uint) (models.Guest, error) {
	var guest models.Guest
	result := db.First(&guest, &models.Guest{
		ID: id,
	})
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Guest{}, fmt.Errorf("user not found")
	}
	return guest, nil
}

// DeleteGuest function delete a guest given its uuid and id
func DeleteGuest(id uint, u uuid.UUID) error {
	guest := models.Guest{}
	if GuestExists(id, u) {
		if n := CountGuests(u); n == 1 {
			return fiber.NewError(fiber.StatusForbidden, "you can't delete the last user")
		}
		result := db.Delete(&guest, "id = ? AND uuid = ?", id, u)
		return result.Error
	} else {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
}

func GetAllUsers() models.Guests {
	guests := &models.Guests{}
	db.Find(guests)
	return *guests
}
