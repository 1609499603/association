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
		//获取社团共有多少数据，多少页
		auth.GET("/", system.AssociationFirst)
		//获取当前页面用户信息
		auth.GET("/all", system.AssociationList)

		//根据社团名称检索后加入缓存 返回有多少数据，多少页
		auth.GET("/name", system.AssociationNameFirst)
		//返回根据社团名称检索的数据
		auth.GET("/name/all", system.AssociationName)

		//登出
		auth.DELETE("/logout", system.LogoutU)

	}
	r.Run(":8080")
}
