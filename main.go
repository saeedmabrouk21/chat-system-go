package main

import (
	"go-chat-system/internal/database"
	"go-chat-system/internal/handlers"
	"log"
	"net/http"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	http.HandleFunc("/chats", handlers.CreateChat)
	http.HandleFunc("/messages", handlers.CreateMessage)

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
