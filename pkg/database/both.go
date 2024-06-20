package database

import (
	"fmt"

	entities "github.com/AlekseyAnikushin/book_library/pkg/entities"
	myerr "github.com/AlekseyAnikushin/book_library/pkg/my_errors"
)

func UpdBookAndAuthor(ba *entities.BookAndAuthor) error {
	var result int
	var err error
	if ba.Book.Isbn != "" {
		err = db.QueryRow(fmt.Sprintf(
			"SELECT  public.UpdateBookAndAuthor ("+
				"%d,"+ //.................................... book_id
				"CAST('%s' AS character varying(500)),"+ //.. book_title
				"%d,"+ //.................................... book_author_id
				"CAST(%d AS smallint),"+ //.................. book_year
				"CAST('%s' AS character(13)),"+ //............book_isbn
				"%d,"+ //.................................... author_id
				"CAST('%s' AS character varying(50)),"+ //... author_name
				"CAST('%s' AS character varying(50)),"+ //... author_surname
				"CAST('%s' AS text),"+ //.................... author_biorgaphy
				"CAST('%s' AS date))", //.................... author_birthdate
			ba.Book.Id,
			ba.Book.Title,
			ba.Book.AuthorId,
			ba.Book.Year,
			ba.Book.Isbn,
			ba.Author.Id,
			ba.Author.Name,
			ba.Author.Surname,
			ba.Author.Biography,
			ba.Author.Birthdate)).Scan(&result)
	} else {
		err = db.QueryRow(fmt.Sprintf(
			"SELECT  public.UpdateBookAndAuthor ("+
				"%d,"+ //.................................... book_id
				"CAST('%s' AS character varying(500)),"+ //.. book_title
				"%d,"+ //.................................... book_author_id
				"CAST(%d AS smallint),"+ //.................. book_year
				"NULL,"+ //.................................. book_isbn
				"%d,"+ //.................................... author_id
				"CAST('%s' AS character varying(50)),"+ //... author_name
				"CAST('%s' AS character varying(50)),"+ //... author_surname
				"CAST('%s' AS text),"+ //.................... author_biorgaphy
				"CAST('%s' AS date))", //.................... author_birthdate
			ba.Book.Id,
			ba.Book.Title,
			ba.Book.AuthorId,
			ba.Book.Year,
			ba.Author.Id,
			ba.Author.Name,
			ba.Author.Surname,
			ba.Author.Biography,
			ba.Author.Birthdate)).Scan(&result)
	}

	if err != nil {
		return err
	}

	if result == -1 {
		return &myerr.ErrDatabaseQuery{}
	}

	if result > 0 {
		return &myerr.ErrRecordNotFound{}
	}

	return nil
}
