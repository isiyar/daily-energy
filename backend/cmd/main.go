package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/http/router"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := ports.InitDatabase(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	r := gin.Default()
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup)
	r.Run()
}
