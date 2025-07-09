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
		users.POST("", h.User.CreateUser)           // POST /users
		users.GET("/:utgid", h.User.GetUser)        // GET /users/:utgid
		users.PUT("/:utgid", h.User.UpdateUser)     // PUT /users/:utgid
		users.DELETE("/:utgid", h.User.DeleteUser)  // DELETE /users/:utgid

		// Подгруппа /users/:utgid (действия и история веса)
		usersUtgid := users.Group("/:utgid")
		{
			usersUtgid.POST("/actions", h.Action.CreateAction)         // POST /users/:utgid/actions
			usersUtgid.GET("/actions", h.Action.GetActions)            // GET /users/:utgid/actions
			usersUtgid.GET("/weight-history", h.UserWeightHistory.GetUserWeightHistory)    // GET /users/:utgid/weight-history
			usersUtgid.POST("/weight-history", h.UserWeightHistory.CreateUserWeightHistory) // POST /users/:utgid/weight-history
		}
	}

	// Группа /actions (требует аутентификации)
	actions := r.Group("/actions", handler.TelegramAuthMiddleware(c.TelegramBotToken))
	{
		actions.GET("/:id", h.Action.GetAction) // GET /actions/:id
	}

	// Группа /ai (требует аутентификации)
	ai := r.Group("/ai", handler.TelegramAuthMiddleware(c.TelegramBotToken))
	{
		ai.POST("/calories", h.Ai.CalculationCalories) // POST /ai/calories
	}

	// Чат (отдельный маршрут, уже с middleware в main.go)
	r.GET("/ws/chat", handler.TelegramAuthMiddleware(c.TelegramBotToken), h.Chat.HandleChat) // GET /ws/chat
}