package model

import (
	"database/sql"
	"fmt"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/db"
)

type Book struct {
	Title   			sql.NullString
	Description   sql.NullString
	Content   		sql.NullString
}

// Read - read data from database
func Read(BookID int, d db.Connection) (Book, error) {
	o, err := readBook(BookID, d)
	if err != nil {
		return Book{}, err
	}

  return o, nil
}

// Update - update content in database
func Update(BookID int, Content string,  d db.Connection) (Book, error) {
	err := updateBook(BookID, Content, d)
	if err != nil {
		return Book{}, err
	}

  return Read(BookID, d)
}

func readBook(BookID int, d db.Connection) (Book, error) {
	o := Book{}
	r := d.Read(readBookStmnt(BookID))

	err := r.Scan(&o.Title, &o.Description, &o.Content)
	if err != nil {
		return o, err
	}

	return o, nil
}

func updateBook(BookID int, Content string, d db.Connection) error {
	_, err := d.Update(updateContentStmnt(BookID, Content))
	if err != nil {
		return err
	}

	return nil
}

func readBookStmnt(bookID int) string {
	tmplt := "SELECT Title, Description, Content FROM Books WHERE Id = %d"
	return fmt.Sprintf(tmplt, bookID)
}

func updateContentStmnt(bookID int, content string) string {
	tmplt := "UPDATE Books SET Content='%s' WHERE Id=%d"
	return fmt.Sprintf(tmplt, content, bookID)
}

// NewNullString - create null string
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
