package models

import (
	"association/utils/snowflake"
	"gorm.io/gorm"
)

// Teacher 教师表
type Teacher struct {
	gorm.Model

	//职工号
	TeacherNumber string `json:"teacher_number"`
	//学院
	CollegeId int64 `json:"college_id"`
	//姓名
	Name string `json:"name"`
	//性别
	Gender int `json:"gender"`
	//手机号
	Phone string `json:"phone"`
	//邮箱
	Email string `json:"email"`
	//用户id
	UserId uint `json:"user_id"`
	//所属社团id
	AssociationId uint `json:"association_id"`
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uint(uint64(snowflake.GenID()))
	return
}
