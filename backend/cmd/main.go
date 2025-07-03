package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/db"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/http/router"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/repository"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/handler"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := db.InitDatabase(c)
	if err != nil {
		panic(err)
	}

	UserUC := usecase.NewUserUseCase(repository.NewUserRepository(db))
	userHandler := handler.NewUserHandler(UserUC)

	h := handler.NewHandler(userHandler)

	r := gin.Default()
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup, h)
	r.Run()
}
