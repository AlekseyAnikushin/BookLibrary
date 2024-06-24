package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	entities "entities"
	myerr "my_errors"
	db "storages"
)

func UpdateBookAndAuthor(reqData *[]byte, bookId string, authorId string) (int, string) {
	nb, bErr := strconv.Atoi(bookId)
	if bErr != nil {
		return http.StatusBadRequest, "Invalid book Id: " + bookId
	}

	na, aErr := strconv.Atoi(authorId)
	if aErr != nil {
		return http.StatusBadRequest, "Invalid author Id: " + authorId
	}

	ba := entities.BookAndAuthor{}

	err := json.Unmarshal(*reqData, &ba)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	if v := validBook(&ba.Book); v != "" {
		return http.StatusBadRequest, v
	}

	if v := validAuthor(&ba.Author); v != "" {
		return http.StatusBadRequest, v
	}

	ba.Book.Id = nb
	ba.Author.Id = na

	err = db.UpdBookAndAuthor(&ba)
	if err != nil {
		serr1, ok1 := err.(*myerr.ErrRecordNotFound)
		if ok1 {
			return http.StatusNotFound, serr1.Error()
		}
		serr2, ok2 := err.(*myerr.ErrDatabaseQuery)
		if ok2 {
			return http.StatusInternalServerError, serr2.Error()
		}
		serr3, ok3 := err.(*myerr.ErrAuthorNotFound)
		if ok3 {
			return http.StatusBadRequest, serr3.Error()
		}
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, ""
}
