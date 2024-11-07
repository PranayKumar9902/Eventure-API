package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *AuthCredentials) (*User_Booking, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User_Booking, error)
}

type AuthService interface {
	Login(ctx context.Context, loginData *AuthCredentials) (string, *User_Booking, error)
	Register(ctx context.Context, registerData *AuthCredentials) (string, *User_Booking, error)
}

//compare Password

func MatchesHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Valid email

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
