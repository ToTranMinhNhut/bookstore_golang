package bookservice

import (
	"bookstoreupdate/models"

	"github.com/stretchr/testify/mock"
)

type MockBookRepo struct {
	mock.Mock
}

func (_self *MockBookRepo) CreateBook(book *models.BookRepository) (int, error) {
	args := _self.Called(book)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self *MockBookRepo) GetBook(BookId int) (models.BookRepository, error) {
	args := _self.Called(BookId)
	r0 := args.Get(0).(models.BookRepository)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
