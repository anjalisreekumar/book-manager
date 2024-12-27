package handlers

import (
	"anjalisreekumar.github/book-manager/models"
	"anjalisreekumar.github/book-manager/db"
	"encoding/json"
	"net/http"
	"fmt"
)

//fetch all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Books)

}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	// Auto-generate ID for new book
	newBook.ID = fmt.Sprintf("%d", len(db.Books)+1)

	// Append the new book to the book list
	db.Books = append(db.Books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}
