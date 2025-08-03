package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/ai"
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

	conversationHistory := []ai.Message{
		{Role: "system", Content: string(h.c.CaloriesAnalyzer)},
	}

	client := &http.Client{}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		if messageType == websocket.TextMessage {
			userMessage := string(p)
			log.Printf("Received message: %s", userMessage)

			conversationHistory = append(conversationHistory, ai.Message{Role: "user", Content: userMessage})

			chatRequest := ai.ChatRequest{
				Model:    "openrouter/cypher-alpha:free",
				Messages: conversationHistory,
			}

			jsonData, err := json.Marshal(chatRequest)
			if err != nil {
				log.Printf("Error marshaling chat request: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Error processing your request"))
				continue
			}

			req, err := ai.GenerateRequest(h.c, jsonData)
			if err != nil {
				log.Printf("Error generating AI request: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Error processing your request"))
				continue
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error sending AI request: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Error processing your request"))
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("AI request failed with status: %d", resp.StatusCode)
				conn.WriteMessage(websocket.TextMessage, []byte("Error from AI service"))
				continue
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading AI response: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Error processing your request"))
				continue
			}
			fmt.Println(bodyBytes)

			var apiResponse ai.APIResponse
			err = ai.Deserialization(bodyBytes, &apiResponse)
			if err != nil {
				log.Printf("Error parsing AI response: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Error processing your request"))
				continue
			}

			aiMessage := apiResponse.Choices[0].Message.Content
			conversationHistory = append(conversationHistory, ai.Message{Role: "assistant", Content: aiMessage})

			err = conn.WriteMessage(websocket.TextMessage, []byte(aiMessage))
			if err != nil {
				log.Printf("Error sending message: %v", err)
				break
			}
			log.Printf("Sent response: %s", aiMessage)
		}
	}
	log.Println("WebSocket connection closed")
}
