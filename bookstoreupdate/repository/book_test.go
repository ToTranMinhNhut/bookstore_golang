package repository

import (
	"bookstoreupdate/db"
	"bookstoreupdate/models"
	"bookstoreupdate/testhelpers"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBookRepo_CreateBook(t *testing.T) {
	testCases := []struct {
		name           string
		input          *models.BookRepository
		expectedResult int
		expectedErr    error
		mockDB         *db.DB
	}{
		{
			name: "Create user success",
			input: &models.BookRepository{
				BookId: 50,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			expectedResult: 99,
			expectedErr:    nil,
			mockDB:         testhelpers.ConnectDB(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			dbMock := tc.mockDB

			bookRepo := BookRepo{
				Conn: dbMock,
			}

			defer bookRepo.Conn.Close()
			// When
			id, err := bookRepo.CreateBook(tc.input)

			// Then
			if tc.expectedErr != nil && id != -1 {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserRepo_GetBook(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult models.BookRepository
		expectedErr    error
		preparePath    string
		mockDb         *db.DB
	}{
		{
			name:           "The user does not exist",
			input:          80,
			expectedResult: models.BookRepository{},
			expectedErr:    errors.New("sql: no rows in result set"),
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
		{
			name:  "Get Book by ID success",
			input: 99,
			expectedResult: models.BookRepository{
				BookId: 99,
				Title:  "Golang",
				Author: "raa",
				Price:  600,
			},
			expectedErr: nil,
			mockDb:      testhelpers.ConnectDB(),
			preparePath: "../testhelpers/preparedata/datafortest",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(tc.mockDb, tc.preparePath)

			bookRepo := BookRepo{
				Conn: tc.mockDb,
			}

			// When
			result, err := bookRepo.GetBook(tc.input)

			// Then
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, tc.expectedResult)
			}
		})
	}
}
