package main

import (
	"github.com/marcoshuck/streaming/pkg/controller"
	"github.com/marcoshuck/streaming/pkg/model"
	"github.com/marcoshuck/streaming/pkg/repository"
	"github.com/marcoshuck/streaming/pkg/server"
	"github.com/marcoshuck/streaming/pkg/service"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err)
	}

	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatalln("Error:", err)
	}

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	s := server.NewServer(bookController)

	s.RegisterRoutes()

	log.Println("Running server on port :3000")

	err = s.ListenAndServe(":3000")
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
