package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	// Простые CRUD‑маршруты для User
	//r.POST("/login", login.Login)
	//r.POST("/register", register.Register)
	return r
}
