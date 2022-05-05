package system

import (
	"association/common/response"
	"association/global"
	models "association/modules"
	"association/modules/dto"
	"association/utils"
	"context"
	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {

	u := new(dto.RegUser)
	//ShouldBindJSON的作用就是检查 获取到的数据是否与结构体u的类型一致，若不一致，抛出错误，若一致则赋值
	if err := c.ShouldBindJSON(u); err != nil {
		response.FailWithMessage("JSON inconsistent type", c)
		return
	}
	if IsUsername(u.Username) {
		//账号存在
		response.FailWithMessage("username already exists", c)
		return
	}
	user := models.User{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
	if u.Role == 0 {
		user.StatusId = 1
	} else if u.Role == 1 {
		user.StatusId = 4
	}
	if err := registerService.InsertUser(user); err != nil {
		response.FailWithMessage("insert error", c)
		return
	}
	global.ASS_LOG.Info("User added successfully,Username:" + user.Username)
	_, m := registerService.IsUsername(user.Username)
	response.OkWithDetailed(gin.H{"statusId": user.StatusId, "Id": m.ID}, "注册成功", c)
}

// InsertTeacher 身份为老师
func InsertTeacher(c *gin.Context) {
	t := new(dto.RegTeacher)
	if err := c.ShouldBindJSON(t); err != nil {
		response.FailWithMessage("JSON inconsistent type", c)
		return
	}
	//验证邮箱验证码是否正确
	if emailStr := global.ASS_REDIS.Get(context.Background(), t.Email).Val(); emailStr != t.EmailStr {
		response.FailWithMessage("Email verification code failed", c)
		return
	}
	teacher := models.Teacher{
		TeacherNumber: t.TeacherNumber,
		CollegeId:     t.CollegeId,
		Name:          t.Name,
		Gender:        t.Gender,
		Phone:         t.Phone,
		Email:         t.Email,
		UserId:        t.UserId,
	}

	if err := registerService.InsertTeacher(teacher); err != nil {
		response.FailWithMessage("insert error", c)
		return
	}
	global.ASS_LOG.Info("Teacher added successfully,TeacherName:" + teacher.Name)
	response.Ok(c)
}

// InsertStudent 身份为学生
func InsertStudent(c *gin.Context) {
	s := new(dto.RegStudent)
	if err := c.ShouldBindJSON(s); err != nil {
		response.FailWithMessage("JSON inconsistent type", c)
		return
	}
	//验证邮箱验证码是否正确
	if emailStr := global.ASS_REDIS.Get(context.Background(), s.Email).Val(); emailStr != s.EmailStr {
		response.FailWithMessage("Email verification code failed", c)
		return
	}

	student := models.Student{
		CollegeId:     s.CollegeId,
		StudentNumber: s.StudentNumber,
		Name:          s.Name,
		Gender:        s.Gender,
		Phone:         s.Phone,
		Email:         s.Email,
		Major:         s.Major,
		Class:         s.Class,
		UserId:        s.UserId,
	}

	if err := registerService.InsertStudent(student); err != nil {
		response.FailWithMessage("insert error", c)
		return
	}
	global.ASS_LOG.Info("Teacher added successfully,StudentName:" + student.Name)
	response.Ok(c)
}

// Email 根据邮箱获取验证码
func Email(c *gin.Context) {
	email := new(dto.RegEmail)
	if err := c.ShouldBindJSON(email); err != nil {
		response.FailWithMessage("email type failed", c)
		return
	}
	//先判断redis中是否存在验证码
	if emailStr := global.ASS_REDIS.Get(context.Background(), email.Email).String(); emailStr != "" {
		//存在则删除原有验证码
		global.ASS_REDIS.Del(context.Background(), email.Email)
	}
	//发送邮件，并获取验证码
	code := utils.Send(email.Email)
	//添加到redis设置过期时间为1分钟
	global.ASS_REDIS.Set(context.Background(), email.Email, code, 60*1000*1000*1000)
	m := make(map[string]string, 1)
	m["emailStr"] = code
	global.ASS_LOG.Info("邮箱验证码:" + code)
	response.OkWithMessage("发送成功", c)
}

// IsUsername 检查账号是否存在
func IsUsername(username string) bool {
	_, s := registerService.IsUsername(username)
	if s.Username == username {
		return true
	}
	return false
}
