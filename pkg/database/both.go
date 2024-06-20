package database

import (
	entities "github.com/AlekseyAnikushin/book_library/pkg/entities"
)

func UpdBookAndAuthor(ba *entities.BookAndAuthor) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = UpdBook(&ba.Book, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = UpdAuthor(&ba.Author, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
