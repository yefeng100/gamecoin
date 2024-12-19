package entity

import "time"

// UserScoreLog 用户金币日志
// TypeID:0未知,1充值,2提现,3银行取出,4银行存入,10000以上是游戏服务器ID
type UserScoreLog struct {
	ID          int64     `gorm:"primaryKey; autoIncrement; comment:主键"`
	UserId      int32     `gorm:"not null; default:0; comment:用户ID"`
	CurScore    int64     `gorm:"not null; default:0; comment:当前金币"`
	UpdScore    int64     `gorm:"not null; default:0; comment:修改金币"`
	AlterScore  int64     `gorm:"not null; default:0; comment:最终金币"`
	TypeId      int       `gorm:"not null; default:0; comment:类型:0未知,1充值,2提现,3银行取出,4银行存入,10000跑得快初级输赢"`
	Description string    `gorm:"size:128; comment:说明:跑得快初级场输赢"`
	CreatedAt   time.Time `gorm:"autoCreateTime; not null; <-:create; default:CURRENT_TIMESTAMP(3); comment:创建时间"`
}
