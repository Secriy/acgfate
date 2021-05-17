package utils

import (
	"time"

	config "acgfate/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UID uint64 `json:"uid"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(uid uint64) (string, error) {
	// JWT配置
	jwtSecret := []byte(config.Conf.JWT.Secret)
	expireDuration := time.Hour * time.Duration(config.Conf.JWT.ExpireDuration)
	// 生成JWT字段
	c := Claims{
		uid, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(), // Expire time
			Issuer:    "ACGFATE",                             // 签发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// JWT配置
	jwtSecret := []byte(config.Conf.JWT.Secret)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, err
}
