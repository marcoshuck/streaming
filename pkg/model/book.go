package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
}
