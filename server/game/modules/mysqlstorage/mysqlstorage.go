package mysqlstorage

import (
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/interfaces"
	"github.com/topfreegames/pitaya/v2/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/common/cfg"
	"project/constants"
	"time"
)

// GetAcc 账号库
func GetAcc() *gorm.DB {
	mysqlModule, err := pitaya.GetModule(constants.ModuleMysqlStorage)
	if err != nil {
		logger.Log.Errorf("GetAcc GetModule err:%v", err)
		return nil
	}
	if mysqlModule == nil {
		logger.Log.Errorf("GetAcc mysqlModule is nil")
		return nil
	}
	return mysqlModule.(*MysqlModule).AccountDB
}

// GetLog 日志库
func GetLog() *gorm.DB {
	mysqlModule, err := pitaya.GetModule(constants.ModuleMysqlStorage)
	if err != nil {
		logger.Log.Errorf("GetLog GetModule err:%v", err)
		return nil
	}
	if mysqlModule == nil {
		logger.Log.Errorf("GetLog mysqlModule is nil")
		return nil
	}
	return mysqlModule.(*MysqlModule).LogDB
}

type MysqlModule struct {
	interfaces.Module
	AccountDB *gorm.DB
	LogDB     *gorm.DB
}

func NewMysqlStorage(app pitaya.Pitaya) *MysqlModule {
	p := new(MysqlModule)
	p.start()

	err := app.RegisterModule(p, constants.ModuleMysqlStorage)
	if err != nil {
		logger.Log.Fatal("mysql add module fatal")
		return nil
	}
	//初始化表
	//go dao.InitTable()
	return p
}

// 启动Mysql
func (p *MysqlModule) start() {
	ip := cfg.GetIns().GetsSvConf().GetString("mysql.addr")
	port := cfg.GetIns().GetsSvConf().GetInt32("mysql.port")
	user := cfg.GetIns().GetsSvConf().GetString("mysql.user")
	pwd := cfg.GetIns().GetsSvConf().GetString("mysql.pwd")
	poolMax := cfg.GetIns().GetsSvConf().GetInt("mysql.pool")
	if poolMax < 100 || poolMax > 500 {
		poolMax = 100
	}

	var err error = nil
	{
		dbName := cfg.GetIns().GetsSvConf().GetString("mysql.dbacc")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, ip, port, dbName)
		p.AccountDB, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			//Logger: gormLog.Default.LogMode(gormLog.Info),
		})
		if err != nil {
			logger.Log.Fatalf("mysql init fatal! %s", dsn)
			return
		}
		sqlDB, err := p.AccountDB.DB()
		if err != nil {
			logger.Log.Fatalf("mysql AccountDB.DB fatal! %s, %v", dsn, err)
			return
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(poolMax)
		// SetConnMaxIdleTime 设置了空闲连接最大时间。
		//sqlDB.SetConnMaxIdleTime(time.Hour)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(24 * time.Hour)
		// ping
		err = sqlDB.Ping()
		if err != nil {
			logger.Log.Fatalf("mysql Ping fatal! %s", dsn)
			return
		}
	}
	{
		dbName := cfg.GetIns().GetsSvConf().GetString("mysql.dblog")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, ip, port, dbName)
		p.LogDB, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			//Logger: gormLog.Default.LogMode(gormLog.Info),
		})
		if err != nil {
			logger.Log.Fatalf("mysql init fatal! %s", dsn)
			return
		}
		sqlDB, err := p.LogDB.DB()
		if err != nil {
			logger.Log.Fatalf("mysql LogDB.DB fatal! %s, %v", dsn, err)
			return
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(poolMax)
		// SetConnMaxIdleTime 设置了空闲连接最大时间。
		//sqlDB.SetConnMaxIdleTime(time.Hour)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
		// ping
		err = sqlDB.Ping()
		if err != nil {
			logger.Log.Fatalf("mysql Ping fatal! %s", dsn)
			return
		}
	}
}

func (p *MysqlModule) Init() error {
	logger.Log.Infof("mysql module init")
	return nil
}
func (p *MysqlModule) AfterInit() {
	logger.Log.Infof("mysql module AfterInit")
}
func (p *MysqlModule) BeforeShutdown() {
	logger.Log.Infof("mysql module BeforeShutdown")
}
func (p *MysqlModule) Shutdown() error {
	logger.Log.Infof("mysql module Shutdown")
	return nil
}
