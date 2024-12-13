package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase struct {
	des            string
	books          map[int]Book
	body           string
	method         string
	expectedStatus int
	expectedBody   string
	url            string
	updateAuthor   string
}

func TestGetBooks(t *testing.T) {
	tests := []testCase{
		{
			des: "Getting books",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"1":{"id":1,"title":"Go","author":"ABC"},"2":{"id":2,"title":"Java","author":"MNO"}}`,
		},
		{
			des: "Getting books",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
				3: {BookId: 3, Title: "python", Author: "ppp"},
			},
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"1":{"id":1,"title":"Go","author":"ABC"},"2":{"id":2,"title":"Java","author":"MNO"},"3":{"id":3,"title":"python","author":"ppp"}}`,
		},
	}

	for _, tt := range tests {
		books = tt.books

		req, _ := http.NewRequest(tt.method, "/book/", nil)
		w := httptest.NewRecorder()
		GetBooks(w, req)

		if status := w.Code; status != tt.expectedStatus {
			t.Errorf("expected status %v, got %v", tt.expectedStatus, status)
		}

		if w.Body.String() != tt.expectedBody {
			t.Errorf("expected body %v, got %v", tt.expectedBody, w.Body.String())
		}
	}
}

func TestGetBooksByID(t *testing.T) {
	tests := []testCase{
		{
			des: "get book by id",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"title":"Go","author":"ABC"}`,
			url:            "/book/1",
		},

		{
			des: "No Books",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			method:         http.MethodGet,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Book not found\n",
			url:            "/book/3",
		},
	}
	for _, tt := range tests {
		books = tt.books
		//body := bytes.NewBuffer(tt.body)

		req, _ := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
		w := httptest.NewRecorder()
		GetBookByID(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if status := w.Code; status != tt.expectedStatus {
			t.Errorf("Test case '%s': expected status %v, got %v", tt.des, tt.expectedStatus, status)
		}

		if string(b) != tt.expectedBody {
			t.Errorf("Test case '%s': expected body %v, got %v", tt.des, tt.expectedBody, string(b))
		}
	}
}

func TestAddBook(t *testing.T) {
	tests := []testCase{
		{
			des: "Adding of books",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},

			body:           `{"title": "Python", "author": "Nayana"}`,
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"1":{"id":1,"title":"Go","author":"ABC"},"2":{"id":2,"title":"Java","author":"MNO"},"3":{"id":3,"title":"Python","author":"Nayana"}}`,
		},
		{
			des: "Invalid JSON format",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			body:           `{"title": "C++", "author":"nnnn"`,
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Error unmarshalling request body\n",
		},
	}

	for _, tt := range tests {
		books = tt.books

		req, _ := http.NewRequest(tt.method, "/book/", strings.NewReader(tt.body))
		w := httptest.NewRecorder()
		AddBook(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if w.Code != tt.expectedStatus {
			t.Errorf("expected status %v, got %v", tt.expectedStatus, w.Code)
		}

		if string(b) != tt.expectedBody {
			t.Errorf("Expected %s, but got %s", tt.expectedBody, string(b))
			//fmt.Println(strings.Contains(tt.expectedBody, string(b)))
		}

	}
}
func TestUpdateBooks(t *testing.T) {
	tests := []testCase{
		{
			des: "Update books Successfully",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			method:         http.MethodPut,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"title":"Go","author":"New author"}`,
			url:            "/book/1",
			updateAuthor:   `{"author": "New author"}`,
		},
		{
			des: "Book not found",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
			},
			method:         http.MethodPut,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Book not found",

			updateAuthor: `{"author": "New author"}`,
			url:          "/book/5",
		},
	}
	for _, tt := range tests {

		books = tt.books

		req := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.updateAuthor))
		w := httptest.NewRecorder()

		UpdateBook(w, req)

		if w.Result().StatusCode != tt.expectedStatus {
			t.Errorf("%s: expected status %d, got %d", tt.des, tt.expectedStatus, w.Result().StatusCode)
		}

		body := strings.TrimSpace(w.Body.String())
		if body != tt.expectedBody {
			t.Errorf("%s: expected body %s, got %s", tt.des, tt.expectedBody, body)
		}
	}
}

func TestDeleteBooks(t *testing.T) {
	tests := []testCase{
		{
			des: "Deleted books Successfully",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			method:         http.MethodDelete,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Book successfully deleted"}`,
			url:            "/book/2",
		},
		{
			des: "Book not exist",
			books: map[int]Book{
				1: {BookId: 1, Title: "Go", Author: "ABC"},
				2: {BookId: 2, Title: "Java", Author: "MNO"},
			},
			method:         http.MethodDelete,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Book not found\n",
			url:            "/book/3",
		},
	}

	for _, tt := range tests {
		books = tt.books

		req, _ := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
		w := httptest.NewRecorder()
		DeleteBook(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if w.Code != tt.expectedStatus {
			t.Errorf("Expected status %v, got %v", tt.expectedStatus, w.Code)
		}

		if string(b) != tt.expectedBody {
			t.Errorf("Expected %s, but got %s", tt.expectedBody, string(b))
		}

		if tt.method == http.MethodDelete {
			if tt.url == "/book/2" && len(books) != 1 {
				t.Errorf("Test case '%s': expected 1 book, got %v", tt.des, len(books))
			}
		}

	}
}
