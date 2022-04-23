package dto

// RegUser （注册）获取账号,密码,角色
type RegUser struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,max=20,min=6"`
	Role     int    `json:"role"`
}

type RegTeacher struct {
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
	//邮箱验证码
	EmailStr string `json:"email_str"`
	//用户id
	UserId int64 `json:"user_id"`
}

type RegStudent struct {
	//学院
	CollegeId int64 `json:"college_id"`
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
	//邮箱验证码
	EmailStr string `json:"email_str"`
	//专业
	Major string `json:"major"`
	//班级
	Class string `json:"class"`
	//用户id
	UserId int64 `json:"user_id"`
}

type RegEmail struct {
	//邮箱
	Email string `json:"email" validate:"required,email"`
}

type RegUsername struct {
	Username string `json:"username"`
}
