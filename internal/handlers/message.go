package handlers

import (
	"encoding/json"
	"go-chat-system/internal/database"
	"go-chat-system/internal/models"
	"net/http"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the message into the database using GORM's Create method
	if err := database.DB.Create(&message).Error; err != nil {
		http.Error(w, "Error inserting message", http.StatusInternalServerError)
		return
	}

	// Return the newly created message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
