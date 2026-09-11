package handlers

import (
	"encoding/json"
	"net/http"
)

// UploadProfileImageHandler is a stub that will be implemented using the Storage interface
func UploadProfileImageHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// This is just a stub for now. Phase 8 will implement full storage.
	// We return a fake URL for now to satisfy the skeleton.
	fakeURL := "/uploads/fake-image.jpg"

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "profile image uploaded (stub)",
		"url":     fakeURL,
		"user_id": userID,
	})
}
