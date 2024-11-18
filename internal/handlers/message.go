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

	// Use a SQL query to insert the message
	query := "INSERT INTO messages (chat_id, body, number) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, message.ChatID, message.Body, message.Number)
	if err != nil {
		http.Error(w, "Error inserting message", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	message.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
