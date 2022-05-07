package dto

type UpdateStudent struct {
	//学院
	CollegeId uint `json:"college_id"`
	//学号
	StudentNumber string `json:"student_number"`
	//姓名
	Name string `json:"name"`
	//性别
	Gender int `json:"gender"`
	//手机号
	Phone string `json:"phone" validate:"required,len=11"`
	//邮箱
	Email string `json:"email" validate:"required,email"`
	//专业
	Major string `json:"major"`
	//班级
	Class string `json:"class"`
}

type UpdateTeacher struct {
	//职工号
	TeacherNumber string `json:"teacher_number"`
	//学院
	CollegeId int64 `json:"college_id"`
	//姓名
	Name string `json:"name"`
	//性别
	Gender int `json:"gender"`
	//手机号
	Phone string `json:"phone" validate:"required,len=11"`
	//邮箱
	Email string `json:"email" validate:"required,email"`
}
