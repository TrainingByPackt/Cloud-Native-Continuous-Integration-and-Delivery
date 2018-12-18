package books

import (
	"database/sql"

	"github.com/xo/dburl"
)

type SQLBookDatabase struct {
	db *sql.DB
}

func NewSQLBookDatabase(dbAddress string) (BookDatabase, error) {

	db, err := dburl.Open(dbAddress)
	if err != nil {
		return nil, err
	}

	sdb := SQLBookDatabase{db: db}
	return sdb, nil

}

func (sbd SQLBookDatabase) GetBooks() ([]Book, error) {

	books := make([]Book, 0)

	rows, err := sbd.db.Query(`SELECT * FROM books`)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.ISBN, &book.Title, &book.Author)
		if err != nil {
			return books, err

		}
		books = append(books, book)
	}

	return books, nil
}

func (sbd SQLBookDatabase) Initialize() error {

	var schema = `CREATE TABLE books (isbn text, title text, author text);`

	_, err := sbd.db.Exec(schema)
	if err != nil {
		return err
	}

	var firstBooks = `INSERT INTO books (isbn, title, author) VALUES ('978-1789619270', 'Kubernetes Design Patterns and Extensions', 'Onur Yilmaz'),  
																	 ('B07HHDVSJK', 'Cloud-Native Continuous Integration and Delivery', 'Onur Yilmaz')`

	_, err = sbd.db.Exec(firstBooks)
	if err != nil {

		return err
	}

	return nil
}
