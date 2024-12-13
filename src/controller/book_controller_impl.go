package controller

import (
	"encoding/json"
	"net/http"
	"riz-it/api-go/helper"
	"riz-it/api-go/model/web"
	"riz-it/api-go/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookCreateRequest := web.BookCreateRequest{}
	helper.ReadFromRequestBody(request, &bookCreateRequest)

	bookResponse := controller.BookService.Create(request.Context(), bookCreateRequest)

	webResponse := web.WebResponse{
		Code:    201,
		Status:  true,
		Message: "Book created successfully",
		Data:    bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(request, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	bookResponse := controller.BookService.Update(request.Context(), bookUpdateRequest)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  true,
		Message: "Book updated successfully",
		Data:    bookResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  true,
		Message: "Book deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  true,
		Message: "Show details of the book",
		Data:    bookResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookResponses := controller.BookService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Status:  true,
		Message: "Show all books",
		Data:    bookResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
