package handlers

import (
	"encoding/json"
	"go-chat-system/internal/database"
	"go-chat-system/internal/models"
	"log"
	"net/http"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {
	var chat models.Chat
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use a SQL query to insert the chat
	query := "INSERT INTO chats (application_token, number) VALUES (?, ?)"
	result, err := database.DB.Exec(query, chat.ApplicationToken, chat.Number)
	if err != nil {
		log.Println("Error inserting chat:", err)
		http.Error(w, "Error inserting chat", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	chat.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}
