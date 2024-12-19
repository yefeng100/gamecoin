package daobase

import (
	"errors"
	"github.com/topfreegames/pitaya/v2/logger"
	"gorm.io/gorm"
	"project/base/entity"
	"project/modules/mysqlstorage"
)

// GetHomeRollList 获取首页滚动列表
func GetHomeRollList() []entity.HomeRoll {
	list := make([]entity.HomeRoll, 0)
	dataDb := mysqlstorage.GetAcc()
	if dataDb == nil {
		return list
	}
	err := dataDb.Find(&list).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.Errorf("数据库获取首页滚动列表失败 err:%v", err)
		}
	}

	return list
}
