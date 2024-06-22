package handlers

import (
	"net/http"
)

type response struct {
	Code    int
	Message string
	Result  []byte
}

var mux *http.ServeMux

func GetMux() *http.ServeMux {
	mux = http.NewServeMux()

	mux.HandleFunc("POST /authors", addAuthor)
	mux.HandleFunc("GET /authors", getAuthors)
	mux.HandleFunc("GET /authors/{id}", getAuthor)
	mux.HandleFunc("PUT /authors/{id}", updAuthor)
	mux.HandleFunc("DELETE /authors/{id}", delAuthor)

	mux.HandleFunc("POST /books", addBook)
	mux.HandleFunc("GET /books", getBooks)
	mux.HandleFunc("GET /books/{id}", getBook)
	mux.HandleFunc("PUT /books/{id}", updBook)
	mux.HandleFunc("DELETE /books/{id}", delBook)

	mux.HandleFunc("PUT /books/{book_id}/authors/{author_id}", updBookAuthor)

	return mux
}
