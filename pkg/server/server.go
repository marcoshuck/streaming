package server

import (
	"github.com/go-chi/chi"
	"github.com/marcoshuck/streaming/pkg/controller"
	"net/http"
)

type Server struct {
	book   controller.Book
	router *chi.Mux
}

func NewServer(book controller.Book) *Server {
	r := chi.NewRouter()
	return &Server{
		router: r,
		book:   book,
	}
}

func (s *Server) RegisterRoutes() {
	s.router.Get("/books", s.book.GetAll)
	s.router.Get("/books/{isbn}", s.book.GetByISBN)
}

func (s *Server) ListenAndServe(address string) error {
	return http.ListenAndServe(address, s.router)
}
