package models

import (
	"association/utils/snowflake"
	"gorm.io/gorm"
)

// Student 学生表
type Student struct {
	gorm.Model

	//学院
	CollegeId uint `json:"college_id"`
	//学号
	StudentNumber string `json:"student_number"`
	//姓名
	Name string `json:"name"`
	//性别
	Gender int `json:"gender"`
	//手机号
	Phone string `json:"phone"`
	//邮箱
	Email string `json:"email"`
	//专业
	Major string `json:"major"`
	//班级
	Class string `json:"class"`
	//用户id
	UserId uint `json:"user_id"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uint(uint64(snowflake.GenID()))
	return
}
