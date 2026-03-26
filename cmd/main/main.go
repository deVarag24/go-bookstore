package main

import (
	"github.com/deVarag24/go-bookstore/pkg/config"
	dicontainer "github.com/deVarag24/go-bookstore/pkg/diContainer"
	"github.com/deVarag24/go-bookstore/pkg/models"
	"github.com/deVarag24/go-bookstore/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	config.ConnectDb()
	// Migrate the database schema
	migerateEntities()

	// Initialize the Fiber app
	app := fiber.New()

	// Register routes
	setupRoutes(app)

	// Start the server
	app.Listen(":3000")
}

func migerateEntities() {
	// Migrate the database schema
	db := config.GetDb()
	db.AutoMigrate(&models.Book{})
}

func setupRoutes(app *fiber.App) {
	db := config.GetDb()
	diContainer := dicontainer.NewDIContainer(db)
	routes.RegisterRoutes(app, diContainer)
}
