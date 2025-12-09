package middleware

import (
	"authvault/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}
		split := strings.Split(header, "Bearer ")
		if len(split) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		token := split[1]
		claims, err := utils.ParseToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r.Header.Set("userEmail", claims["email"].(string))
		r.Header.Set("userRole", claims["role"].(string))
		next.ServeHTTP(w, r)
	}
}
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("userRole") != "admin" {
			http.Error(w, "Access denied: Admin only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
