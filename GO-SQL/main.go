package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	password = "Sijar123@"
	dbname   = "db_course_go"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID     int
	Title  string
	Author string
	Stock  int
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err = sql.Open("mysql", mysqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connect to Database")

	CreateBook()

	RetrieveBooks()
	UpdateBook()
	DeleteBook()
}

func CreateBook() {
	var book = Book{}

	sqlStatement := `INSERT INTO books (title, author, stock) VALUES (?,?,?)`

	result, err := db.Exec(sqlStatement, "Malisia", "Nonok Karuniawan", 30)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM books WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertId).Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New book data: %v", book)
}

func RetrieveBooks() {
	var books = []Book{}

	sqlStatement := "SELECT * FROM books"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
	}
	fmt.Println()
	fmt.Printf("Book datas: %+v\n", books)
}

func UpdateBook() {
	sqlStatement := `UPDATE books SET title = ?, author = ?, stock = ? WHERE id = ?;`

	result, err := db.Exec(sqlStatement, "Biagy", "Nana Karuniawan", 30, 1)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}

func DeleteBook() {
	sqlStatement := `DELETE FROM books WHERE id = ?`

	result, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted amount: ", count)
}
