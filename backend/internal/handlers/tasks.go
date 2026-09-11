package handlers

import (
	"encoding/json"
	"net/http"

	"cloudboard/internal/auth"
	"cloudboard/internal/database"
	"cloudboard/internal/models"
	"github.com/go-chi/chi/v5"
)

func getUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return 0, err
	}
	claims, err := auth.ValidateToken(cookie.Value)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query("SELECT id, user_id, title, description, status, priority, created_at, updated_at FROM tasks WHERE user_id = $1 ORDER BY created_at DESC", userID)
	if err != nil {
		http.Error(w, `{"error":"could not fetch tasks"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.CreatedAt, &t.UpdatedAt); err != nil {
			continue
		}
		tasks = append(tasks, t)
	}

	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var req models.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	if req.Status == "" {
		req.Status = "pending"
	}
	if req.Priority == "" {
		req.Priority = "medium"
	}

	var newID int
	err = database.DB.QueryRow(
		"INSERT INTO tasks (user_id, title, description, status, priority) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		userID, req.Title, req.Description, req.Status, req.Priority,
	).Scan(&newID)

	if err != nil {
		http.Error(w, `{"error":"could not create task"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": newID})
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	var t models.Task
	err = database.DB.QueryRow("SELECT id, user_id, title, description, status, priority, created_at, updated_at FROM tasks WHERE id = $1 AND user_id = $2", id, userID).
		Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		http.Error(w, `{"error":"task not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	var req models.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE tasks SET title = $1, description = $2, status = $3, priority = $4, updated_at = NOW() WHERE id = $5 AND user_id = $6",
		req.Title, req.Description, req.Status, req.Priority, id, userID)

	if err != nil {
		http.Error(w, `{"error":"could not update task"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "task updated"})
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	_, err = database.DB.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2", id, userID)
	if err != nil {
		http.Error(w, `{"error":"could not delete task"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "task deleted"})
}

func UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")
	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request"}`, http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE tasks SET status = $1, updated_at = NOW() WHERE id = $2 AND user_id = $3", req.Status, id, userID)
	if err != nil {
		http.Error(w, `{"error":"could not update task status"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "task status updated"})
}
