package handlers

import (
	"io"
	"net/http"

	services "booklib/pkg/services"
)

func addAuthor(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	a, err := io.ReadAll(req.Body)
	if err != nil {
		writeResponse(res, "Error reading the request body", http.StatusBadRequest, nil)
		return
	}
	if len(a) == 0 {
		writeResponse(res, "The request body is empty", http.StatusBadRequest, nil)
		return
	}

	resultCode, resultMsg := services.AddAuthor(&a)
	writeResponse(res, resultMsg, resultCode, nil)
}

func getAuthors(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetAuthors()
	writeResponse(res, resultMsg, resultCode, resultData)
}

func getAuthor(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg, resultData := services.GetAuthor(req.PathValue("id"))
	writeResponse(res, resultMsg, resultCode, resultData)
}

func updAuthor(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	a, err := io.ReadAll(req.Body)
	if err != nil {
		writeResponse(res, "Error reading the request body", http.StatusBadRequest, nil)
		return
	}
	if len(a) == 0 {
		writeResponse(res, "The request body is empty", http.StatusBadRequest, nil)
		return
	}

	resultCode, resultMsg := services.UpdateAuthor(req.PathValue("id"), &a)
	writeResponse(res, resultMsg, resultCode, nil)
}

func delAuthor(res http.ResponseWriter, req *http.Request) {
	resultCode, resultMsg := services.DeleteAuthor(req.PathValue("id"))
	writeResponse(res, resultMsg, resultCode, nil)
}
