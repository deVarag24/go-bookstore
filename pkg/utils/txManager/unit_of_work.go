package txmanager

import "github.com/deVarag24/go-bookstore/pkg/repository"

type UnitOfWork interface {
	Books() repository.BooksRepository
}

type unitOfWork struct {
	booksRepo repository.BooksRepository
}

func (u *unitOfWork) Books() repository.BooksRepository {
	return u.booksRepo
}
