package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	Des            string
	Method         string
	Body           string
	expectedStatus int
	expectedBody   string
	url            string
}

func TestDBBookStore_GetBookByID(t *testing.T) {
	tests := []testCase{
		/*{Des: "Getting books by id",
			Method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":3,"title":"go","author":"Thanush"}`,
			url:            "/Book/3",
		},*/
		{
			Des:            "Book not available in database",
			Method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Book not found\n",
			url:            "/Book/2",
		},
	}
	for _, test := range tests {
		db, _ := sql.Open("mysql", "root:password@tcp/org_db")
		store := NewStore(db)
		req, _ := http.NewRequest(test.Method, test.url, nil)
		w := httptest.NewRecorder()
		store.GetBookByID(w, req)
		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if status := w.Code; status != test.expectedStatus {
			t.Errorf("Test case '%s': expected status %v, got %v", test.Des, test.expectedStatus, status)
		}

		if string(b) != test.expectedBody {
			t.Errorf("Test case '%s': expected body %v, got %v", test.Des, test.expectedBody, string(b))
		}

	}
}

func refreshTable(db *sql.DB) {
	const query = "TRUNCATE Book"

	db.Exec(query)
}

func TestDBBookStore_AddBook(t *testing.T) {
	db, _ := sql.Open("mysql", "root:password@tcp/org_db")

	//refreshTable(db)

	defer refreshTable(db)

	tests := []testCase{
		{Des: "Add new book",
			Method:         http.MethodPost,
			Body:           `{"title":"abbbbb","author":"amul"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":3,"title":"abbbbb","author":"amul"}`,
			url:            "/Book/",
		},
		/*{
			Des:            "Already exist",
			Method:         http.MethodPost,
			Body:           `{"title":"abbbbb","author":"amul"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrorEntityAlreadyExist,
			url:            "/Book/",
		},*/
	}
	for _, test := range tests {

		store := NewStore(db)
		req, _ := http.NewRequest(test.Method, test.url, strings.NewReader(test.Body))
		w := httptest.NewRecorder()
		store.AddBook(w, req)
		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if status := w.Code; status != test.expectedStatus {
			t.Errorf("Test case '%s': expected status %v, got %v", test.Des, test.expectedStatus, status)
		}

		if string(b) != test.expectedBody {
			t.Errorf("Test case '%s': expected body %v, got %v", test.Des, test.expectedBody, string(b))
		}

	}

}
func TestDBBookStore_DeleteBook(t *testing.T) {
	tests := []testCase{
		{
			Des:            "Delete book by id",
			Method:         http.MethodDelete,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Book successfully deleted"}`,
			url:            "/Book/4",
		},
		/*{
			Des:            "Book not found in database",
			Method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Book id not found\n",
			url:            "/Book/10",
		},*/
	}
	for _, test := range tests {
		db, _ := sql.Open("mysql", "root:password@tcp/org_db")
		store := NewStore(db)
		req, _ := http.NewRequest(test.Method, test.url, nil)
		w := httptest.NewRecorder()
		store.DeleteBook(w, req)
		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if status := w.Code; status != test.expectedStatus {
			t.Errorf("Test case '%s': expected status %v, got %v", test.Des, test.expectedStatus, status)
		}

		if string(b) != test.expectedBody {
			t.Errorf("Test case '%s': expected body %v, got %v", test.Des, test.expectedBody, string(b))
		}

	}
}
func TestDBBookStore_UpdateBook(t *testing.T) {
	tests := []testCase{
		/*{
			Des:            "Update the book author name",
			Method:         http.MethodPut,
			Body:           `{"author":"Chandana"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":6,"title":"Python","author":"Chandana"}`,
			url:            "/Book/6",
		},*/
		{
			Des:            "Update the book not present",
			Method:         http.MethodPut,
			Body:           `{"author":"alex"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "The book not found\n",
			url:            "/Book/11",
		},
	}
	for _, test := range tests {
		db, _ := sql.Open("mysql", "root:password@tcp/org_db")
		store := NewStore(db)
		req, _ := http.NewRequest(test.Method, test.url, strings.NewReader(test.Body))
		w := httptest.NewRecorder()
		store.UpdateBook(w, req)
		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if status := w.Code; status != test.expectedStatus {
			t.Errorf("Test case '%s': expected status %v, got %v", test.Des, test.expectedStatus, status)
		}

		if string(b) != test.expectedBody {
			t.Errorf("Test case '%s': expected body %v, got %v", test.Des, test.expectedBody, string(b))
		}

	}
}

func Test_GetsBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	rows := sqlmock.NewRows([]string{"BookId", "Title", "Author"}).
		AddRow(1, "Go Programming", "Dhanush").
		AddRow(2, "Mastering Go", "Smith")

	mock.ExpectQuery("SELECT BookId, Title, Author FROM Book").
		WillReturnRows(rows)

	_, err = GetsBooks(db)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

}
func TestGetBookByID_Success(t *testing.T) {

	mockRows := sqlmock.NewRows([]string{"BookId", "Title", "Author"}).
		AddRow(1, "Go Programming", "Dhanush")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT BookId, Title, Author FROM Book WHERE BookId = ?").
		WithArgs(1).
		WillReturnRows(mockRows)

	book, err := GetBookByID(db, 1)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expectedBook := Book{
		BookId: 1,
		Title:  "Go Programming",
		Author: "Dhanush",
	}

	if !reflect.DeepEqual(book, expectedBook) {
		t.Errorf("Expected book %v, but got %v", expectedBook, book)
	}
}

func TestGetBookByID_fail(t *testing.T) {

	_ = sqlmock.NewRows([]string{"BookId", "Title", "Author"}).
		AddRow(1, "Go Programming", "Dhanush")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT BookId, Title, Author FROM Book WHERE BookId = ?").WithArgs(2).
		WillReturnError(sql.ErrNoRows)
	_, err = GetBookByID(db, 2)
	if err != nil {
		errors.New("error in querying")

	}
	if !strings.Contains(err.Error(), "no book found with ID 2") {
		t.Errorf("expected error message to contain 'no book found with ID 2', got: %v", err)
	}
}
func Test_AddBooks(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	book := &Book{
		Title:  "Go Programming",
		Author: "nkjgw",
	}
	mock.ExpectExec(`INSERT INTO Book \(Title, Author\) VALUES \(\?, \?\)`).
		WithArgs(book.Title, book.Author).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = AddBook(db, book)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if book.BookId != 1 {
		t.Fatalf("expected BookId to be 1, got: %d", book.BookId)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}

}

func TestAddBook_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	book := &Book{
		Title:  "Pythhon Programming",
		Author: "ppppp",
	}

	mock.ExpectExec(`INSERT INTO Book \(Title, Author\) VALUES \(\?, \?\)`).
		WithArgs(book.Title, book.Author).
		WillReturnResult(sqlmock.NewResult(1, 1)) // LastInsertId = 1

	err = AddBook(db, book)

	assert.NoError(t, err)

	assert.Equal(t, 1, book.BookId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}
func TestAddBook_fail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	book := &Book{
		Title:  "Go Programming",
		Author: "omm",
	}

	mock.ExpectExec(`INSERT INTO Book \(Title, Author\) VALUES \(\?, \?\)`).
		WithArgs(book.Title, book.Author).
		WillReturnError(fmt.Errorf("SQL execution error"))

	err = AddBook(db, book)

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "SQL execution error") {
		t.Errorf("expected error message to contain 'SQL execution error', got: %v", err)
	}
}

func TestDeleteBook_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM Book WHERE BookId = \?`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = DeleteBook(db, 1)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestDeleteBook_fail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM Book WHERE BookId = \?`).
		WithArgs(3).
		WillReturnError(fmt.Errorf("SQL execution error"))

	err = DeleteBook(db, 3)

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "SQL execution error") {
		t.Errorf("expected error message to contain 'SQL execution error', got: %v", err)
	}
}

func TestUpdateBooks(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`UPDATE Book SET Author = \? WHERE BookId = \?`).
		WithArgs("Chandana", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = UpdateBooks(db, 1, "Chandana")

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected output is different: %v", err)
	}
}

func TestUpdateBookFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database connection: %v", err)
	}
	defer db.Close()
	mock.ExpectExec(`UPDATE Book SET Author = \? WHERE BookId = \?`).
		WithArgs("hgfh", 10).
		WillReturnError(fmt.Errorf("SQL execution error"))

	err = UpdateBooks(db, 10, "hgfh")

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "SQL execution error") {
		t.Errorf("expected error message to contain 'SQL execution error', got: %v", err)
	}
}
