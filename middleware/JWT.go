package middleware

import (
	"association/common/response"
	"association/modules/system/request"
	"association/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		//续签
		a := mc.ExpiresAt - time.Now().Unix()
		b := int64(86400000)
		if a < b {
			claims := utils.CreateClaims(request.BaseClaims{
				ID:        mc.BaseClaims.ID,
				Username:  mc.BaseClaims.Username,
				Authority: mc.BaseClaims.Authority,
			}, time.Now().Add(time.Hour*8).Unix())
			token, err := utils.GenToken(claims)
			if err != nil {
				response.FailWithMessage("续签token失败", c)
				return
			}
			contextSet(c, mc)
			c.Request.Header.Set("Authorization", token)
			c.Next()
		}
		// 将当前请求的username信息保存到请求的上下文c上
		contextSet(c, mc)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func contextSet(c *gin.Context, mc *request.CustomClaims) {
	c.Set("username", mc.BaseClaims.Username)
	c.Set("id", mc.BaseClaims.ID)
	c.Set("authority", mc.BaseClaims.Authority)
	c.Set("expiresAt", mc.ExpiresAt)
}
