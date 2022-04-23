package models

// Teacher 教师表
type Teacher struct {
	//教师id
	Id int64 `json:"id"`
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
	UserId int64 `json:"user_id"`
	//所属社团id
	AssociationId int64 `json:"association_id"`

	//逻辑删除，记录时间
	Record
}
