package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"bookstoreupdate/db"
	"bookstoreupdate/models"
)

type IBookRepo interface {
	CreateBook(book *models.BookRepository) (int, error)
	GetBook(BookId int) (models.BookRepository, error)
}

type BookRepo struct {
	Conn *db.DB
}

func (_self *BookRepo) CreateBook(book *models.BookRepository) (int, error) {
	ctx := context.Background()
	tx, err := _self.Conn.Client.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id int
	query := `INSERT INTO book (bookid, title, author, price) VALUES ($1, $2, $3, $4) RETURNING bookid`
	err = tx.QueryRowContext(ctx, query, book.BookId, book.Title, book.Author, book.Price).Scan(&id)
	if err != nil {
		return -1, err
	}

	tx.Commit()
	return id, nil
}

func (_self *BookRepo) GetBook(BookId int) (models.BookRepository, error) {
	ctx := context.Background()
	tx, err := _self.Conn.Client.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	book := models.BookRepository{}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT * FROM book WHERE bookid = $1;`
	row := tx.QueryRowContext(ctx, query, BookId)
	switch err := row.Scan(&book.BookId, &book.Title, &book.Author, &book.Price); err {
	case sql.ErrNoRows:
		return book, err
	default:
		return book, err
	}
}
