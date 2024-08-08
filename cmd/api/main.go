package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/satyam-jha-16/event-manager/config"
	"github.com/satyam-jha-16/event-manager/db"
	"github.com/satyam-jha-16/event-manager/handlers"
	"github.com/satyam-jha-16/event-manager/repositories"
)

func main() {
	
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "EventLink",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
