package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getBooks")

	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "error in marshalling", http.StatusBadRequest)
		return
	}

	w.Write([]byte(json))
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get book by id")
	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr) //converts string to integer
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

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add books")
	body, err := ioutil.ReadAll(r.Body)
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

func updateBook(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	var updateData map[string]string
	err = json.Unmarshal(body, &updateData)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	newAuthor, exists := updateData["author"]
	if exists {

		book.Author = newAuthor
		books[id] = book
	} else {
		http.Error(w, "Author not provided in the request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling updated book", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {

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

	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/books/add", addBook)
	http.HandleFunc("/books/{id}", getBookByID)
	http.HandleFunc("/books/update/{id}", updateBook)
	http.HandleFunc("/books/delete/{id}", deleteBook)

	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
