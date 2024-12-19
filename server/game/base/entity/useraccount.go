package entity

import "time"

// UserAccount 账号主表
type UserAccount struct {
	UserId    int32     `gorm:"primaryKey; not null; comment:主键,用户ID"` // 设置为主键，会自动为自增字段
	AccName   string    `gorm:"size:32; not null; uniqueIndex; comment:账号"`
	Pwd       string    `gorm:"size:64; default:''; comment:密码"`
	NickName  string    `gorm:"size:64; default:''; comment:昵称"`
	FaceUrl   string    `gorm:"size:128; default:''; comment:头像"`
	Machine   string    `gorm:"size:64; not null; index; comment:机器码"`
	AccType   int8      `gorm:"not null; default:0; comment:账号类型(0:普通玩家,1:机器人,2:测试账号)"`
	RegPlat   int8      `gorm:"not null; default:0; comment:注册平台(0:未知,1:安卓,2:IOS,3:WEB,4:PC)"`
	PhoneNum  string    `gorm:"size:32; default:''; comment:手机号(国际编码-手机号)"`
	CreatedAt time.Time `gorm:"autoCreateTime; not null; <-:create; default:CURRENT_TIMESTAMP(3); comment:创建时间"`
}
