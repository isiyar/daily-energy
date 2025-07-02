package main

import (
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/http/router"
)
import "github.com/gin-gonic/gin"

func main() {
	// config, err := LoadConfig()
	// db, err := InitDatabase(config)
	// if err != nil {
	// 	return
	// }

	r := gin.Default()
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup)
	r.Run()
}
