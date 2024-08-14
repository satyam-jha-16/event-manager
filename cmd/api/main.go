package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/satyam-jha-16/event-manager/config"
	"github.com/satyam-jha-16/event-manager/db"
	"github.com/satyam-jha-16/event-manager/handlers"
	"github.com/satyam-jha-16/event-manager/middlewares"
	"github.com/satyam-jha-16/event-manager/repositories"
	"github.com/satyam-jha-16/event-manager/services"
)

func main() {
	
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "EventLink",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	server := app.Group("/api")
	//services 
	authService := services.NewAuthService(authRepository)
	handlers.NewAuthHandler(server.Group("/auth"), authService)
	
	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTickethandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
