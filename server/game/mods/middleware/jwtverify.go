package middleware

import (
	"github.com/topfreegames/pitaya/v2/logger"
	"project/constants"
)

// JwtVerify jwt验证
func JwtVerify(jwtToken string) (int32, int32) {
	if len(jwtToken) == 0 {
		return constants.ResCodeJwtTokenErr, 0
	}
	jwt := NewJWT()
	// 检查token
	err := jwt.CheckToken(jwtToken)
	if err != nil {
		logger.Log.Errorf("CheckToken token error, jwtToken:%v, err:%v", jwtToken, err)
		return constants.ResCodeJwtTokenErr, 0
	}

	claims, err := jwt.ParserToken(jwtToken)
	if err != nil {
		logger.Log.Errorf("ParserToken token error, jwtToken:%v, err:%v", jwtToken, err)
		return constants.ResCodeJwtTokenErr, 0
	}
	// 刷新jwt过期时间
	_, err = jwt.UpdateToken(jwtToken)
	if err != nil {
		logger.Log.Errorf("UpdateToken token error, jwtToken:%v, err:%v", jwtToken, err)
	}
	return constants.ResCodeSuc, claims.UserId
}
