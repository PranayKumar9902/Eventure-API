package handlers

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pranay/ticket-booking-app/models"
)

type AuthHandler struct {
	authService models.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	loginData := new(models.AuthCredentials)

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	token, user, err := h.authService.Login(context, loginData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	registerData := new(models.AuthCredentials)

	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(registerData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	token, user, err := h.authService.Register(context, registerData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func NewAuthHandler(route fiber.Router, service models.AuthService) {
	handler := &AuthHandler{
		authService: service,
	}

	route.Post("/login", handler.Login)
	route.Post("/register", handler.Register)
}
