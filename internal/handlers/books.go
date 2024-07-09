package handlers

import (
	"io"
	"net/http"

	"book_library/internal/services"
)

func addBook(res http.ResponseWriter, req *http.Request) {
	var resp response
	b, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		resp = response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
		writeResponse(res, &resp)
		return
	}
	if len(b) == 0 {
		resp = response{Code: http.StatusBadRequest, Message: "The request body is empty"}
		writeResponse(res, &resp)
		return
	}

	resultCode, resultMsg := services.AddBook(&b)
	resp = response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}

func getBooks(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetBooks()
	resp := response{Code: resultCode, Message: resultMsg, Result: resultData}
	writeResponse(res, &resp)
}

func getBook(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetBook(req.PathValue("id"))
	resp := response{Code: resultCode, Message: resultMsg, Result: resultData}
	writeResponse(res, &resp)
}

func updBook(res http.ResponseWriter, req *http.Request) {
	var resp response
	b, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		resp = response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
		writeResponse(res, &resp)
		return
	}
	if len(b) == 0 {
		resp = response{Code: http.StatusBadRequest, Message: "The request body is empty"}
		writeResponse(res, &resp)
		return
	}

	resultCode, resultMsg := services.UpdateBook(req.PathValue("id"), &b)
	resp = response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}

func delBook(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg := services.DeleteBook(req.PathValue("id"))
	resp := response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}
