package handlers

import (
	"anjalisreekumar.github/book-manager/db"
	"encoding/json"
	"net/http"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Books)

}
