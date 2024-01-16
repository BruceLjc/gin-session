package router

import (
	"mysingo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/user/register", service.RegisterHandler)
	r.POST("/user/login", service.LoginHandle)
	r.Run(":8000")
}
