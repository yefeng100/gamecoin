package daobase

import (
	"github.com/topfreegames/pitaya/v2/logger"
	"gorm.io/gorm"
	"project/base/entity"
	"project/constants"
	"project/modules/mysqlstorage"
)

// CreateUserIdPool 插入UserID数据
func CreateUserIdPool() {
	accountDb := mysqlstorage.GetAcc()
	if accountDb == nil {
		return
	}
	maxCount := int32(100000) //保留10万用户ID
	count := int64(0)
	accountDb.Model(&entity.UserIdPool{}).Where("is_use = ?", 0).Count(&count)
	if count >= int64(maxCount) {
		return
	}
	userPool := new(entity.UserIdPool)
	accountDb.Select("max(user_id) as UserId").First(userPool)
	beginUserId := userPool.UserId + 1
	if userPool.UserId < 10000 {
		beginUserId = 10000
	}
	//插入数据
	for i := beginUserId; i < beginUserId+maxCount; i++ {
		userPool.UserId = i
		userPool.IsUse = false
		tx := accountDb.Create(userPool)
		if tx.Error != nil {
			logger.Log.Errorf("accountDb Create err:%v", tx.Error)
			return
		}
	}
}

// GetNewUserId 获取新的用户ID， 创建用户用
func GetNewUserId(conn *gorm.DB) (int32, error) {
	if conn == nil {
		conn = mysqlstorage.GetAcc()
		if conn == nil {
			return 0, constants.ErrConnectionFailure
		}
	}
	var uidData entity.UserIdPool
	tx := conn.Where("is_use = ?", 0).Order("RAND()").Take(&uidData)
	if tx.Error != nil {
		logger.Log.Errorf("GetNewUserId Take err:%v", tx.Error)
		return 0, tx.Error
	}
	uidData.IsUse = true
	tx = conn.Model(&entity.UserIdPool{}).Where("user_id = ?", uidData.UserId).Update("is_use", true)
	if tx.Error != nil {
		logger.Log.Errorf("GetNewUserId Save err:%v, uidData:%v", tx.Error, uidData)
		return 0, tx.Error
	}
	logger.Log.Infof("GetNewUserId %v", uidData)
	return uidData.UserId, nil
}

// SetNewUserId 激活用户ID， 创建用户失败
func SetNewUserId(uid int32, isUse bool) {
	accountDb := mysqlstorage.GetAcc()
	if accountDb == nil {
		return
	}
	tx := accountDb.Model(&entity.UserIdPool{}).Where("user_id = ?", uid).Update("is_use", isUse)
	if tx.Error != nil {
		logger.Log.Errorf("SetNewUserId Save err:%v, uid:%v, isUse:%v", tx.Error, uid, isUse)
		return
	}
}
