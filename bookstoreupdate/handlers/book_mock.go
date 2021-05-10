package handlers

import (
	"bookstoreupdate/models"

	"github.com/stretchr/testify/mock"
)

type MockBookService struct {
	mock.Mock
}

func (_self *MockBookService) CreateBook(book *models.BookService) (int, error) {
	args := _self.Called(book)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self *MockBookService) GetBook(BookId int) (models.BookService, error) {
	args := _self.Called(BookId)
	r0 := args.Get(0).(models.BookService)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
