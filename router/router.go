package router

import (
	"association/apis/system"
	"association/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	//注册相关
	reg := r.Group("/register")
	{
		reg.POST("/", system.Register)
		reg.POST("/teacher", system.InsertTeacher)
		reg.POST("/student", system.InsertStudent)
		reg.POST("/email", system.Email)
	}
	auth := r.Group("/association")
	{
		//登录认证
		auth.POST("/login", system.Login)

	}
	auth.Use(middleware.JWTAuthMiddleware()).Use(middleware.CasbinHander())
	{
		//123
		auth.GET("/home")
	}
	r.Run(":8080")
}
