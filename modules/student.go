package models

// Student 学生表
type Student struct {
	//学生id
	Id int64 `json:"id"`
	//学院
	CollegeId int64 `json:"college_id"`
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
	UserId int64 `json:"user_id"`
	//所属社团id
	AssociationId int64 `json:"association_id"`

	//逻辑删除，记录时间
	Record
}
