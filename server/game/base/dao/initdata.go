package dao

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/logger"
	"project/base/dao/daobase"
	"project/base/entity"
	"project/modules/mysqlstorage"
	"time"
)

func InitTable() {
	//账号库
	accountDb := mysqlstorage.GetAcc()
	migratorAcc := accountDb.Migrator()
	if !migratorAcc.HasTable(&entity.UserIdPool{}) {
		_ = migratorAcc.AutoMigrate(&entity.UserIdPool{}) //用户ID池子
	}
	if !migratorAcc.HasTable(&entity.UserAccount{}) {
		_ = migratorAcc.AutoMigrate(&entity.UserAccount{}) //账号
	}
	if !migratorAcc.HasTable(&entity.UserScore{}) {
		_ = migratorAcc.AutoMigrate(&entity.UserScore{}) //用户金币
	}
	if !migratorAcc.HasTable(&entity.ConfigVip{}) {
		_ = migratorAcc.AutoMigrate(&entity.ConfigVip{}) //vip基础经验表
	}
	if !migratorAcc.HasTable(&entity.ConfigExp{}) {
		_ = migratorAcc.AutoMigrate(&entity.ConfigExp{}) //vip基础经验表
	}
	if !migratorAcc.HasTable(&entity.HomeRoll{}) {
		_ = migratorAcc.AutoMigrate(&entity.HomeRoll{}) //首页滚动数据
	}
	//日志库
	logDb := mysqlstorage.GetLog()
	migratorLog := logDb.Migrator()
	if !migratorLog.HasTable(&entity.UserScoreLog{}) {
		_ = migratorLog.AutoMigrate(&entity.UserScoreLog{}) //金币修改日志
	}
	//初始化数据
	initData()
	//定时器
	dbTask()
}

func initData() {
	daobase.CreateUserIdPool() //用户ID池子
	daobase.InitConfigVip()    //初始化VipLevel
	daobase.InitConfigExp()    //初始化ExpLevel
}

// 任务
func dbTask() {
	pitaya.NewTimer(time.Hour, func() {
		logger.Log.Info("dbTask 进入DB任务定时器")
		daobase.CreateUserIdPool() //用户ID池子
	})
}
