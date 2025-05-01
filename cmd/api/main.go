package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:8082, http://192.168.29.250:8082", // add all expected origins here
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
    }))

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
