package routes

import (
	"bookstoreupdate/db"
	"bookstoreupdate/handlers"
	"bookstoreupdate/repository"
	"bookstoreupdate/services/bookservice"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(db *db.DB) *chi.Mux {
	mux := chi.NewRouter()

	handler := handlers.BookHandler{
		IBookService: &bookservice.BookSV{
			IBookRepo: &repository.BookRepo{
				Conn: db,
			},
		},
	}

	mux.Route("/books", func(mux chi.Router) {
		mux.Post("/create", handler.CreateBook)
		mux.Get("/{bookid}", handler.GetBook)
	})
	return mux
}
