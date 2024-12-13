package repository

import (
	"context"
	"database/sql"
	"errors"
	"riz-it/api-go/helper"
	"riz-it/api-go/model/entity"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book {
	SQL := "insert into book(title) values (?)"
	result, err := tx.ExecContext(ctx, SQL, book.Title)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	book.Id = int(id)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book {
	SQL := "update book set title = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Title, book.Id)
	helper.PanicIfError(err)
	return book
}
func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book entity.Book) {
	SQL := "delete from book where id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	helper.PanicIfError(err)
}
func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (entity.Book, error) {
	SQL := "select id, title from book where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	book := entity.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book is not found")
	}
}
func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Book {
	SQL := "select id, title from book"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		book := entity.Book{}
		err := rows.Scan(&book.Id, &book.Title)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books
}
