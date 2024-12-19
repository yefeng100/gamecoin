package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"project/common/cfg"
	"time"
)

// token生成器
func GenerateToken(userId int32) (token string, err error) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := NewJWT()
	//生效时间
	notBefore := &jwt.NumericDate{}
	notBefore.Time = time.Now()
	//前面用户
	jwtUser := cfg.GetIns().GetsSvConf().GetString("server.jwtUser")
	//过期时间
	jwtTokenTime := cfg.GetIns().GetsSvConf().GetInt64("server.jwtTokenTime")
	expiresAt := &jwt.NumericDate{}
	expiresAt.Time = time.Now().Add(time.Second * time.Duration(jwtTokenTime))
	// 构造用户claims信息(负荷)
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: notBefore, // 签名生效时间
			ExpiresAt: expiresAt, // 签名过期时间
			Issuer:    jwtUser,   // 签名颁发者
		},
	}
	// 根据claims生成token对象
	token, err = j.CreateToken(claims)
	return
}
