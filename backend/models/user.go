package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  UserRole = "manager"
	Attendee UserRole = "attendee"
)

type User_Booking struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"text;unique;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:'attendee'"`
	Password  string    `json:"-"` //Donot Compute Password in JSON
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User_Booking) AfterCreate(db *gorm.DB) (err error) {

	if u.ID == 1 {
		db.Model(u).Update("role", Manager)
	}
	return
}
