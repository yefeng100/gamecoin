package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2/logger"
	"sync"
)

var contCfg *ContractCfg
var contOnce sync.Once

type ContractCfg struct {
	cfg *viper.Viper //合约配置
}

func NewContractCfg(confName string) *ContractCfg {
	contOnce.Do(func() {
		contCfg = &ContractCfg{}
		contCfg.init(confName)
	})
	return contCfg
}

// GetContIns 获取配置管理指针
func GetContIns() *ContractCfg {
	if contCfg == nil {
		NewContractCfg("contract")
	}
	return contCfg
}

// 初始化server配置
func (t *ContractCfg) init(confName string) {
	t.cfg = viper.New()
	t.cfg.SetConfigName(fmt.Sprintf("config/%s", confName))
	t.cfg.AddConfigPath(".")
	t.cfg.SetConfigType("yaml")
	err := t.cfg.ReadInConfig()
	if err != nil {
		logger.Log.Errorln("加载ContractCfg配置错误 path:config/server.toml")
		panic(err.Error())
	}

	logger.Log.Infof("加载配置完成 path:config/contract.yaml %v", t.cfg)
}

// GetConf 获取系统配置
func (t *ContractCfg) GetConf() *viper.Viper {
	return t.cfg
}
