package handlers

import (
	"authvault/models"
	"authvault/storage"
	"authvault/utils"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)

	if _, exists := storage.Users[data.Email]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	hashed, _ := utils.HashPassword(data.Password)
	data.Password = hashed
	storage.Users[data.Email] = data
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered Successfully",
	})
}
func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)
	stored, exists := storage.Users[creds.Email]
	if !exists || !utils.CheckPasswordHash(stored.Password, creds.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	access, _ := utils.CreateToken(stored.Email, stored.Role)
	refresh, _ := utils.CreateRefreshToken(stored.Email)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Refresh string `json:"refresh_token"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	claims, err := utils.ParseToken(body.Refresh)
	if err != nil {
		http.Error(w, "Invalid refresh Token", http.StatusUnauthorized)
		return
	}
	email := claims["email"].(string)
	role := storage.Users[email].Role
	newAccess, _ := utils.CreateToken(email, role)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": newAccess,
	})
}
