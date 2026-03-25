package services

import (
	"errors"
	"math/rand"

	"github.com/deVarag24/go-bookstore/pkg/models"
	"github.com/deVarag24/go-bookstore/pkg/repository"
)

type BookStoreService interface {
	CreateBook(name, author string, price float64) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

type bookStoreService struct {
	booksRepo repository.BooksRepository
}

func NewBookStoreService(booksRepo repository.BooksRepository) BookStoreService {
	return &bookStoreService{booksRepo: booksRepo}
}

func (s *bookStoreService) CreateBook(name, author string, price float64) (*models.Book, error) {
	book := &models.Book{
		Name:   name,
		Author: author,
		Price:  price,
		ID:     uint(rand.Intn(10000)),
	}

	err := s.booksRepo.CreateBook(book)
	if err != nil {
		return nil, errors.New("failed to create book: " + err.Error())
	}

	return book, nil
}

func (s *bookStoreService) GetAllBooks() ([]*models.Book, error) {
	books, err := s.booksRepo.GetAllBooks()
	if err != nil {
		return nil, errors.New("failed to retrieve books: " + err.Error())
	}

	var bookPointers []*models.Book
	for i := range books {
		bookPointers = append(bookPointers, &books[i])
	}

	return bookPointers, nil
}

func (s *bookStoreService) GetBookByID(id uint) (*models.Book, error) {
	book, err := s.booksRepo.GetBookByID(id)
	if err != nil {
		return nil, errors.New("book not found: " + err.Error())
	}
	return book, nil
}

func (s *bookStoreService) UpdateBook(book *models.Book) error {
	err := s.booksRepo.UpdateBook(book)
	if err != nil {
		return errors.New("failed to update book: " + err.Error())
	}
	return nil
}

func (s *bookStoreService) DeleteBook(id uint) error {
	err := s.booksRepo.DeleteBook(id)
	if err != nil {
		return errors.New("failed to delete book: " + err.Error())
	}
	return nil
}
