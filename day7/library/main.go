package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Book struct to represent a book
type Book struct {
	BookId int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = make(map[int]Book)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(books)
	if err != nil {
		fmt.Println("error in marshalling")
		http.Error(w, "error in marshalling", http.StatusBadRequest)
		return
	}

	w.Write(json)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)

	if err != nil {

		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}
	var book *Book
	for _, b := range books {
		if b.BookId == id {
			book = &b
			break
		}
	}
	if book == nil {
		http.Error(w, "Book not found", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func AddBook(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the request from body", http.StatusBadRequest)
		return
	}
	var newBook Book
	err = json.Unmarshal(body, &newBook)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}
	newBook.BookId = len(books) + 1

	books[newBook.BookId] = newBook

	w.Header().Set("Content-Type", "application/json")

	booksJSON, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error marshalling books", http.StatusInternalServerError)
		return
	}

	w.Write(booksJSON)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, exists := books[id]
	if !exists {
		fmt.Println("Error: book not found")
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	var updateBookData Book
	err = json.Unmarshal(body, &updateBookData)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	if updateBookData.Author != "" {
		book.Author = updateBookData.Author
	}

	books[id] = book

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling updated book", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	_, exists := books[id]
	if !exists {

		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Book successfully deleted",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func main() {
	books[1] = Book{BookId: 1, Title: "Go", Author: "ABC"}
	books[2] = Book{BookId: 2, Title: "Java", Author: "MNO"}

	http.HandleFunc("/book/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			path := r.URL.Path
			parts := strings.Split(path, "/")
			if len(parts) > 2 && parts[len(parts)-1] != "" {
				GetBookByID(w, r)
			} else {
				GetBooks(w, r)
			}
		case http.MethodPost:
			AddBook(w, r)
		case http.MethodPut:
			UpdateBook(w, r)
		case http.MethodDelete:
			DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
