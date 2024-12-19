package daobase

import (
	"github.com/topfreegames/pitaya/v2/logger"
	"project/base/entity"
	"project/constants"
	"project/modules/mysqlstorage"
)

func InitConfigVip() []*entity.ConfigVip {
	logger.Log.Warnf("InitVipLevel 初始化VipLevel表")
	list := make([]*entity.ConfigVip, 0)

	dataDb := mysqlstorage.GetAcc()
	if dataDb == nil {
		return list
	}
	count := int64(0)
	dataDb.Model(&entity.ConfigVip{}).Count(&count)
	if count > 0 {
		return list
	}
	list = append(list, &entity.ConfigVip{Level: 1, Exp: 100 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 2, Exp: 200 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 3, Exp: 500 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 4, Exp: 1000 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 5, Exp: 2000 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 6, Exp: 5000 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 7, Exp: 10000 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 8, Exp: 20000 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 9, Exp: 50000 * constants.ScoreRatio})
	list = append(list, &entity.ConfigVip{Level: 10, Exp: 100000 * constants.ScoreRatio})
	for _, v := range list {
		SetVipLevel(v.Level, v.Exp)
	}
	return list
}

// SetVipLevel 插入vipLevel
func SetVipLevel(level int32, exp int64) {
	dataDb := mysqlstorage.GetAcc()
	if dataDb == nil {
		return
	}
	vipLevel := &entity.ConfigVip{
		Level: level,
		Exp:   exp,
	}
	tx := dataDb.Create(vipLevel)
	if tx.Error != nil {
		logger.Log.Errorf("数据库插入VipLevel失败 err:%v", tx.Error)
	}
}

// GetConfigVipList 获取ConfigVip列表
func GetConfigVipList() []*entity.ConfigVip {
	list := make([]*entity.ConfigVip, 0)
	dataDb := mysqlstorage.GetAcc()
	if dataDb == nil {
		return list
	}
	tx := dataDb.Find(&list)
	if tx.Error != nil {
		logger.Log.Errorf("数据库获取ConfigVip失败 err:%v", tx.Error)
		return list
	}
	return list
}
