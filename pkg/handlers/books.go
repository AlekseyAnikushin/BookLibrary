package handlers

import (
	"io"
	"net/http"

	services "booklib/pkg/services"
)

func addBook(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		writeResponse(res, "Error reading the request body", http.StatusBadRequest, nil)
		return
	}
	if len(b) == 0 {
		writeResponse(res, "The request body is empty", http.StatusBadRequest, nil)
		return
	}

	resultCode, resultMsg := services.AddBook(&b)
	writeResponse(res, resultMsg, resultCode, nil)
}

func getBooks(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetBooks()
	writeResponse(res, resultMsg, resultCode, resultData)
}

func getBook(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetBook(req.PathValue("id"))
	writeResponse(res, resultMsg, resultCode, resultData)
}

func updBook(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		writeResponse(res, "Error reading the request body", http.StatusBadRequest, nil)
		return
	}
	if len(b) == 0 {
		writeResponse(res, "The request body is empty", http.StatusBadRequest, nil)
		return
	}

	resultCode, resultMsg := services.UpdateBook(req.PathValue("id"), &b)
	writeResponse(res, resultMsg, resultCode, nil)
}

func delBook(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg := services.DeleteBook(req.PathValue("id"))
	writeResponse(res, resultMsg, resultCode, nil)
}
