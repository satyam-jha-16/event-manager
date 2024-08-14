package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/satyam-jha-16/event-manager/models"
	"github.com/skip2/go-qrcode"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userId := uint(c.Locals("userId").(float64))

	tickets, err := h.repository.GetMany(context, userId)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   tickets,
	})
}

func (h *TicketHandler) GetOne(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userId := uint(c.Locals("userId").(float64))

	ticketId, _ := strconv.Atoi(c.Params("ticketId"))

	ticket, err := h.repository.GetOne(context, userId, uint(ticketId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	
	var QRCode []byte 
	QRCode, err = qrcode.Encode(
		fmt.Sprintf("ticketId: %v, ownerId:%v", ticket.ID, ticket.UserID), 
		qrcode.Medium, 
		256)
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}


	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   &fiber.Map{
			"ticket" : ticket,
			"qrcode" : QRCode,
		},
	})

}

func (h *TicketHandler) Create(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	
		userId := uint(c.Locals("userId").(float64))

	ticket := &models.Ticket{}

	// userId := uint(c.Locals("userId").(float64))
	if err := c.BodyParser(ticket); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	ticket, err := h.repository.Create(context, userId, ticket)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "failed",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket created",
		"data":    ticket,
	})
}

func (h *TicketHandler) ValidateOne(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if err := c.BodyParser(validateBody); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
			"data":   nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := h.repository.UpdateOne(context, validateBody.OwnerID, validateBody.TicketID, validateData)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "failed",
			"error":  err.Error(),
			"data":   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket validated",
		"data":    ticket,
	})
}

func NewTickethandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{repository: repository}

	router.Get("/", handler.GetMany)
	// router.Get("/", handler.Create)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/", handler.Create)
	router.Post("/validate", handler.ValidateOne)

}
