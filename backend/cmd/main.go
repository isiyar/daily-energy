package main

import (
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/http/router"
)
import "github.com/gin-gonic/gin"

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		return
	}

	if !config.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}

	//db, err := ports.InitDatabase(config)
	//if err != nil {
	//	return
	//}
	//
	//println(db.Config)

	r := gin.Default()
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup)
	r.Run()
}
