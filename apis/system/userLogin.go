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
	u, loginErr := userLoginService.LoginUser(*loginUser)
	if loginErr != nil {
		global.ASS_LOG.Error("login failed username:" + loginUser.Username)
		response.FailWithMessage("账号或密码错误", c)
		return
	}
	online := new(models.OnlineUser)

	bytes, _ := global.ASS_REDIS.Get(context.Background(), strconv.FormatUint(uint64(u.ID), 10)).Bytes()
	_ = json.Unmarshal(bytes, online)
	u.Password = "******"
	if online.Token != "" {
		response.OkWithDetailed(gin.H{
			"User":  u,
			"token": online.Token,
		}, "登陆成功", c)
		global.ASS_LOG.Info("用户id:" + strconv.FormatUint(uint64(u.ID), 10) + ",登录成功")
		return
	}

	claims := utils.CreateClaims(request.BaseClaims{
		ID:        int64(u.ID),
		Username:  u.Username,
		Authority: u.StatusId,
	}, time.Now().Add(time.Hour*24*7).Unix())
	token, err := utils.GenToken(claims)
	if err != nil {
		global.ASS_LOG.Error("获取token失败,error:" + err.Error())
		response.FailWithMessage("获取token失败", c)
		return
	}
	ip := utils.GetLocalIp()
	online.Id = u.ID
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
	onlineErr := global.ASS_REDIS.Set(context.Background(), strconv.FormatUint(uint64(u.ID), 10), marshal, time.Duration(num)).Err()
	if onlineErr != nil {
		response.FailWithMessage("存入redis错误", c)
		return
	}

	response.OkWithDetailed(gin.H{
		"User":  u,
		"token": token,
	}, "登陆成功", c)
	global.ASS_LOG.Info("用户id:" + strconv.FormatUint(uint64(u.ID), 10) + ",登录成功")

}

// LogoutU 登出
func LogoutU(c *gin.Context) {
	id := c.GetInt64("id")
	err := global.ASS_REDIS.Get(context.Background(), strconv.FormatInt(id, 10)).Err()
	if err != nil {
		response.FailWithMessage("下线失败,请联系管理员", c)
		return
	}
	response.OkWithMessage("已下线", c)
}
