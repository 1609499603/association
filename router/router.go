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

		//加入社团
		auth.PUT("/join", system.Join)

		//审批加入社团页面
		auth.POST("/examine", system.Examine)

		//同意加入社团
		auth.POST("/agree_association", system.AgreeAssociation)

		//退出社团
		auth.PUT("/exit", system.Exit)

		//我的社团信息
		auth.GET("/my", system.MyAssociation)

		//我的社团用户信息
		auth.GET("/my_association_user", system.MyAssociationUser)

		//发送通知
		auth.POST("/send_notice", system.SendNotice)

		//我的通知列表（社团）
		auth.GET("/notice", system.Notice)

	}
	personal := r.Group("/personal")
	personal.Use(middleware.JWTAuthMiddleware()).Use(middleware.CasbinHander())
	{
		//个人信息主页
		personal.GET("/", system.Personal)

		//修改教师信息
		personal.PUT("/revise_teacher", system.RevisePersonalTeacher)

		//修改学生信息
		personal.PUT("/revise_student", system.RevisePersonalStudent)

	}
	r.Run(":8080")
}
