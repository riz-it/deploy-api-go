package exception

import (
	"net/http"
	"riz-it/api-go/helper"
	"riz-it/api-go/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  false,
			Message: "Book not found",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	response := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Status:  false,
		Message: "Internal Server error",
		Data:    err,
	}
	helper.WriteToResponseBody(writer, response)
}
