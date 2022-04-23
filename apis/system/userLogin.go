package system

import (
	"association/common/response"
	"association/global"
	models "association/modules"
	"association/modules/dto"
	"association/modules/system/request"
	"association/utils"
	"github.com/gin-gonic/gin"
)

// Login 登录认证
func Login(c *gin.Context) {
	loginUser := new(dto.LoginUser)
	if err := c.ShouldBindJSON(loginUser); err != nil {
		response.FailWithMessage("JSON inconsistent type", c)
		return
	}
	if !IsUsername(loginUser.Username) {
		response.FailWithMessage("账号不存在", c)
		return
	}
	loginUser.Password = utils.MD5V([]byte(loginUser.Password))
	u, loginErr := userLoginService.LoginUser(*loginUser)
	if loginErr != nil {
		global.ASS_LOG.Error("login failed username:" + loginUser.Username)
		response.FailWithMessage("账号或密码错误", c)
		return
	}

	claims := utils.CreateClaims(request.BaseClaims{
		ID:        u.Id,
		Username:  u.Username,
		Authority: u.StatusId,
	})
	token, err := utils.GenToken(claims)
	if err != nil {
		global.ASS_LOG.Error("获取token失败,error:" + err.Error())
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(models.LoginRes{
		Username:  u.Username,
		Token:     token,
		ExpiresAt: claims.ExpiresAt,
	}, "登陆成功", c)
	global.ASS_LOG.Info("用户:" + u.Username + ",登录成功")

}
