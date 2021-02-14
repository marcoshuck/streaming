package controller

import (
	"encoding/json"
	"fmt"
	"github.com/marcoshuck/streaming/pkg/service"
	"net/http"
)

type Book interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByISBN(w http.ResponseWriter, r *http.Request)
}

type book struct {
	service service.Book
}

func (b *book) GetByISBN(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (b *book) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := b.service.ListAll()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err),
			http.StatusInternalServerError,
		)
	}

	bslice, err := json.Marshal(list)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err),
			http.StatusInternalServerError,
		)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bslice)
}

func NewBookController(service service.Book) Book {
	return &book{
		service: service,
	}
}
