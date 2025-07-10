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

	planUC := usecase.NewPlanUseCase(repository.NewPlanRepository(db))
	planHandler := handler.NewPlanHandler(c, planUC, userUC)

	aiHandler := handler.NewAiHandler(c)
	chatHandler := handler.NewChatHandler(c)

	h := handler.NewHandler(actionHandler, userHandler, userWeightHistoryHandler, planHandler, aiHandler, chatHandler)

	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://frontend-dev:5173", "https://test-srvr.ru"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "initData"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	apiGroup := r.Group("/api")
	router.RegisterRoutes(apiGroup, h, c)
	r.Run()
}