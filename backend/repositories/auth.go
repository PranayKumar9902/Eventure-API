package repositories

import (
	"context"

	"github.com/pranay/ticket-booking-app/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (a *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User_Booking, error) {
	user := &models.User_Booking{
		Email:    registerData.Email,
		Password: registerData.Password,
	}
	err := a.db.Create(user).Error
	return user, err
}

func (a *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User_Booking, error) {
	user := &models.User_Booking{}

	if err := a.db.Where(query, args...).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
