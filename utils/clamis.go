package utils

import (
	"association/global"
	"association/modules/system/request"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	header := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(header, " ", 2)
	token := parts[1]
	claims, err := ParseToken(token)
	if err != nil {
		global.ASS_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在Authorization且claims是否为规定结构")
	}
	return claims, err
}
