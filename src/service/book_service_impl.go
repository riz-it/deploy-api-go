package service

import (
	"context"
	"database/sql"
	"riz-it/api-go/exception"
	"riz-it/api-go/helper"
	"riz-it/api-go/model/entity"
	"riz-it/api-go/model/web"
	"riz-it/api-go/repository"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookService(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := entity.Book{
		Title: request.Title,
	}

	book = service.BookRepository.Save(ctx, tx, book)
	return helper.ToBookResponse(book)

}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book.Title = request.Title

	book = service.BookRepository.Update(ctx, tx, book)
	return helper.ToBookResponse(book)
}
func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, book)
}
func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBookResponse(book)
}
func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAll(ctx, tx)

	return helper.ToBookResponses(books)

}
