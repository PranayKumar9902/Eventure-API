package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pranay/ticket-booking-app/models"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events, err := h.repository.GetMany(context)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"events": events,
	})

}

func (h *EventHandler) GetOne(c *fiber.Ctx) error {

	eventId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event, err := h.repository.GetOne(context, uint(eventId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"event":  event,
	})
}

func (h *EventHandler) CreateOne(c *fiber.Ctx) error {

	event := models.Event{}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createdEvent, err := h.repository.CreateOne(context, event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Event created successfully",
		"event":   createdEvent,
	})
}

func (h *EventHandler) UpdateOne(c *fiber.Ctx) error {

	eventId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updateData := make(map[string]interface{})

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updatedEvent, err := h.repository.UpdatedOne(context, uint(eventId), updateData)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Event updated successfully",
		"event":   updatedEvent,
	})
}

func (h *EventHandler) DeleteOne(c *fiber.Ctx) error {

	eventId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.repository.DeleteOne(context, uint(eventId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Event deleted successfully",
	})
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {

	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:id", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Put("/:id", handler.UpdateOne)
	router.Delete("/:id", handler.DeleteOne)
}
