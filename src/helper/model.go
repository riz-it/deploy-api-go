package helper

import (
	"riz-it/api-go/model/entity"
	"riz-it/api-go/model/web"
)

func ToBookResponse(book entity.Book) web.BookResponse {
	return web.BookResponse{
		Id:    book.Id,
		Title: book.Title,
	}
}

func ToBookResponses(books []entity.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}
