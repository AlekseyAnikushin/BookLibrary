package handlers

import (
	"io"
	"net/http"

	services "booklib/pkg/services"
)

func updBookAuthor(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	ba, err := io.ReadAll(req.Body)
	if err != nil {
		writeResponse(res, "Error reading the request body", http.StatusBadRequest, nil)
		return
	}
	if len(ba) == 0 {
		writeResponse(res, "The request body is empty", http.StatusBadRequest, nil)
		return
	}

	resultCode, resultMsg := services.UpdateBookAndAuthor(&ba, req.PathValue("book_id"), req.PathValue("author_id"))
	writeResponse(res, resultMsg, resultCode, nil)
}
