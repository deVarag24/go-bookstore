package dicontainer

import (
	"github.com/deVarag24/go-bookstore/pkg/controllers"
	"github.com/deVarag24/go-bookstore/pkg/repository"
	"github.com/deVarag24/go-bookstore/pkg/services"
	"gorm.io/gorm"
)

type DIContainer struct {
	Controllers Controllers
}

type Controllers struct {
	BookStoreController controllers.BookStoreController
}

func NewDIContainer(db *gorm.DB) *DIContainer {
	booksRepository := repository.NewBooksRepository(db)
	bookStoreService := services.NewBookStoreService(booksRepository)
	bookStoreController := controllers.NewBookStoreController(bookStoreService)
	return &DIContainer{
		Controllers: Controllers{
			BookStoreController: bookStoreController,
		},
	}
}
