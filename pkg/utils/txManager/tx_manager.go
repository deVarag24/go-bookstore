package txmanager

import (
	"github.com/deVarag24/go-bookstore/pkg/repository"
	"gorm.io/gorm"
)

type TxManager interface {
	WithTransaction(fn func(uow UnitOfWork) error) error
}

type txManager struct {
	db *gorm.DB
}

func NewTxManager(db *gorm.DB) TxManager {
	return &txManager{db: db}
}

func (tm *txManager) WithTransaction(fn func(uow UnitOfWork) error) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	uow := &unitOfWork{
		booksRepo: repository.NewBooksRepository(tx),
	}
	if err := fn(uow); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
