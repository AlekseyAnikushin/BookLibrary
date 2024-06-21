package database

import (
	"database/sql"
	"fmt"
	"strconv"

	entities "booklib/pkg/entities"
	myerr "booklib/pkg/my_errors"
)

func AddAuthor(a *entities.Author) (int, error) {
	var newId int
	err := db.QueryRow(fmt.Sprintf("INSERT INTO public.authors (\"Name\", \"Surname\", \"Biography\", \"Birthdate\") VALUES('%s', '%s', '%s', '%s') RETURNING \"ID\"", a.Name, a.Surname, a.Biography, a.Birthdate)).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func GetAuthors(id int) ([]entities.Author, error) {
	var authors []entities.Author
	if id > 0 {
		authors = make([]entities.Author, 0, 10)
	} else {
		authors = make([]entities.Author, 0, 1)
	}

	query := "SELECT \"ID\", \"Name\", \"Surname\", COALESCE(\"Biography\",'') AS \"Biography\", to_char(\"Birthdate\", 'YYYY-MM-DD') FROM public.authors"
	if id > 0 {
		query += " WHERE \"ID\" = " + strconv.Itoa(id)
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var a entities.Author

	for rows.Next() {
		a = entities.Author{}
		err = rows.Scan(&a.Id, &a.Name, &a.Surname, &a.Biography, &a.Birthdate)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func UpdAuthor(a *entities.Author, tx *sql.Tx) error {
	var res sql.Result
	var err error

	query := fmt.Sprintf("UPDATE public.authors SET \"Name\"='%s', \"Surname\"='%s', \"Biography\"='%s', \"Birthdate\"='%s' WHERE \"ID\"=%d", a.Name, a.Surname, a.Biography, a.Birthdate, a.Id)
	if tx == nil {
		res, err = db.Exec(query)
	} else {
		res, err = tx.Exec(query)
	}
	if err != nil {
		return err
	}

	i, err2 := res.RowsAffected()
	if err2 == nil && i == 0 {
		return &myerr.ErrRecordNotFound{}
	}

	return nil
}

func DelAuthor(id int) error {
	var bookCount int
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(\"ID\") FROM public.books WHERE \"AuthorID\"=%d", id)).Scan(&bookCount)
	if err != nil {
		return err
	}

	if bookCount > 0 {
		return &myerr.ErrBooksExist{BookCount: bookCount}
	}

	res, err := db.Exec(fmt.Sprintf("DELETE FROM public.authors WHERE \"ID\"=%d", id))
	if err != nil {
		return err
	}

	i, err2 := res.RowsAffected()
	if err2 == nil && i == 0 {
		return &myerr.ErrRecordNotFound{}
	}

	return nil
}
