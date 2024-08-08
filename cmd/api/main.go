package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satyam-jha-16/event-manager/handlers"
	"github.com/satyam-jha-16/event-manager/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "EventLink",
		ServerHeader: "Fiber",
	})
	
	eventRepository := repositories.NewEventRepository(nil)
	
	server := app.Group("/api")
	
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	
	app.Listen(":8000")
}