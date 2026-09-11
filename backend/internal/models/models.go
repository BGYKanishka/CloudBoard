package models

import "time"

type User struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"`
	ProfileImageURL string    `json:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`   // pending, in_progress, completed
	Priority    string    `json:"priority"` // low, medium, high
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
