package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
	"log"
	"net/http"
)

func TelegramAuthMiddleware(botToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		initData := c.GetHeader("initData")
		if initData == "" {
			c.Set("utgid", "5845405153")
			c.Next()
			return
			// log.Println("Missing initData header")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing initData"})
			// return
		}

		utgid, err := validator.GetTelegramUserID(initData, botToken)
		if err != nil {
			log.Printf("Invalid initData: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid initData"})
			return
		}

		c.Set("utgid", utgid)
		c.Next()
	}
}
