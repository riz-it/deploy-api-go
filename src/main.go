package main

import (
	"net/http"
	"riz-it/api-go/app"
	"riz-it/api-go/controller"
	"riz-it/api-go/exception"
	"riz-it/api-go/helper"
	"riz-it/api-go/repository"
	"riz-it/api-go/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)

	router := httprouter.New()

	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
