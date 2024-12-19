package main

import (
	"flag"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/logger"
	"project/base/dao"
	"project/common/cfg"
	"project/constants"
	"project/mods/logfile"
	"project/modules/mysqlstorage"
	"project/modules/redisstorage"
	"project/services/db"
	"project/services/gw"
	"project/services/lobby"
	"project/services/web"
)

var app pitaya.Pitaya

func main() {
	//--type=login --areaId=0
	svType := flag.String("type", "", "the server type")
	areaId := flag.Int("areaid", 0, "the server areaId")
	logLevel := flag.Int("loglevel", 6, "the server loglevel")
	confName := flag.String("confname", "server", "the server main config")
	flag.Parse()
	//加载日志
	logfile.LogLoad(*svType, *logLevel, true)
	//加载配置
	cfg.NewServerConfig(*confName)
	conf := cfg.GetIns().LoadServerConf(*svType)
	//开始
	start(*svType, *areaId, conf)
}

// 开始
func start(svType string, areaId int, conf *config.Config) {
	builder := pitaya.NewBuilderWithConfigs(svType == constants.ServerNameGw, svType, pitaya.Cluster, map[string]string{}, conf)
	if svType == constants.ServerNameGw {
		gw.NetAcceptor(builder)
	}
	app = builder.Build()
	defer app.Shutdown()
	pitaya.DefaultApp = app

	switch svType {
	case constants.ServerNameGw:
		//网关
		gw.Start(app) //初始化业务逻辑
	case constants.ServerNameLobby:
		//大厅
		lobby.Start(app)
	case constants.ServerNameDb:
		//db
		mysqlStorage()      //初始化mysql
		redisStorage()      //初始化redis
		dbserver.Start(app) //初始化业务逻辑
	case constants.ServerNameWeb:
		//web
		mysqlStorage() //初始化mysql
		redisStorage() //初始化redis
		web.Start(app) //初始化业务逻辑
	default:
		logger.Log.Errorf("启动失败, svType err： svType:%s, areaId:%v", svType, areaId)
		logger.Log.Fatal("启动失败")
		return
	}

	logger.Log.Infof("启动成功： svType:%s, areaId:%v", svType, areaId)
	app.Start()
}

func mysqlStorage() {
	mysqlstorage.NewMysqlStorage(app) //初始化mysql
	go dao.InitTable()
}

func redisStorage() {
	redisstorage.NewRedisStorage(app) //初始化redis
}
