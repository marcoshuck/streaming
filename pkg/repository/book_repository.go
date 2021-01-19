package repository

import (
	"github.com/marcoshuck/streaming/pkg/model"
	"gorm.io/gorm"
	"log"
)

type Book interface {
	Find(isbn *string) ([]model.Book, error)
}

type book struct {
	db *gorm.DB
}

func (b *book) Find(isbn *string) ([]model.Book, error) {
	var books []model.Book

	q := b.db.Model(&model.Book{})

	if isbn != nil {
		q = q.Where("isbn = ?", *isbn)
	}

	if err := q.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func NewBookRepository(db *gorm.DB) Book {
	if db == nil {
		log.Fatalln("db has not been provided")
		return nil
	}

	return &book{
		db: db,
	}
}
