package entity

// UserScore 用户积分
type UserScore struct {
	UserId   int32 `gorm:"primaryKey; not null; comment:主键,用户ID"`
	Score    int64 `gorm:"not null; default:0; comment:积分"`
	ScoreBox int64 `gorm:"not null; default:0; comment:保险箱积分"`
}
