package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/satyam-jha-16/event-manager/models"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancel()
	
	events, err := h.repository.GetMany(context)
	
	if err!= nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status" : "fail",
			"error": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status" : "success",
		"data": events,
		"message" : "",
	})
}

func (h *EventHandler) GetOne(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) Create(c *fiber.Ctx)error {
	return nil
}


func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.Create)
	router.Get("/:eventId", handler.GetOne)
}
