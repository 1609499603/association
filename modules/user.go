package models

import (
	"association/utils/snowflake"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model

	//账号
	Username string `json:"username"`
	//密码
	Password string `json:"password"`
	//角色（老师或学生，0学生，1老师）
	Role int `json:"role"`
	//身份(1普通，2副社长，3社长，4指导老师，5超级管理)
	StatusId int `json:"status_id"`
	//所属社团id
	AssociationId uint `json:"association_id"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uint(uint64(snowflake.GenID()))
	return
}

// OnlineUser 用户线上数据
type OnlineUser struct {
	Id            uint   `json:"id"`            //用户id
	Username      string `json:"username"`      //用户名
	LoginTime     int64  `json:"loginTime"`     //登录时间
	LoginLocation string `json:"loginLocation"` // 归属地
	Ip            string `json:"ip"`            //ip地址
	Token         string `json:"key"`           // token
}
