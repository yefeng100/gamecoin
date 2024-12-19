package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var cfg *viper.Viper = nil

func InitServer() (*viper.Viper, error) {
	cfg = viper.New()
	cfg.AddConfigPath("config") // 添加搜索路径
	cfg.SetConfigName("server")
	cfg.SetConfigType("toml")
	err := cfg.ReadInConfig()
	if err != nil {
		fmt.Println("InitConf err:", err)
		return nil, err
	}
	return cfg, nil
}

func GetIns() *viper.Viper {
	if cfg == nil {
		_, _ = InitServer()
	}
	return cfg
}
