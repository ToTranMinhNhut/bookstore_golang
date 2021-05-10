package bookservice

import (
	"bookstoreupdate/models"
	"bookstoreupdate/repository"
)

type IBookService interface {
	CreateBook(book *models.BookService) (int, error)
	GetBook(BookId int) (models.BookService, error)
}

type BookSV struct {
	IBookRepo repository.IBookRepo
}

func (_self *BookSV) CreateBook(bookService *models.BookService) (int, error) {
	bookRepo := &models.BookRepository{
		BookId: bookService.BookId,
		Title:  bookService.Title,
		Author: bookService.Author,
		Price:  bookService.Price,
	}

	var id int
	id, err := _self.IBookRepo.CreateBook(bookRepo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (_self *BookSV) GetBook(BookId int) (models.BookService, error) {
	bookRepo, err := _self.IBookRepo.GetBook(BookId)
	bookService := models.BookService(bookRepo)
	return bookService, err
}
