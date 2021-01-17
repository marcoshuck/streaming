package service

import (
	"github.com/marcoshuck/streaming/pkg/model"
	"github.com/marcoshuck/streaming/pkg/repository"
)

type Book interface {
	ListAll() ([]model.Book, error)
	ListByISBN(isbn string) ([]model.Book, error)
}

type book struct {
	repository repository.Book
}

func (b *book) ListAll() ([]model.Book, error) {
	books, err := b.repository.Find(nil)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *book) ListByISBN(isbn string) ([]model.Book, error) {
	err := b.validateISBN(isbn)
	if err != nil {
		return nil, err
	}

	books, err := b.repository.Find(&isbn)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *book) validateISBN(isbn string) error {
	return nil
}

func NewBookService(repository repository.Book) Book {
	return &book{
		repository: repository,
	}
}
