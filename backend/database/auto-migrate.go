package database

import (
	"github.com/pranay/ticket-booking-app/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User_Booking{})
}
