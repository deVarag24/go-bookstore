package routes

import (
	dicontainer "github.com/deVarag24/go-bookstore/pkg/diContainer"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, diContainer *dicontainer.DIContainer) {
	routes := app.Group("/api/v1")
	RegisterBookStoreRoutes(routes, diContainer.Controllers.BookStoreController)
}
