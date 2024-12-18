package main

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
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
			expectedBody:   `{"id":17,"title":"abbbbb","author":"amul"}`,
			url:            "/Book/",
		},
		/*{
			Des:            "Already exist",
			Method:         http.MethodPost,
			Body:           `{"title":"go","author":"Thanush"}`,
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
		{
			Des:            "Update the book author name",
			Method:         http.MethodPut,
			Body:           `{"author":"Chandana"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":6,"title":"Python","author":"Chandana"}`,
			url:            "/Book/6",
		},
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
