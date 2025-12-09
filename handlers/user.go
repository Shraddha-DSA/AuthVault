package handlers

import (
	"authvault/storage"
	"encoding/json"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("userEmail")
	user := storage.Users[email]
	json.NewEncoder(w).Encode(user)
}
func AdminRoute(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome Admin!",
	})
}
func Logout(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	token := auth[len("Bearer"):]
	storage.Blacklist[token] = true

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful!",
	})
}
