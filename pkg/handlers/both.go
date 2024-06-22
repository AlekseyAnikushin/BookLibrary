package handlers

import (
	"io"
	"net/http"

	services "github.com/AlekseyAnikushin/book_library/pkg/services"
)

func updBookAuthor(res http.ResponseWriter, req *http.Request) {
	ch := make(chan response)
	go func() {
		ba, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			ch <- response{Code: http.StatusBadRequest, Message: "Error reading the request body"}
			return
		}
		if len(ba) == 0 {
			ch <- response{Code: http.StatusBadRequest, Message: "The request body is empty"}
			return
		}

		resultCode, resultMsg := services.UpdateBookAndAuthor(&ba, req.PathValue("book_id"), req.PathValue("author_id"))
		ch <- response{Code: resultCode, Message: resultMsg}
	}()
	resp := <-ch
	writeResponse(res, &resp)
}
