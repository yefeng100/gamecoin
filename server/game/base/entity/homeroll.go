package entity

import "time"

// HomeRoll 首页滚动数据
type HomeRoll struct {
	ID          int64     `gorm:"primaryKey; autoIncrement; comment:主键"`
	Img         string    `gorm:"size:128; default:''; comment:图片(有http为url地址,否则为本地图片名字)"`
	Url         string    `gorm:"size:128; default:''; comment:跳转连接,点击图片打开连接(url为空,点击图片不做任何反应)"`
	UrlOpenType int8      `gorm:"not null; default:0; comment:跳转连接打开方式(0默认跳转App页面,1浏览器打开URL)"`
	Valid       int8      `gorm:"not null; default:0; comment:是否有效(0:有效,1:无效)"`
	CreatedAt   time.Time `gorm:"autoCreateTime; not null; <-:create; default:CURRENT_TIMESTAMP(3); comment:创建时间"`
}
