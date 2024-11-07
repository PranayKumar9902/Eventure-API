package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pranay/ticket-booking-app/models"
	"github.com/pranay/ticket-booking-app/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository models.AuthRepository
}

func (a *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User_Booking, error) {

	user, err := a.authRepository.GetUser(ctx, "email = ?", loginData.Email)
	if err != nil {
		return "", nil, err
	}

	if !models.MatchesHash(loginData.Password, user.Password) {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	secretKey := os.Getenv("JWT_SECRET")

	tokenString, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, secretKey)

	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func (a *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User_Booking, error) {

	if !models.IsValidEmail(registerData.Email) {
		return "", nil, fmt.Errorf("invalid email")
	}

	_, err := a.authRepository.GetUser(ctx, "email = ?", registerData.Email)

	if err == nil {
		return "", nil, fmt.Errorf("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", nil, err
	}

	registerData.Password = string(hashedPassword)

	user, err := a.authRepository.RegisterUser(ctx, registerData)

	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	secretKey := os.Getenv("JWT_SECRET")

	tokenString, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, secretKey)

	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		authRepository: repository,
	}
}
