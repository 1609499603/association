package system

import (
	"association/common/response"
	"association/global"
	models "association/modules"
	"association/modules/dto"
	"association/modules/system/request"
	"association/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
	online := new(models.OnlineUser)
	global.ASS_REDIS.Get(context.Background(), strconv.FormatInt(u.Id, 10)).Scan(online)
	if online.Token == "" {
		response.OkWithDetailed(gin.H{"token": online.Token}, "用户已在线", c)
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
	ip := utils.GetLocalIp()
	online.Id = u.Id
	online.Username = u.Username
	online.LoginTime = time.Now().Unix()
	online.LoginLocation = utils.Ip2LocationCity(ip)
	online.Ip = ip
	online.Token = token

	marshal, err := json.Marshal(online)
	if err != nil {
		return
	}

	parseToken, _ := utils.ParseToken(token)
	var num int64
	num = (parseToken.ExpiresAt - time.Now().Unix()) * 1000 * 1000 * 1000
	onlineErr := global.ASS_REDIS.Set(context.Background(), strconv.FormatInt(u.Id, 10), marshal, time.Duration(num)).Err()
	if onlineErr != nil {
		response.FailWithMessage("存入redis错误", c)
		return
	}

	response.OkWithDetailed(models.LoginRes{
		Username: u.Username,
		Token:    token,
	}, "登陆成功", c)
	global.ASS_LOG.Info("用户:" + u.Username + ",登录成功")

}

// Logout 登出,h
func Logout(c *gin.Context) {
	id := c.GetInt64("id")
	a, _ := c.Get("onlineUser")
	v := time.Duration(c.GetInt64("expiresAt"))
	err := global.ASS_REDIS.Set(context.Background(), strconv.FormatInt(id, 10), a, v).Err()
	if err != nil {
		response.FailWithMessage("下线失败,请联系管理员", c)
		return
	}
	response.OkWithMessage("已下线", c)
}

func OnlineNumber(c *gin.Context) {

}
