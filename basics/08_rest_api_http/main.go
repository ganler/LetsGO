package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FistName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBook := Book{}
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	newID := 1
	if len(books) != 0 {
		id, err := strconv.Atoi(books[len(books)-1].ID)
		if err != nil {
			panic(err)
		}
		newID = id + 1
	}
	newBook.ID = strconv.Itoa(newID) // Mock ID - not safe...
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

// Create a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	newBook := Book{}
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	for index, item := range books {
		if item.ID == params["id"] {
			newBook.ID = item.ID
			books[index] = newBook
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Create a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	lrz := Author{FistName: "Runzhong", LastName: "Lee"}

	books = append(books, Book{
		ID: "1", Isbn: "112135", Title: "LRZNB", Author: &lrz,
	})

	books = append(books, Book{
		ID: "2", Isbn: "151217", Title: "SuccessBook", Author: &lrz,
	})

	books = append(books, Book{
		ID: "3", Isbn: "666666", Title: "ReinforcementLearning", Author: &Author{
			FistName: "Yi", LastName: "San",
		},
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
