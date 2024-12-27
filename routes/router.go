package routes

import (
	"net/http"
	"anjalisreekumar.github/book-manager/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Get all books
			handlers.GetBooks(w, r)
		} else if r.Method == http.MethodPost {
			// Add a new book
			handlers.AddNewBook(w, r)
		} else {
			// Method not allowed
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
