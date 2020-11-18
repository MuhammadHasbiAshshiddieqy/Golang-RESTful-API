package model

import (
	"database/sql"
	"strconv"
	"fmt"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/db"
)

type Book struct {
	Title   			sql.NullString 
	Description   sql.NullString
	Content   		sql.NullString
}

// Read - read data from database
func Read(BookID string, d db.Connection) (Book, error) {
	o, err := readBook(BookID, d)
	if err != nil {
		return Book{}, err
	}

  return o, nil
}

func readBook(BookID string, d db.Connection) (Book, error) {
	o := Book{}
	id, _ := strconv.Atoi(BookID)
	r := d.Read(bookStmnt(id))

	err := r.Scan(&o.Title, &o.Description, &o.Content)
	if err != nil {
		return o, err
	}

	return o, nil
}

func bookStmnt(bookID int) string {
	tmplt := "SELECT Title, Description, Content FROM Books WHERE Id = %d"
	return fmt.Sprintf(tmplt, bookID)
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
