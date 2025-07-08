package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/adapters/db"
	"github.com/isiyar/daily-energy/backend/internal/adapters/http/router"
	"github.com/isiyar/daily-energy/backend/internal/adapters/repository"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
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

	userUC := usecase.NewUserUseCase(repository.NewUserRepository(db))
	userHandler := handler.NewUserHandler(userUC)

	UserWeightHistoryUC := usecase.NewUserWeightHistoryUseCase(repository.NewUserWeightHistoryRepository(db))
	userWeightHistoryHandler := handler.NewUserWeightHistoryHandler(userUC, UserWeightHistoryUC)

	actionUC := usecase.NewActionUseCase(repository.NewActionRepository(db))
	actionHandler := handler.NewActionHandler(actionUC, userUC)

	aiHandler := handler.NewAiHandler(c)

	h := handler.NewHandler(actionHandler, userHandler, userWeightHistoryHandler, aiHandler)

	r := gin.Default()
	r.Use(cors.Default())
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup, h)
	r.Run()
}
