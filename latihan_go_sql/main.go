package main

import (
	"database/sql"
	"fmt"
	"log"
	"mini_challenge/models"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

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

	createProduct()

	updateProduct("Bubur")

	getProductById(5)

	createVariant()

	fmt.Println("")
	updateVariantById(1)
	deleteVariantById(3)
	getProductWithVariant()
}

func createProduct() {
	var product = models.Product{}

	sqlStatement := `INSERT INTO products (name, created_at, updated_at) VALUES (?,?,?)`

	result, err := db.Exec(sqlStatement, "Susu", time.Now().UTC(), time.Now().UTC())
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	var createdAtBytes, updatedAtBytes []byte
	err = db.QueryRow(sqlRetrieve, lastInsertId).Scan(&product.ID, &product.Name, &createdAtBytes, &updatedAtBytes)
	if err != nil {
		panic(err)
	}

	// Assuming the database returns the timestamps in a standard format, e.g., RFC3339
	const layout = "2006-01-02 15:04:05" // Adjust this layout as necessary

	createdAt, err := time.Parse(layout, string(createdAtBytes))
	if err != nil {
		panic(err)
	}

	updatedAt, err := time.Parse(layout, string(updatedAtBytes))
	if err != nil {
		panic(err)
	}

	product.CreatedAt = createdAt
	product.UpdatedAt = updatedAt

	fmt.Printf("New product data: %v", product)
}

func updateProduct(name string) {
	sqlStatement := `UPDATE products SET name = ?, created_at = ? WHERE id = ?;`

	result, err := db.Exec(sqlStatement, name, time.Now().UTC(), 1)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}

func getProductById(id int) {
	var product = models.Product{}
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	var createdAtBytes, updatedAtBytes []byte
	err = db.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &createdAtBytes, &updatedAtBytes)
	if err != nil {
		panic(err)
	}

	// Assuming the database returns the timestamps in a standard format, e.g., RFC3339
	const layout = "2006-01-02 15:04:05" // Adjust this layout as necessary

	createdAt, err := time.Parse(layout, string(createdAtBytes))
	if err != nil {
		panic(err)
	}

	updatedAt, err := time.Parse(layout, string(updatedAtBytes))
	if err != nil {
		panic(err)
	}

	product.CreatedAt = createdAt
	product.UpdatedAt = updatedAt

	fmt.Printf("Retrieved Data %v", product)
}

func createVariant() {
	var variant = models.Variant{}

	sqlStatement := `INSERT INTO variants (variant_name, quantity, product_id, created_at, updated_at) VALUES (?,?,?,?,?)`

	result, err := db.Exec(sqlStatement, "Susu Sapi", 4, 1, time.Now().UTC(), time.Now().UTC())
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	sqlRetrieve := `SELECT * FROM variants WHERE id = ?`

	var createdAtBytes, updatedAtBytes []byte
	err = db.QueryRow(sqlRetrieve, lastInsertId).Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.ProductID, &createdAtBytes, &updatedAtBytes)
	if err != nil {
		panic(err)
	}

	// Assuming the database returns the timestamps in a standard format, e.g., RFC3339
	const layout = "2006-01-02 15:04:05" // Adjust this layout as necessary

	createdAt, err := time.Parse(layout, string(createdAtBytes))
	if err != nil {
		panic(err)
	}

	updatedAt, err := time.Parse(layout, string(updatedAtBytes))
	if err != nil {
		panic(err)
	}

	variant.CreatedAt = createdAt
	variant.UpdatedAt = updatedAt

	fmt.Printf("New variant data: %v", variant)
}

func updateVariantById(id int) {
	sqlStatement := `UPDATE variants SET variant_name = ?, quantity = ?, updated_at = ? WHERE id = ?;`

	result, err := db.Exec(sqlStatement, "Susu Cabai", 27, time.Now().UTC(), id)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}

func deleteVariantById(id int) {
	sqlStatement := "DELETE FROM variants WHERE id = ?"

	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted row amount:", count)
}

func getProductWithVariant() {
	var variantsWithProduct = []struct {
		name        string
		quantity    int
		productName string
	}{}

	sqlStatement := "SELECT variants.variant_name, variants.quantity, products.name FROM variants JOIN products ON variants.product_id = products.id"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var variant = struct {
			name        string
			quantity    int
			productName string
		}{}

		err = rows.Scan(&variant.name, &variant.quantity, &variant.productName)
		if err != nil {
			panic(err)
		}

		variantsWithProduct = append(variantsWithProduct, variant)

	}

	fmt.Println()
	fmt.Printf("Variants datas: %v\n", variantsWithProduct)

}

// func RetrieveBooks() {
// 	var books = []Book{}

// 	sqlStatement := "SELECT * FROM books"

// 	rows, err := db.Query(sqlStatement)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var book = Book{}

// 		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
// 		if err != nil {
// 			panic(err)
// 		}

// 		books = append(books, book)
// 	}
// 	fmt.Println()
// 	fmt.Printf("Book datas: %+v\n", books)
// }

// func DeleteBook() {
// 	sqlStatement := `DELETE FROM books WHERE id = ?`

// 	result, err := db.Exec(sqlStatement, 3)
// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := result.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Deleted amount: ", count)
// }
