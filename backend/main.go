package main

import "github.com/isiyar/daily-energy/backend/routers"
import "github.com/gin-gonic/gin"

func registerRouters(r *gin.Engine) {
	r.GET("/", routers.HelloWorldRouter)
}

func main() {
	// config, err := LoadConfig()
	// db, err := InitDatabase(config)
	// if err != nil {
	// 	return
	// }
	router := gin.Default()
	registerRouters(router)
	router.Run()
}
