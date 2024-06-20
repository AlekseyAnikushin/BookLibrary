package my_errors

import "fmt"

type ErrBooksExist struct {
	BookCount int
}

func (r *ErrBooksExist) Error() string {
	return fmt.Sprintf("there are %d books by this author", r.BookCount)
}

type ErrRecordNotFound struct{}

func (r *ErrRecordNotFound) Error() string {
	return "record not found"
}

type ErrDatabaseQuery struct{}

func (r *ErrDatabaseQuery) Error() string {
	return "database query error"
}

type ErrAuthorNotFound struct {
	Id int
}

func (r *ErrAuthorNotFound) Error() string {
	return fmt.Sprintf("the author with id %d not found", r.Id)
}
