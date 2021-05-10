package bookservice

import (
	"errors"
	"testing"

	"bookstoreupdate/models"

	"github.com/stretchr/testify/require"
)

func TestBookService_CreateBook(t *testing.T) {
	testCases := []struct {
		name           string
		input          *models.BookService
		expectedErr    error
		expectedResult int
		mockRepoInput  *models.BookRepository
		mockRepoResult int
		mockRepoErr    error
	}{
		{
			name: "Create book failed with error",
			input: &models.BookService{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			expectedErr:    errors.New("Create book failed with error"),
			expectedResult: -1,
			mockRepoInput: &models.BookRepository{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			mockRepoResult: -1,
			mockRepoErr:    errors.New("Create book failed with error"),
		},
		{
			name: "Create booksuccess",
			input: &models.BookService{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			expectedErr:    nil,
			expectedResult: 99,
			mockRepoInput: &models.BookRepository{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			mockRepoResult: 99,
			mockRepoErr:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Given
			mockBookRepo := new(MockBookRepo)
			mockBookRepo.On("CreateBook", tc.mockRepoInput).
				Return(tc.mockRepoResult, tc.mockRepoErr)

			bookService := BookSV{
				IBookRepo: mockBookRepo,
			}

			//When
			id, err := bookService.CreateBook(tc.input)

			//Then
			if tc.expectedErr != nil && id != -1 {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBookService_GetBook(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedErr    error
		expectedResult models.BookService
		mockRepoInput  int
		mockRepoResult models.BookRepository
		mockRepoErr    error
	}{
		{
			name:           "Get failed with error",
			input:          100,
			expectedErr:    errors.New("get failed with error"),
			expectedResult: models.BookService{},
			mockRepoInput:  100,
			mockRepoResult: models.BookRepository{},
			mockRepoErr:    errors.New("get failed with error"),
		},
		{
			name:        "Check existed success",
			input:       99,
			expectedErr: nil,
			expectedResult: models.BookService{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			mockRepoInput: 99,
			mockRepoResult: models.BookRepository{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			mockRepoErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Given
			mockBookRepo := new(MockBookRepo)
			mockBookRepo.On("GetBook", tc.mockRepoInput).
				Return(tc.mockRepoResult, tc.mockRepoErr)
			service := BookSV{
				IBookRepo: mockBookRepo,
			}

			//When
			existed, err := service.GetBook(tc.input)

			//Then
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResult, existed)
			}
		})
	}
}
