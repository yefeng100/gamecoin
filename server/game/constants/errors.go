package constants

import "errors"

var (
	ErrConnectionFailure = errors.New("db connection failure") //数据库连接失败
	ErrNilUserId         = errors.New("nil user id")           //userId错误
)
