package handlers

import (
	"io"
	"net/http"

	"book_library/internal/services"
)

func addAuthor(res http.ResponseWriter, req *http.Request) {
	var resp response
	a, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		resp = response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
		writeResponse(res, &resp)
		return
	}
	if len(a) == 0 {
		resp = response{Code: http.StatusBadRequest, Message: "The request body is empty"}
		writeResponse(res, &resp)
		return
	}

	resultCode, resultMsg := services.AddAuthor(&a)
	resp = response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}

func getAuthors(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetAuthors()
	resp := response{Code: resultCode, Message: resultMsg, Result: resultData}
	writeResponse(res, &resp)
}

func getAuthor(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetAuthor(req.PathValue("id"))
	resp := response{Code: resultCode, Message: resultMsg, Result: resultData}
	writeResponse(res, &resp)
}

func updAuthor(res http.ResponseWriter, req *http.Request) {
	var resp response
	a, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		resp = response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
		writeResponse(res, &resp)
		return
	}
	if len(a) == 0 {
		resp = response{Code: http.StatusBadRequest, Message: "The request body is empty"}
		writeResponse(res, &resp)
		return
	}

	resultCode, resultMsg := services.UpdateAuthor(req.PathValue("id"), &a)
	resp = response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}

func delAuthor(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg := services.DeleteAuthor(req.PathValue("id"))
	resp := response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}
