package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/isiyar/daily-energy/backend/config"
)

type ChatHandler struct {
	c config.Config
}

func NewChatHandler(c config.Config) *ChatHandler {
	return &ChatHandler{
		c: c,
	}
}

func (h *ChatHandler) HandleChat(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection established for chat")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		if messageType == websocket.TextMessage {
			userMessage := string(p)
			log.Printf("Received message: %s", userMessage)

			response := "Echo: " + userMessage

			err = conn.WriteMessage(websocket.TextMessage, []byte(response))
			if err != nil {
				log.Printf("Error sending message: %v", err)
				break
			}
			log.Printf("Sent response: %s", response)
		}
	}
	log.Println("WebSocket connection closed")
}