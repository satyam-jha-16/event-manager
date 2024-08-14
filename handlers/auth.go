package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/satyam-jha-16/event-manager/models"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	creds := &models.AuthCredentials{}
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	//validate the credentials
	if err := validate.Struct(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  fmt.Errorf("please, provide a valid name, email and password").Error(),
		})
	}
	token, user, err := h.service.Login(context, creds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
				"message": "Successfully logged in",
				"data": &fiber.Map{
					"token": token,
					"user":  user,
				},
	})
	
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	creds := &models.AuthCredentials{}
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	//validate the credentials
	if err := validate.Struct(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  fmt.Errorf("please, provide a valid name, email and password").Error(),
		})
	}
	
	token, user, err := h.service.Register(context, creds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
				"message": "Successfully logged in",
				"data": &fiber.Map{
					"token": token,
					"user":  user,
				},
	})

}

func NewAuthHandler(router fiber.Router, service models.AuthService) {
	handler := &AuthHandler{service: service}

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
}
