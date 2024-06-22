package handlers

import (
	"io"
	"net/http"

	services "github.com/AlekseyAnikushin/book_library/pkg/services"
)

func addBook(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		b, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			ch <- response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
			return
		}
		if len(b) == 0 {
			ch <- response{Code: http.StatusBadRequest, Message: "The request body is empty"}
			return
		}

		resultCode, resultMsg := services.AddBook(&b)
		ch <- response{Code: resultCode, Message: resultMsg}

	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func getBooks(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		resultCode, resultMsg, resultData := services.GetBooks()
		ch <- response{Code: resultCode, Message: resultMsg, Result: resultData}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func getBook(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		resultCode, resultMsg, resultData := services.GetBook(req.PathValue("id"))
		ch <- response{Code: resultCode, Message: resultMsg, Result: resultData}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func updBook(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		b, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			ch <- response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
			return
		}
		if len(b) == 0 {
			ch <- response{Code: http.StatusBadRequest, Message: "The request body is empty"}
			return
		}

		resultCode, resultMsg := services.UpdateBook(req.PathValue("id"), &b)
		ch <- response{Code: resultCode, Message: resultMsg}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func delBook(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		resultCode, resultMsg := services.DeleteBook(req.PathValue("id"))
		ch <- response{Code: resultCode, Message: resultMsg}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}
