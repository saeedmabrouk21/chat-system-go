package handlers

import (
	"encoding/json"
	"go-chat-system/internal/database"
	"go-chat-system/internal/models"
	"net/http"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {
	// Automatically generate the number for the chat by finding the max number in the DB and incrementing it
	var lastChat models.Chat
	if err := database.DB.Order("number desc").First(&lastChat).Error; err != nil {
		// If no chats exist, start from number 1
		lastChat.Number = 0
	}

	// Create a new chat with the incremented number
	chat := models.Chat{
		Number: lastChat.Number + 1,
	}

	// Insert the chat into the database using GORM's Create method
	if err := database.DB.Create(&chat).Error; err != nil {
		http.Error(w, "Error inserting chat", http.StatusInternalServerError)
		return
	}

	// Return the generated number in the response
	response := map[string]interface{}{
		"number": chat.Number,
	}

	// Set the response header and return the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
