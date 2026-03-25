package repository

import (
	"github.com/deVarag24/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

type BooksRepository interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

type booksRepository struct {
	db *gorm.DB
}

func NewBooksRepository(db *gorm.DB) BooksRepository {
	return &booksRepository{db: db}
}

func (r *booksRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *booksRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *booksRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *booksRepository) UpdateBook(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *booksRepository) DeleteBook(id uint) error {
	return r.db.Where("id=?", id).Delete(&models.Book{}).Error
}
