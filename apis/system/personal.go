package system

import (
	"association/common/response"
	"association/modules/dto"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Personal 个人信息主页
func Personal(c *gin.Context) {
	getId, _ := c.Get("id")
	id := cast.ToString(getId)
	user, err := personalService.SelectPersonalUser(id)
	if err != nil {
		response.FailWithMessage("查询信息失败，请重试", c)
		return
	}
	user.Password = "******"
	if user.StatusId == 5 {
		response.OkWithData(gin.H{
			"user": user,
		}, c)
		return
	}
	if user.Role == 0 {
		student, err := personalService.SelectPersonalStudent(id)
		if err != nil {
			response.FailWithMessage("查询信息失败", c)
			return
		}
		response.OkWithData(gin.H{
			"user":    user,
			"student": student,
		}, c)
		return
	} else {
		teacher, err := personalService.SelectPersonalTeacher(id)
		if err != nil {
			response.FailWithMessage("查询信息失败", c)
			return
		}
		response.OkWithData(gin.H{
			"user":    user,
			"teacher": teacher,
		}, c)
		return
	}

}

// RevisePersonalTeacher 修改教师信息
func RevisePersonalTeacher(c *gin.Context) {
	teacher := new(dto.UpdateTeacher)
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		response.FailWithMessage("json 格式错误", c)
		return
	}
	getId, _ := c.Get("id")
	id := cast.ToString(getId)
	updateErr := personalService.UpdateTeacher(*teacher, id)
	if updateErr != nil {
		response.FailWithMessage("修改信息失败", c)
		return
	}
	response.OkWithMessage("修改信息成功", c)

}

// RevisePersonalStudent 修改学生信息
func RevisePersonalStudent(c *gin.Context) {
	student := new(dto.UpdateStudent)
	err := c.ShouldBindJSON(&student)
	if err != nil {
		response.FailWithMessage("json 格式错误", c)
		return
	}
	getId, _ := c.Get("id")
	id := cast.ToString(getId)
	updateErr := personalService.UpdateStudent(*student, id)
	if updateErr != nil {
		response.FailWithMessage("修改信息失败", c)
		return
	}
	response.OkWithMessage("修改信息成功", c)

}
