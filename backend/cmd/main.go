package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/http/router"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		return
	}

	if !config.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := ports.InitDatabase(config)
	if err != nil {
		return
	}
	
	fmt.Println(db.Config)

	r := gin.Default()
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup)
	r.Run()
}
