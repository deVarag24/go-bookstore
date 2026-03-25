package routes

import (
	"github.com/deVarag24/go-bookstore/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookStoreRoutes(routes fiber.Router, controller controllers.BookStoreController) {
	bookStoreRoute := routes.Group("/book")

	bookStoreRoute.Post("/", controller.CreateBook)
	bookStoreRoute.Get("/", controller.GetBooks)
	bookStoreRoute.Get("/:id", controller.GetBookByID)
	bookStoreRoute.Put("/:id", controller.UpdateBook)
	bookStoreRoute.Delete("/:id", controller.DeleteBook)
}
