package cache

import "time"

// redis expiration time
const (
	RRExpirationSecond30 = time.Second * 30 //30秒超时
	RExpirationHour1     = time.Hour        //1小时
	RExpirationDay1      = time.Hour * 24   //1天
)

// redis key
const (
	RKeyUserNonce  = "user:nonce:%s"  //随机数, %s=用户自己生成的唯一值
	RKeyLoginToken = "token:login:%d" //登录token, %d=UserId
)

// redis key(table)
const (
	RKeyUserAccountID = "table:useraccount:id:%d" //账号数据, %d=userId
	RKeyVipLevelList  = "table:viplevel"
)
