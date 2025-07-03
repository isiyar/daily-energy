package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/http/router"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/repository"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := repository.InitDatabase(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	r := gin.Default()
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup)
	r.Run()
}
