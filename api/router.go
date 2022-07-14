package api

import (
	"github.com/gin-gonic/gin"
	"login-test/middleware"
)

func InitEngine() {
	engine := gin.Default()

	//处理跨域问题
	engine.Use(middleware.Cors)

	engine.POST("/register", Register)
	engine.POST("/login", Login)

	engine.Run(":8090")
}
