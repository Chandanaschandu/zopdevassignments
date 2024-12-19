package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	ErrorEntityNotFound     = "Entity not found"
	ErrorEntityAlreadyExist = "Entity already exists"
)

type DBBookStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *DBBookStore {
	return &DBBookStore{db: db}
}

// Book struct to represent a book
type Book struct {
	BookId int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookStore interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBookByID(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

func (store *DBBookStore) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := store.db.Query("SELECT BookId, Title, Author FROM Book")
	if err != nil {
		log.Printf("error :%v", err)
		return
	}

	var books = make([]Book, 0)
	for rows.Next() {
		var book Book
		_ = rows.Scan(&book.BookId, &book.Title, &book.Author)
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error marshalling books", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func (store *DBBookStore) GetBookByID(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request, book ID missing", http.StatusBadRequest)
		return
	}

	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	var book Book
	query := "SELECT BookId, Title, Author FROM Book WHERE BookId= ?"

	log.Printf("Executing query: %s", query)

	err = store.db.QueryRow(query, id).Scan(&book.BookId, &book.Title, &book.Author)

	if err == sql.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	} else if err != nil {

		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error marshalling book: %v", err)
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (store *DBBookStore) AddBook(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var newBook Book
	_ = json.Unmarshal(body, &newBook)

	var exists bool
	_ = store.db.QueryRow("SELECT EXISTS(SELECT 1 FROM Book WHERE Title = ? AND Author = ?)", newBook.Title, newBook.Author).Scan(&exists)

	if exists {
		http.Error(w, ErrorEntityAlreadyExist, http.StatusBadRequest)
		return
	}

	// Insert the new book
	result, _ := store.db.Exec("INSERT INTO Book (Title, Author) VALUES (?, ?)", newBook.Title, newBook.Author)
	id, _ := result.LastInsertId()
	newBook.BookId = int(id)

	jsonData, _ := json.Marshal(newBook)
	w.Write(jsonData)
}

func (store *DBBookStore) UpdateBook(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book Book
	ans := store.db.QueryRow("SELECT BookId, Title, Author FROM Book WHERE BookId = ?", id).Scan(&book.BookId, &book.Title, &book.Author)
	if ans != nil {
		http.Error(w, "The book not found", http.StatusBadRequest)
		return
	}
	body, _ := io.ReadAll(r.Body)

	var updateBookData Book
	_ = json.Unmarshal(body, &updateBookData)

	if updateBookData.Author != "" {
		book.Author = updateBookData.Author
	}
	if updateBookData.Title != "" {
		book.Title = updateBookData.Title
	}

	_, err = store.db.Exec("UPDATE Book SET Title = ?, Author = ? WHERE BookId= ?", book.Title, book.Author, book.BookId)
	if err != nil {
		http.Error(w, "Book not exists", http.StatusInternalServerError)
	}

	jsonData, _ := json.Marshal(book)
	w.Write(jsonData)
}

func (store *DBBookStore) DeleteBook(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID format", http.StatusBadRequest)
		return
	}

	_, err = store.db.Exec("DELETE FROM Book WHERE BookId= ?", id)
	/*if exists != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}*/
	if err != nil {
		http.Error(w, "Book id not found", http.StatusNotFound)
	}

	response := map[string]string{
		"message": "Book successfully deleted",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
	//w.Write([]byte("Book deleted"))
}

func main() {
	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@tcp/org_db")
	if err != nil {
		fmt.Println("Error ", err)
	}

	err = db.Ping()
	if err != nil {

		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the MySQL database!")
	store := NewStore(db)

	http.HandleFunc("/Book/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			path := r.URL.Path
			parts := strings.Split(path, "/")
			if len(parts) > 1 && parts[len(parts)-1] != "" {
				store.GetBookByID(w, r)
			} else {
				store.GetBooks(w, r)
			}
		case http.MethodPost:
			store.AddBook(w, r)
		case http.MethodPut:
			store.UpdateBook(w, r)
		case http.MethodDelete:
			store.DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func GetsBooks(db *sql.DB) ([]Book, error) {

	rows, err := db.Query("SELECT BookId, Title, Author FROM Book")
	if err != nil {
		log.Printf("error :%v", err)
		return nil, err
	}

	var books = make([]Book, 0)
	for rows.Next() {
		var book Book
		_ = rows.Scan(&book.BookId, &book.Title, &book.Author)
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating over rows: %v", err)
		return nil, err
	}
	return books, nil
}

func GetBookByID(db *sql.DB, BookId int64) (Book, error) {
	var book Book
	query := "SELECT BookId, Title, Author FROM Book WHERE BookId = ?"

	err := db.QueryRow(query, BookId).Scan(&book.BookId, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return Book{}, fmt.Errorf("no book found with ID %d", BookId)
		}
		return Book{}, fmt.Errorf("error querying book: %v", err)
	}
	return book, nil
}

func AddBook(db *sql.DB, book *Book) error {

	query := "INSERT INTO Book (Title, Author) VALUES (?, ?)"

	result, err := db.Exec(query, book.Title, book.Author)
	if err != nil {
		return fmt.Errorf("error inserting book: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error retrieving last insert ID: %v", err)
	}

	book.BookId = int(lastInsertID)
	return nil
}
func DeleteBook(db *sql.DB, BookId int) error {
	query := "DELETE FROM Book WHERE BookId = ?"
	result, err := db.Exec(query, BookId)
	if err != nil {
		return fmt.Errorf("error deleting book: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error retrieving rows affected: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no rows affected, book with ID %d not found", BookId)
	}

	return nil
}
func UpdateBooks(db *sql.DB, BookId int, Author string) error {
	query := "UPDATE Book SET Author = ? WHERE BookId = ?"
	res, err := db.Exec(query, Author, BookId)
	if err != nil {
		return fmt.Errorf("error updating book: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error in  retrieving : %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, book with ID %d not found", BookId)
	}

	return nil
}
