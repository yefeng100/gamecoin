package daobase

import (
	"github.com/topfreegames/pitaya/v2/logger"
	"gorm.io/gorm"
	"project/base/entity"
	"project/constants"
	"project/modules/mysqlstorage"
)

// CreateUserScore 账号创建
func CreateUserScore(conn *gorm.DB, eData *entity.UserScore) error {
	if conn == nil {
		conn = mysqlstorage.GetAcc()
		if conn == nil {
			return constants.ErrConnectionFailure
		}
	}
	tx := conn.Create(eData)
	if tx.Error != nil {
		logger.Log.Errorf("CreateUserScore acc:%v, err:%v", eData, tx.Error)
		return tx.Error
	}
	return nil
}

func GetUserScoreById(userId int32) *entity.UserScore {
	accConn := mysqlstorage.GetAcc()
	if accConn == nil {
		logger.Log.Errorf("GetUserScoreById accConn is nil")
		return nil
	}
	retData := &entity.UserScore{}
	tx := accConn.First(retData, userId)
	if tx.Error != nil {
		logger.Log.Errorf("GetUserScoreById err:%v", tx.Error)
		return nil
	}
	return retData
}
