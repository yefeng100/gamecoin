package entity

// UserIdPool 用户ID池子
type UserIdPool struct {
	UserId int32 `gorm:"primaryKey; not null; comment:主键,用户ID"`
	IsUse  bool  `gorm:"not null; default:0; index; comment:是否使用"`
}
