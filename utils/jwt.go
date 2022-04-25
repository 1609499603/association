package utils

import (
	"association/global"
	"association/modules/system/request"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	// 创建一个我们自己的声明
	c := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.ASS_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Issuer:    global.ASS_CONFIG.JWT.Issuer,
		},
	}
	return c
}

// GenToken 生成JWT
func GenToken(c request.CustomClaims) (string, error) {
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(global.ASS_CONFIG.JWT.SigningKey))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*request.CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(global.ASS_CONFIG.JWT.SigningKey), nil
	})
	if err != nil {
		global.ASS_LOG.Error("token error:" + err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
