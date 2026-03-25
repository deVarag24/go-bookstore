package controllers

import (
	"github.com/deVarag24/go-bookstore/pkg/models"
	"github.com/deVarag24/go-bookstore/pkg/services"
	"github.com/deVarag24/go-bookstore/pkg/utils/apiResponse"
	"github.com/gofiber/fiber/v2"
)

type BookStoreController interface {
	GetBookByID(ctx *fiber.Ctx) error
	GetBooks(ctx *fiber.Ctx) error
	CreateBook(ctx *fiber.Ctx) error
	UpdateBook(ctx *fiber.Ctx) error
	DeleteBook(ctx *fiber.Ctx) error
}

type bookStoreController struct {
	bookStoreService services.BookStoreService
}

func NewBookStoreController(boolStoreService services.BookStoreService) BookStoreController {
	return &bookStoreController{bookStoreService: boolStoreService}
}

func (c *bookStoreController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse.NewErrorResponse("Failed to Retrieve Book", map[string]interface{}{"error": "Invalid book ID"}))
	}

	book, err := c.bookStoreService.GetBookByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(apiResponse.NewErrorResponse("Failed to Retrieve Book", map[string]interface{}{"error": err.Error()}))
	}

	return ctx.JSON(apiResponse.NewSuccessResponse("Book retrieved successfully", book))

}

func (c *bookStoreController) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.bookStoreService.GetAllBooks()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse.NewErrorResponse("Failed to Retrieve Books", map[string]interface{}{"error": err.Error()}))
	}

	return ctx.JSON(apiResponse.NewSuccessResponse("Books retrieved successfully", books))
}

func (c *bookStoreController) CreateBook(ctx *fiber.Ctx) error {
	type CreateBookRequest struct {
		Name   string  `json:"name"`
		Author string  `json:"author"`
		Price  float64 `json:"price"`
	}

	var req CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse.NewErrorResponse("Failed to Create Book", map[string]interface{}{"error": "Invalid request body"}))
	}

	book, err := c.bookStoreService.CreateBook(req.Name, req.Author, req.Price)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse.NewErrorResponse("Failed to Create Book", map[string]interface{}{"error": err.Error()}))
	}

	return ctx.JSON(apiResponse.NewSuccessResponse("Book created successfully", book))
}

func (c *bookStoreController) UpdateBook(ctx *fiber.Ctx) error {
	type UpdateBookRequest struct {
		Name   string  `json:"name"`
		Author string  `json:"author"`
		Price  float64 `json:"price"`
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse.NewErrorResponse("Failed to Update Book", map[string]interface{}{"error": "Invalid book ID"}))
	}

	var req UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse.NewErrorResponse("Failed to Update Book", map[string]interface{}{"error": "Invalid request body"}))
	}
	book := &models.Book{
		ID:     uint(id),
		Name:   req.Name,
		Author: req.Author,
		Price:  req.Price,
	}
	err = c.bookStoreService.UpdateBook(book)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse.NewErrorResponse("Failed to Update Book", map[string]interface{}{"error": err.Error()}))
	}
	return ctx.JSON(apiResponse.NewSuccessResponse("Book updated successfully", book))
}

func (c *bookStoreController) DeleteBook(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse.NewErrorResponse("Failed to Delete Book", map[string]interface{}{"error": "Invalid book ID"}))
	}

	err = c.bookStoreService.DeleteBook(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse.NewErrorResponse("Failed to Delete Book", map[string]interface{}{"error": err.Error()}))
	}

	return ctx.JSON(apiResponse.NewSuccessResponse("Book deleted successfully", nil))
}
