package daobase

import (
	"errors"
	"github.com/topfreegames/pitaya/v2/logger"
	"gorm.io/gorm"
	"project/base/entity"
	"project/constants"
	"project/modules/mysqlstorage"
	"strings"
)

// CreateUserAccount 账号创建
func CreateUserAccount(conn *gorm.DB, acc *entity.UserAccount) error {
	if conn == nil {
		conn = mysqlstorage.GetAcc()
		if conn == nil {
			return constants.ErrConnectionFailure
		}
	}
	tx := conn.Create(acc)
	if tx.Error != nil {
		logger.Log.Errorf("CreateUserAccount acc:%v, err:%v", acc, tx.Error)
		return tx.Error
	}
	return nil
}

// GetUserAccByAccName 账号查找
func GetUserAccByAccName(accName string) *entity.UserAccount {
	accConn := mysqlstorage.GetAcc()
	if accConn == nil {
		return nil
	}
	retData := &entity.UserAccount{}
	tx := accConn.Where("acc_name = ?", accName).Take(retData)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			logger.Log.Errorf("GetUserAccByAccName Take err:%v", tx.Error)
		}
		return nil
	}
	return retData
}

// GetUserAccByPhone 手机号查找
func GetUserAccByPhone(phone string) *entity.UserAccount {
	accConn := mysqlstorage.GetAcc()
	if accConn == nil {
		return nil
	}
	retData := &entity.UserAccount{}
	tx := accConn.Where("phone = ?", phone).Take(retData)
	if tx.Error != nil {
		logger.Log.Errorf("GetUserAccByPhone Take err:%v", tx.Error)
		return nil
	}
	return retData
}

// GetUserAccById userId查找
func GetUserAccById(userId int32) *entity.UserAccount {
	accConn := mysqlstorage.GetAcc()
	if accConn == nil {
		return nil
	}
	retData := &entity.UserAccount{}
	tx := accConn.First(retData, userId)
	if tx.Error != nil {
		logger.Log.Errorf("GetUserAccById err:%v", tx.Error)
		return nil
	}
	return retData
}

// GetUserAccCountByMachine 机器码数量
func GetUserAccCountByMachine(machine string) int64 {
	accConn := mysqlstorage.GetAcc()
	if accConn == nil {
		return 0
	}
	machine = strings.TrimSpace(machine)
	if machine == "" {
		return 0
	}
	count := int64(0)
	tx := accConn.Model(&entity.UserAccount{}).Where("machine = ?", machine).Count(&count)
	if tx.Error != nil {
		logger.Log.Errorf("GetUserAccCountByMachine Take err:%v", tx.Error)
		return count
	}
	return count
}

// UpdUserAccByFaceUrl 修改头像
func UpdUserAccByFaceUrl(userId int64, faceUrl string) {
	accConn := mysqlstorage.GetAcc()
	if accConn == nil {
		return
	}
	tx := accConn.Model(&entity.UserAccount{}).Where("user_id = ?", userId).Update("face_url", faceUrl)
	if tx.Error != nil {
		logger.Log.Errorf("UpdUserAccByFaceUrl Take err:%v", tx.Error)
	}
}
