package services

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	entities "entities"
	myerr "my_errors"
	db "storages"
)

func GetAuthors() (int, string, []byte) {
	authors, err := db.GetAuthors(0)
	if err != nil {
		return http.StatusInternalServerError, "Database error", nil
	}

	if len(authors) == 0 {
		e := myerr.ErrRecordNotFound{}
		return http.StatusNotFound, e.Error(), nil
	}

	res, jsonErr := json.Marshal(authors)
	if jsonErr != nil {
		return http.StatusInternalServerError, "Internal error", nil
	}

	return http.StatusOK, "", res
}

func GetAuthor(id string) (int, string, []byte) {
	n, convErr := strconv.Atoi(id)
	if convErr != nil {
		return http.StatusBadRequest, "Invalid author Id: " + id, nil
	}

	authors, err := db.GetAuthors(n)
	if err != nil {
		return http.StatusInternalServerError, "Database error", nil
	}

	if len(authors) == 0 {
		e := myerr.ErrRecordNotFound{}
		return http.StatusNotFound, e.Error(), nil
	}

	res, jsonErr := json.Marshal(authors)
	if jsonErr != nil {
		return http.StatusInternalServerError, "Internal error", nil
	}

	return http.StatusOK, "", res
}

func DeleteAuthor(id string) (int, string) {
	n, convErr := strconv.Atoi(id)
	if convErr != nil {
		return http.StatusBadRequest, "Invalid author Id: " + id
	}

	err := db.DelAuthor(n)
	if err != nil {
		serr1, ok1 := err.(*myerr.ErrRecordNotFound)
		if ok1 {
			return http.StatusNotFound, serr1.Error()
		}
		serr2, ok2 := err.(*myerr.ErrBooksExist)
		if ok2 {
			return http.StatusConflict, serr2.Error()
		}
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, ""
}

func UpdateAuthor(id string, reqData *[]byte) (int, string) {
	author := entities.Author{}

	n, convErr := strconv.Atoi(id)
	if convErr != nil {
		return http.StatusBadRequest, "Invalid author Id: " + id
	}

	var err error
	err = json.Unmarshal(*reqData, &author)

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	if v := validAuthor(&author); v != "" {
		return http.StatusBadRequest, v
	}

	author.Id = n

	err = db.UpdAuthor(&author, nil)
	if err != nil {
		serr, ok := err.(*myerr.ErrRecordNotFound)
		if ok {
			return http.StatusNotFound, serr.Error()
		}
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, ""
}

func AddAuthor(reqData *[]byte) (int, string) {
	author := entities.Author{}

	var err error
	err = json.Unmarshal(*reqData, &author)

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	if v := validAuthor(&author); v != "" {
		return http.StatusBadRequest, v
	}

	var id int
	id, err = db.AddAuthor(&author)

	if err != nil {
		return http.StatusInternalServerError, "Database error"
	}

	return http.StatusOK, "Author Id: " + strconv.Itoa(id)
}

func validAuthor(a *entities.Author) string {
	if a.Surname == "" {
		return "The author's surname is missing"
	}
	if a.Name == "" {
		return "The author's name is missing"
	}
	if _, e := time.Parse("2006-01-02", a.Birthdate); e != nil {
		return "Invalid authors's birthdate - " + a.Birthdate
	}
	return ""
}
