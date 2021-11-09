package util

import (
	"time"

	config "acgfate/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UID uint64 `json:"uid"`
	jwt.StandardClaims
}

// GenToken generate JWT token
func GenToken(uid uint64) (string, error) {
	conf := config.Conf.JWTConf

	expireDuration := time.Hour * time.Duration(conf.ExpireDuration)

	c := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(), // expire time
			Issuer:    "ACG.Fate",                            // 签发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(conf.Secret))
}

// ParseToken Parse JWT token
func ParseToken(tokenString string) (*Claims, error) {
	conf := config.Conf.JWTConf

	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(conf.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // validate token
		return claims, nil
	}
	return nil, err
}
