package repository

import (
	"context"
	"database/sql"
	"riz-it/api-go/model/entity"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book
	Update(ctx context.Context, tx *sql.Tx, book entity.Book) entity.Book
	Delete(ctx context.Context, tx *sql.Tx, book entity.Book)
	FindById(ctx context.Context, tx *sql.Tx, bookId int) (entity.Book, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Book
}
