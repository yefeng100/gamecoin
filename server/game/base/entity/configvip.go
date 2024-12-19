package entity

// ConfigVip Vip等级
// Exp是充值额度
type ConfigVip struct {
	Level int32 `gorm:"primaryKey; not null; comment:等级"`
	Exp   int64 `gorm:"not null; default:0; comment:vip经验"`
}
