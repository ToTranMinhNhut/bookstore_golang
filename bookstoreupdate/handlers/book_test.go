package handlers

import (
	"bookstoreupdate/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_CreateUser(t *testing.T) {
	type mockCreateBookService struct {
		input  *models.BookService
		BookID int
		err    error
	}
	testCases := []struct {
		name                   string
		requestBody            interface{}
		expectedResponseBody   string
		expectedResponseStatus int
		mockCreateBookService  mockCreateBookService
	}{
		{
			name: "Validate request body failed",
			requestBody: map[string]interface{}{
				"bookid": 0,
				"title":  "",
				"author": "",
				"price":  0,
			},
			expectedResponseBody:   "\"title\" field is required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name: "Call services return with error",
			requestBody: map[string]interface{}{
				"bookid": 99,
				"title":  "Golang",
				"author": "raa",
				"price":  600,
			},
			expectedResponseBody:   "services error\n",
			expectedResponseStatus: http.StatusBadRequest,
			mockCreateBookService: mockCreateBookService{
				input: &models.BookService{
					BookId: 99,
					Title:  "Golang",
					Author: "raa",
					Price:  600,
				},
				BookID: -1,
				err:    errors.New("services error"),
			},
		},
		{
			name: "Call services success",
			requestBody: map[string]interface{}{
				"bookid": 99,
				"title":  "Golang",
				"author": "raa",
				"price":  600,
			},
			expectedResponseBody:   "{\"bookid\":99,\"title\":\"\",\"author\":\"\",\"price\":0}\n",
			expectedResponseStatus: http.StatusOK,
			mockCreateBookService: mockCreateBookService{
				input: &models.BookService{
					BookId: 99,
					Title:  "Golang",
					Author: "raa",
					Price:  600,
				},
				BookID: 99,
				err:    nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Given
			mockService := new(MockBookService)

			mockService.On("CreateBook", tc.mockCreateBookService.input).
				Return(tc.mockCreateBookService.BookID, tc.mockCreateBookService.err)

			handlers := BookHandler{
				IBookService: mockService,
			}

			requestBody, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Error(err)
			}

			//When
			req, err := http.NewRequest(http.MethodPost, "/books/create", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateBook)
			handler.ServeHTTP(responseRecorder, req)

			//Then
			require.Equal(t, tc.expectedResponseStatus, responseRecorder.Code)
			require.Equal(t, tc.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}

func TestBookHandler_GetBook(t *testing.T) {
	type mockGetBookByID struct {
		input  int
		result models.BookService
		err    error
	}
	testCases := []struct {
		name                 string
		input                string
		expectedResponseBody string
		expectedStatus       int
		mockGetBookByID      mockGetBookByID
	}{
		{
			name:                 "Get Book failed with error",
			input:                "1",
			expectedResponseBody: "get Book failed with error\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetBookByID: mockGetBookByID{
				input:  1,
				result: models.BookService{},
				err:    errors.New("get Book failed with error"),
			},
		},
		{
			name:                 "Get Book return not exist",
			input:                "55",
			expectedResponseBody: "book does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetBookByID: mockGetBookByID{
				input:  55,
				result: models.BookService{},
				err:    errors.New("book does not exist"),
			},
		},
		{
			name:                 "Success request",
			input:                "99",
			expectedResponseBody: "{\"bookid\":99,\"title\":\"Golang\",\"author\":\"raa\",\"price\":600}\n",
			expectedStatus:       http.StatusOK,
			mockGetBookByID: mockGetBookByID{
				input: 99,
				result: models.BookService{
					BookId: 99,
					Title:  "Golang",
					Author: "raa",
					Price:  600,
				},
				err: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Given
			mockBookService := new(MockBookService)

			mockBookService.On("GetBook", tc.mockGetBookByID.input).
				Return(tc.mockGetBookByID.result, tc.mockGetBookByID.err)

			handlers := BookHandler{
				IBookService: mockBookService,
			}

			r, err := http.NewRequest(http.MethodGet, "/books/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("bookid", tc.input)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetBook)
			handler.ServeHTTP(w, r)

			//Then
			require.Equal(t, tc.expectedStatus, w.Code)
			require.Equal(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
