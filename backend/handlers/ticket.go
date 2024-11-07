package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pranay/ticket-booking-app/models"
	"github.com/skip2/go-qrcode"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := c.Locals("user").(models.User_Booking)

	tickets, err := h.repository.GetMany(context, user.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"tickets": tickets,
	})

}

func (h *TicketHandler) GetOne(c *fiber.Ctx) error {

	ticketId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := c.Locals("user").(models.User_Booking)

	ticket, err := h.repository.GetOne(context, user.ID, uint(ticketId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var QRCode []byte

	QRCode, err = qrcode.Encode(
		fmt.Sprintf("ticketId:%v ownerId:%v", ticketId, user.ID),
		qrcode.Medium,
		256,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error generating QR code",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"ticket": ticket,
			"QRCode": QRCode,
		},
	})
}

func (h *TicketHandler) CreateOne(c *fiber.Ctx) error {

	ticket := models.Ticket{}

	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := c.Locals("user").(models.User_Booking)

	createdTicket, err := h.repository.CreateOne(context, user.ID, ticket)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket created successfully",
		"ticket":  createdTicket,
	})

}

func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if err := ctx.BodyParser(&validateBody); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := h.repository.UpdatedOne(context, validateBody.OwnerID, validateBody.TicketID, validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the show!",
		"data":    ticket,
	})
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {

	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:id", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Post("/validate", handler.ValidateOne)
}
