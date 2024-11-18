package handlers

import (
	"encoding/json"
	"go-chat-system/internal/database"
	"go-chat-system/internal/models"
	"net/http"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the chat into the database using GORM's Create method
	if err := database.DB.Create(&chat).Error; err != nil {
		http.Error(w, "Error inserting chat", http.StatusInternalServerError)
		return
	}

	// Return the newly created chat
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}
