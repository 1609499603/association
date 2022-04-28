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
	login := r.Group("/login")
	{
		//登录认证
		login.POST("/", system.Login)
	}
	auth := r.Group("/association")
	auth.Use(middleware.JWTAuthMiddleware()).Use(middleware.CasbinHander())
	{

		//获取社团信息
		auth.GET("/", system.AssociationList)

		//登出
		auth.DELETE("/logout", system.LogoutU)

	}
	r.Run(":8080")
}
