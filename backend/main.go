package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/routes"
)

func main() {
	router := gin.Default()
	apiGroup := router.Group("/api")
	routes.RegisterRoutes(apiGroup)
	router.Run()
}
