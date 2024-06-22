package handlers

import (
	"io"
	"net/http"

	services "github.com/AlekseyAnikushin/book_library/pkg/services"
)

func addAuthor(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		a, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			ch <- response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
			return
		}
		if len(a) == 0 {
			ch <- response{Code: http.StatusBadRequest, Message: "The request body is empty"}
			return
		}

		resultCode, resultMsg := services.AddAuthor(&a)
		ch <- response{Code: resultCode, Message: resultMsg}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func getAuthors(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		resultCode, resultMsg, resultData := services.GetAuthors()
		ch <- response{Code: resultCode, Message: resultMsg, Result: resultData}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func getAuthor(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		resultCode, resultMsg, resultData := services.GetAuthor(req.PathValue("id"))
		ch <- response{Code: resultCode, Message: resultMsg, Result: resultData}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func updAuthor(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		a, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			ch <- response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
			return
		}
		if len(a) == 0 {
			ch <- response{Code: http.StatusBadRequest, Message: "The request body is empty"}
			return
		}

		resultCode, resultMsg := services.UpdateAuthor(req.PathValue("id"), &a)
		ch <- response{Code: resultCode, Message: resultMsg}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}

func delAuthor(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		resultCode, resultMsg := services.DeleteAuthor(req.PathValue("id"))
		ch <- response{Code: resultCode, Message: resultMsg}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}
