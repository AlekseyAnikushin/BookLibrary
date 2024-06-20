package database

import (
	"database/sql"
	"fmt"
	"strconv"

	entities "github.com/AlekseyAnikushin/book_library/pkg/entities"
	myerr "github.com/AlekseyAnikushin/book_library/pkg/my_errors"
)

func AddBook(b *entities.Book) (int, error) {
	err := checkAuthor(b.AuthorId)
	if err != nil {
		return 0, err
	}

	var newId int
	if len(b.Isbn) > 0 {
		err = db.QueryRow(fmt.Sprintf("INSERT INTO public.Books (\"Title\", \"AuthorID\", \"Year\", \"ISBN\") VALUES('%s', '%d', '%d', '%s') RETURNING \"ID\"", b.Title, b.AuthorId, b.Year, b.Isbn)).Scan(&newId)
	} else {
		err = db.QueryRow(fmt.Sprintf("INSERT INTO public.Books (\"Title\", \"AuthorID\", \"Year\", \"ISBN\") VALUES('%s', '%d', '%d', NULL) RETURNING \"ID\"", b.Title, b.AuthorId, b.Year)).Scan(&newId)
	}
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func GetBooks(id int) ([]entities.Book, error) {
	var Books []entities.Book
	if id > 0 {
		Books = make([]entities.Book, 0, 10)
	} else {
		Books = make([]entities.Book, 0, 1)
	}

	query := "SELECT \"ID\", \"Title\", \"AuthorID\", \"Year\", COALESCE(\"ISBN\",'') AS \"ISBN\" FROM public.Books"
	if id > 0 {
		query += " WHERE \"ID\" = " + strconv.Itoa(id)
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var b entities.Book

	for rows.Next() {
		b = entities.Book{}
		err = rows.Scan(&b.Id, &b.Title, &b.AuthorId, &b.Year, &b.Isbn)
		if err != nil {
			return nil, err
		}
		Books = append(Books, b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return Books, nil
}

func UpdBook(b *entities.Book) error {
	err := checkAuthor(b.AuthorId)
	if err != nil {
		return err
	}

	var res sql.Result
	if len(b.Isbn) > 0 {
		res, err = db.Exec(fmt.Sprintf("UPDATE public.Books SET \"Title\"='%s', \"AuthorID\"='%d', \"Year\"='%d', \"ISBN\"='%s' WHERE \"ID\"=%d", b.Title, b.AuthorId, b.Year, b.Isbn, b.Id))
	} else {
		res, err = db.Exec(fmt.Sprintf("UPDATE public.Books SET \"Title\"='%s', \"AuthorID\"='%d', \"Year\"='%d', \"ISBN\"=NULL WHERE \"ID\"=%d", b.Title, b.AuthorId, b.Year, b.Id))
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

func DelBook(id int) error {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM public.Books WHERE \"ID\"=%d", id))
	if err != nil {
		return err
	}

	i, err2 := res.RowsAffected()
	if err2 == nil && i == 0 {
		return &myerr.ErrRecordNotFound{}
	}

	return nil
}

func checkAuthor(id int) error {
	var a uint64
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(\"ID\") FROM public.Authors WHERE \"ID\"=%d", id)).Scan(&a)
	if err != nil {
		return err
	}
	if a == 0 {
		return &myerr.ErrAuthorNotFound{Id: id}
	}
	return nil
}
