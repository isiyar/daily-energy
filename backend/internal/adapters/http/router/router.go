package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/handler"
)

func RegisterRoutes(r gin.IRouter, h *handler.Handler, c config.Config) {
	r.GET("/ping", handler.PingHandler)

	users := r.Group("/users", handler.TelegramAuthMiddleware(c.TelegramBotToken))
	{
		users.POST("", h.User.CreateUser)
		users.GET("/:utgid", h.User.GetUser)
		users.PUT("/:utgid", h.User.UpdateUser)
		users.DELETE("/:utgid", h.User.DeleteUser)

		usersUtgid := users.Group("/:utgid")
		{
			usersUtgid.POST("/actions", h.Action.CreateAction)
			usersUtgid.GET("/actions", h.Action.GetActions)
      
      usersUtgid.POST("/plans", h.Plan.CreatePlan)
      usersUtgid.GET("/plans", h.Plan.GetPlans)
      
			usersUtgid.GET("/weight-history", h.UserWeightHistory.GetUserWeightHistory)
			usersUtgid.POST("/weight-history", h.UserWeightHistory.CreateUserWeightHistory)
		}
	}
  
  	actions := r.Group("/actions", handler.TelegramAuthMiddleware(c.TelegramBotToken))
	{
		actions.GET("/:id", h.Action.GetAction)
	}

	ai := r.Group("/ai", handler.TelegramAuthMiddleware(c.TelegramBotToken))
	{
		ai.POST("/calories", h.Ai.CalculationCalories)
	}

	r.GET("/ws/chat", handler.TelegramAuthMiddleware(c.TelegramBotToken), h.Chat.HandleChat)
}