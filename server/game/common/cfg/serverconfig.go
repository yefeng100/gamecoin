package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/logger"
	"sync"
)

// ServerConfigData 基础配置数据
type ServerConfigData struct {
}

type ServerConfig struct {
	serverViper *viper.Viper //server配置
}

var instance *ServerConfig
var once sync.Once

func NewServerConfig(confName string) *ServerConfig {
	once.Do(func() {
		instance = &ServerConfig{}
		instance.init(confName)
	})
	return instance
}

// GetIns 获取配置管理指针
func GetIns() *ServerConfig {
	if instance == nil {
		NewServerConfig("server")
	}
	return instance
}

// 初始化server配置
func (t *ServerConfig) init(confName string) {
	t.serverViper = viper.New()
	t.serverViper.SetConfigName(fmt.Sprintf("config/%s", confName))
	t.serverViper.AddConfigPath(".")
	t.serverViper.SetConfigType("yaml")
	err := t.serverViper.ReadInConfig()
	if err != nil {
		logger.Log.Errorln("加载server配置错误 path:config/server.toml")
		panic(err.Error())
	}

	logger.Log.Infof("加载配置完成 path:config/server.toml %v", t.serverViper)
}

// GetsSvConf 获取系统配置
func (t *ServerConfig) GetsSvConf() *viper.Viper {
	return t.serverViper
}

// LoadServerConf 加载系统配置
func (t *ServerConfig) LoadServerConf(svType string) *config.Config {
	conf := viper.New()
	conf.Set("pitaya.cluster.rpc.client.nats.connect", t.serverViper.GetString("nats.url"))
	conf.Set("pitaya.cluster.rpc.server.nats.connect", t.serverViper.GetString("nats.url"))
	conf.Set("pitaya.cluster.sd.etcd.endpoints", t.serverViper.GetString("etcd.url"))
	conf.Set("pitaya.groups.etcd.endpoints", t.serverViper.GetString("etcd.url"))
	conf.Set("pitaya.modules.bindingstorage.etcd.endpoints", t.serverViper.GetString("etcd.url"))
	conf.Set("pitaya.handler.messages.compression", false) //消息是否压缩

	//conf.Set("pitaya.concurrency.handler.dispatch", 1) //携程数量

	configSys := config.NewConfig(conf)
	return configSys
}
