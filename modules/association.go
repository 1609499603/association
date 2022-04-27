package models

import (
	"association/utils/snowflake"
	"gorm.io/gorm"
)

// Association 社团表
type Association struct {
	gorm.Model
	//社团名称
	Name string `json:"name"`
	//简介
	Introduction string `json:"introduction"`
	//招生宣言
	EnrollmentDeclaration string `json:"enrollment_declaration"`
}

func (a *Association) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uint(uint64(snowflake.GenID()))
	return
}

/*

1、目的是 老师想让每个班委先知道，后让同学知道 ，先让班委了解
2、具体要做的 今天晚上8点去910和老师了解一下
3、三月组织的短训，只需要24000
4、可以先体验，暑假再决定留不留下
5、主要方面 ，前端和后端

*/
