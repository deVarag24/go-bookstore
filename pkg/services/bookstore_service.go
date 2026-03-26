package services

import (
	"errors"
	"math/rand"

	"github.com/deVarag24/go-bookstore/pkg/models"
	"github.com/deVarag24/go-bookstore/pkg/repository"
	txmanager "github.com/deVarag24/go-bookstore/pkg/utils/txManager"
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
	txManager txmanager.TxManager
}

func NewBookStoreService(booksRepo repository.BooksRepository, txManager txmanager.TxManager) BookStoreService {
	return &bookStoreService{booksRepo: booksRepo, txManager: txManager}
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
	return s.txManager.WithTransaction(func(uow txmanager.UnitOfWork) error {
		existingBook, err := uow.Books().GetBookByID(book.ID)
		if err != nil {
			return errors.New("book not found: " + err.Error())
		}

		existingBook.Name = book.Name
		existingBook.Author = book.Author
		existingBook.Price = book.Price

		err = uow.Books().UpdateBook(existingBook)
		if err != nil {
			return errors.New("failed to update book: " + err.Error())
		}
		return nil
	})
}

func (s *bookStoreService) DeleteBook(id uint) error {
	err := s.booksRepo.DeleteBook(id)
	if err != nil {
		return errors.New("failed to delete book: " + err.Error())
	}
	return nil
}
