package models

import (
	"errors"
	"net/http"
)

// use for handler
type BookController struct {
	BookId int    `json:"bookid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

func (_self *BookController) Validate() (int, error) {
	if _self.BookId < 0 {
		return http.StatusBadRequest, errors.New("\"bookid\" field is greater than 0")
	}
	if _self.Title == "" {
		return http.StatusBadRequest, errors.New("\"title\" field is required")
	}
	if _self.Author == "" {
		return http.StatusBadRequest, errors.New("\"author\" field is required")
	}
	if _self.Price < 0 {
		return http.StatusBadRequest, errors.New("\"price\" field is greater than 0")
	}
	return http.StatusOK, nil
}

type BookService struct {
	BookId int    `json:"bookid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

type BookRepository struct {
	BookId int    `json:"bookid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

type SuccessResponse struct {
	Success bool `json:"Success"`
}
