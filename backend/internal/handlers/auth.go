package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"cloudboard/internal/auth"
	"cloudboard/internal/database"
	"cloudboard/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error":"server error"}`, http.StatusInternalServerError)
		return
	}

	var newID int
	err = database.DB.QueryRow(
		"INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
		req.Name, req.Email, string(hash),
	).Scan(&newID)

	if err != nil {
		http.Error(w, `{"error":"could not create user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": newID})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	err := database.DB.QueryRow(
		"SELECT id, password_hash FROM users WHERE email = $1", req.Email,
	).Scan(&user.ID, &user.PasswordHash)

	if err != nil {
		http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, `{"error":"could not generate token"}`, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false, // set to true in production
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	json.NewEncoder(w).Encode(map[string]string{"message": "logged in"})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	json.NewEncoder(w).Encode(map[string]string{"message": "logged out"})
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var user models.User
	err = database.DB.QueryRow(
		"SELECT id, name, email, profile_image_url FROM users WHERE id = $1", claims.UserID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.ProfileImageURL)

	if err != nil {
		http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
