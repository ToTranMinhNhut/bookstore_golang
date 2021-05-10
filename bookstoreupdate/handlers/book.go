package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"bookstoreupdate/models"
	"bookstoreupdate/services/bookservice"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	IBookService bookservice.IBookService
}

func (_self *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	bookController := &models.BookController{}
	if err := json.NewDecoder(r.Body).Decode(&bookController); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validation
	if statusCode, err := bookController.Validate(); err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	bookService := &models.BookService{
		BookId: bookController.BookId,
		Title:  bookController.Title,
		Author: bookController.Author,
		Price:  bookController.Price,
	}

	id, err := _self.IBookService.CreateBook(bookService)
	if err != nil && id == -1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Response
	json.NewEncoder(w).Encode(&models.BookController{
		BookId: id,
	})
}

func (_self *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	BookId := chi.URLParam(r, "bookid")
	id, err := strconv.Atoi(BookId)
	if err != nil {
		log.Println("err: ", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	bookService, err := _self.IBookService.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookController := models.BookController(bookService)

	json.NewEncoder(w).Encode(bookController)
}
