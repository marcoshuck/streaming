package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ISBN            string    `json:"isbn"`
	Authors         string    `json:"authors"`
	PublicationDate time.Time `json:"publication_date"`
	Publisher       string    `json:"publisher"`
}
