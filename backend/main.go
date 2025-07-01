package main

import "github.com/isiyar/daily-energy/backend/routers"
import "github.com/gin-gonic/gin"

func main() {
	// config, err := LoadConfig()
	// db, err := InitDatabase(config)
	// if err != nil {
	// 	return
	// }

	router := gin.Default()
	apiGroup := router.Group("/api")
	routers.RingRegisterRoutes(apiGroup)
	routers.HelloWorldRegisterRoutes(apiGroup)
	router.Run()
}
