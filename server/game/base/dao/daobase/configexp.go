package daobase

import (
	"github.com/topfreegames/pitaya/v2/logger"
	"project/base/entity"
	"project/modules/mysqlstorage"
)

func InitConfigExp() []*entity.ConfigExp {
	logger.Log.Warnf("InitConfigExp 初始化ConfigExp表")
	list := make([]*entity.ConfigExp, 0)

	dataDb := mysqlstorage.GetAcc()
	if dataDb == nil {
		return list
	}
	count := int64(0)
	dataDb.Model(&entity.ConfigVip{}).Count(&count)
	if count > 0 {
		return list
	}
	list = append(list, &entity.ConfigExp{Level: 1, Exp: 3600})
	list = append(list, &entity.ConfigExp{Level: 2, Exp: 3600 * 2})
	list = append(list, &entity.ConfigExp{Level: 3, Exp: 3600 * 4})
	list = append(list, &entity.ConfigExp{Level: 4, Exp: 3600 * 8})
	list = append(list, &entity.ConfigExp{Level: 5, Exp: 3600 * 16})
	list = append(list, &entity.ConfigExp{Level: 6, Exp: 3600 * 32})
	list = append(list, &entity.ConfigExp{Level: 7, Exp: 3600 * 64})
	list = append(list, &entity.ConfigExp{Level: 8, Exp: 3600 * 128})
	list = append(list, &entity.ConfigExp{Level: 9, Exp: 3600 * 256})
	list = append(list, &entity.ConfigExp{Level: 10, Exp: 3600 * 512})
	for _, v := range list {
		SetExpLevel(v.Level, v.Exp)
	}
	return list
}

// SetExpLevel 插入ConfigExp
func SetExpLevel(level int32, exp int64) {
	dataDb := mysqlstorage.GetAcc()
	if dataDb == nil {
		return
	}
	vipLevel := &entity.ConfigExp{
		Level: level,
		Exp:   exp,
	}
	tx := dataDb.Create(vipLevel)
	if tx.Error != nil {
		logger.Log.Errorf("数据库插入ConfigExp失败 err:%v", tx.Error)
	}
}

// GetConfigExpList 获取ConfigVip列表
func GetConfigExpList() []*entity.ConfigExp {
	list := make([]*entity.ConfigExp, 0)
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
