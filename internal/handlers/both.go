package handlers

import (
	"io"
	"net/http"

	"book_library/internal/services"
)

func updBookAuthor(res http.ResponseWriter, req *http.Request) {
	var resp response
	ba, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		resp = response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
		writeResponse(res, &resp)
		return
	}
	if len(ba) == 0 {
		resp = response{Code: http.StatusBadRequest, Message: "The request body is empty"}
		writeResponse(res, &resp)
		return
	}

	resultCode, resultMsg := services.UpdateBookAndAuthor(&ba, req.PathValue("book_id"), req.PathValue("author_id"))
	resp = response{Code: resultCode, Message: resultMsg}
	writeResponse(res, &resp)
}
