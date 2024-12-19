package redisstorage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/interfaces"
	"github.com/topfreegames/pitaya/v2/logger"
	"project/common/cfg"
	"project/constants"
)

func GetConn() *redis.Client {
	mysqlModule, err := pitaya.GetModule(constants.ModuleRedisStorage)
	if err != nil {
		logger.Log.Errorf("GetConn RedisModule err:%v", err)
		return nil
	}
	if mysqlModule == nil {
		logger.Log.Errorf("GetConn RedisModule is nil")
		return nil
	}
	return mysqlModule.(*RedisModule).rdbCli
}

type RedisModule struct {
	interfaces.Module
	rdbCli *redis.Client
}

func NewRedisStorage(app pitaya.Pitaya) *RedisModule {
	p := new(RedisModule)
	p.start()

	err := app.RegisterModule(p, constants.ModuleRedisStorage)
	if err != nil {
		logger.Log.Fatal("redis add module fatal")
		return nil
	}
	return p
}

// 启动Redis
func (p *RedisModule) start() {
	ip := cfg.GetIns().GetsSvConf().GetString("redis.addr")
	port := cfg.GetIns().GetsSvConf().GetInt32("redis.port")
	pwd := cfg.GetIns().GetsSvConf().GetString("redis.pwd")
	db := cfg.GetIns().GetsSvConf().GetInt("redis.db")
	p.rdbCli = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ip, port),
		Password: pwd,
		DB:       db,
	})
	_, err := p.rdbCli.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("fatal error connect to redis: %s", err.Error()))
		return
	}
	logger.Log.Infof("connect redis suc")
}

func (p *RedisModule) Init() error {
	logger.Log.Infof("redis module init")
	return nil
}
func (p *RedisModule) AfterInit() {
	logger.Log.Infof("redis module AfterInit")
}
func (p *RedisModule) BeforeShutdown() {
	logger.Log.Infof("redis module BeforeShutdown")
}
func (p *RedisModule) Shutdown() error {
	logger.Log.Infof("redis module Shutdown")
	err := p.rdbCli.Close()
	if err != nil {
		logger.Log.Errorf("redis close fail! err:%v", err)
	}
	return nil
}
