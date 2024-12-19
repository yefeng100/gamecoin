package entity

// ConfigExp 在线时长等级
// Exp在线时长
type ConfigExp struct {
	Level int32 `gorm:"primaryKey; not null; comment:等级"`
	Exp   int64 `gorm:"not null; default:0; comment:在线时长"`
}
