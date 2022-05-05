package models

import (
	"association/utils/snowflake"
	"gorm.io/gorm"
)

// Notice 公告表
type Notice struct {
	gorm.Model

	//公告标题
	Title string `json:"title"`
	//公告内容
	Content string `json:"content"`
	//是否为平台通知
	IsSystem int `json:"is_system"`
	//所属社团id
	AssociationId uint `json:"association_id"`
	//用户id
	UserId uint `json:"user_id"`
}

func (n *Notice) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uint(uint64(snowflake.GenID()))
	return
}
