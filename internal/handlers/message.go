package handlers

import (
	"encoding/json"
	"go-chat-system/internal/database"
	"go-chat-system/internal/models"
	"net/http"
)

// CreateMessage handles the creation of a new message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		ChatNumber int    `json:"chat_number"`
		Body       string `json:"body"` // The body of the message (content)
	}

	// Decode the incoming JSON request body
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the chat exists by its number
	var chat models.Chat
	if err := database.DB.Where("number = ?", requestBody.ChatNumber).First(&chat).Error; err != nil {
		http.Error(w, "Chat not found", http.StatusNotFound)
		return
	}

	// Create a new message object and set the ChatID and Body
	message := models.Message{
		ChatID: chat.ID,
		Body:   requestBody.Body, // Set the body from the request
	}

	// Calculate the next message number for this chat
	var messageCount int64
	database.DB.Model(&models.Message{}).Where("chat_id = ?", message.ChatID).Count(&messageCount)
	message.Number = int(messageCount + 1)

	// Insert the message into the database
	if err := database.DB.Create(&message).Error; err != nil {
		http.Error(w, "Error inserting message", http.StatusInternalServerError)
		return
	}

	// Return only the generated message number
	response := map[string]interface{}{
		"number": message.Number,
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
