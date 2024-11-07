package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/pranay/ticket-booking-app/database"
	"github.com/pranay/ticket-booking-app/handlers"
	"github.com/pranay/ticket-booking-app/middleware"
	"github.com/pranay/ticket-booking-app/repositories"
	"github.com/pranay/ticket-booking-app/services"
)

func main() {

	godotenv.Load(".env")

	database.ConnectToDatabase(database.DBMigrator)

	app := fiber.New()

	app.Use(logger.New())

	eventRepository := repositories.NewEventRepository(database.Database.Db)
	ticketRepository := repositories.NewTicketRepository(database.Database.Db)
	authRepository := repositories.NewAuthRepository(database.Database.Db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middleware.AuthProtected(database.Database.Db))

	handlers.NewEventHandler(privateRoutes.Group("/events"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/tickets"), ticketRepository)

	log.Fatal(app.Listen(":8007"))
}
