package services

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	db "github.com/AlekseyAnikushin/book_library/pkg/database"
	entities "github.com/AlekseyAnikushin/book_library/pkg/entities"
	myerr "github.com/AlekseyAnikushin/book_library/pkg/my_errors"
)

func GetBooks() (int, string, []byte) {
	books, err := db.GetBooks(0)
	if err != nil {
		return http.StatusInternalServerError, "Database error", nil
	}

	if len(books) == 0 {
		e := myerr.ErrRecordNotFound{}
		return http.StatusNotFound, e.Error(), nil
	}

	res, jsonErr := json.Marshal(books)
	if jsonErr != nil {
		return http.StatusInternalServerError, "Internal error", nil
	}

	return http.StatusOK, "", res
}

func GetBook(id string) (int, string, []byte) {
	n, convErr := strconv.Atoi(id)
	if convErr != nil {
		return http.StatusBadRequest, "Invalid book Id: " + id, nil
	}

	books, err := db.GetBooks(n)
	if err != nil {
		return http.StatusInternalServerError, "Database error", nil
	}

	if len(books) == 0 {
		e := myerr.ErrRecordNotFound{}
		return http.StatusNotFound, e.Error(), nil
	}

	res, jsonErr := json.Marshal(books)
	if jsonErr != nil {
		return http.StatusInternalServerError, "Internal error", nil
	}

	return http.StatusOK, "", res
}

func DeleteBook(id string) (int, string) {
	n, convErr := strconv.Atoi(id)
	if convErr != nil {
		return http.StatusBadRequest, "Invalid book Id: " + id
	}

	err := db.DelBook(n)
	if err != nil {
		serr, ok := err.(*myerr.ErrRecordNotFound)
		if ok {
			return http.StatusNotFound, serr.Error()
		}
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, ""
}

func UpdateBook(id string, reqData *[]byte) (int, string) {
	book := entities.Book{}

	n, convErr := strconv.Atoi(id)
	if convErr != nil {
		return http.StatusBadRequest, "Invalid book Id: " + id
	}

	var err error
	err = json.Unmarshal(*reqData, &book)

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	if v := validBook(&book); v != "" {
		return http.StatusBadRequest, v
	}

	book.Id = n

	err = db.UpdBook(&book, nil)
	if err != nil {
		serr1, ok1 := err.(*myerr.ErrRecordNotFound)
		if ok1 {
			return http.StatusNotFound, serr1.Error()
		}
		serr2, ok2 := err.(*myerr.ErrAuthorNotFound)
		if ok2 {
			return http.StatusBadRequest, serr2.Error()
		}
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, ""
}

func AddBook(reqData *[]byte) (int, string) {
	book := entities.Book{}

	var err error
	err = json.Unmarshal(*reqData, &book)

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	if v := validBook(&book); v != "" {
		return http.StatusBadRequest, v
	}

	var id int
	id, err = db.AddBook(&book)

	if err != nil {
		serr, ok := err.(*myerr.ErrAuthorNotFound)
		if ok {
			return http.StatusBadRequest, serr.Error()
		}
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, "Book Id: " + strconv.Itoa(id)
}

func validBook(b *entities.Book) string {
	if b.Title == "" {
		return "The book's title is missing"
	}
	if b.AuthorId == 0 {
		return "The books's author id is missing"
	}
	if b.Year == 0 || int(b.Year) > time.Now().Year() {
		return "Invalid year value"
	}
	if len(b.Isbn) > 0 && len(b.Isbn) != 13 {
		return "The ISBN length is incorrect, it should be 13"
	}
	return ""
}
