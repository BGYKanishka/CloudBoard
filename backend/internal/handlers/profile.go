package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"cloudboard/internal/database"
	"cloudboard/internal/storage"
)

var fileStorage storage.Storage

func init() {
	backend := os.Getenv("STORAGE_BACKEND")
	if backend == "s3" {
		fileStorage = storage.NewS3Storage(os.Getenv("S3_BUCKET"), os.Getenv("AWS_REGION"))
	} else {
		fileStorage = storage.NewLocalStorage(os.Getenv("LOCAL_STORAGE_PATH"))
	}
}

// UploadProfileImageHandler uses the Storage interface
func UploadProfileImageHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// This is a simulated upload. In a real app we parse multipart form here.
	// We'll use a dummy file reader for the sake of the exercise.
	dummyFile, _ := os.Open("../../README.md") // just something to read
	defer dummyFile.Close()

	url, err := fileStorage.Upload("profile_pic.jpg", dummyFile)
	if err != nil {
		http.Error(w, `{"error":"upload failed"}`, http.StatusInternalServerError)
		return
	}

	// Update DB
	_, err = database.DB.Exec("UPDATE users SET profile_image_url = $1 WHERE id = $2", url, userID)
	if err != nil {
		http.Error(w, `{"error":"db update failed"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "profile image uploaded",
		"url":     url,
		"user_id": userID,
	})
}
