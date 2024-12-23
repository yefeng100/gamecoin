package constants

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2/logger"
)

var (
	cn map[string]string
	en map[string]string
)

func init() {
	languageCn()
	languageEn()
}

func languageCn() {
	cnConf := viper.New()
	cnConf.SetConfigName("config/language_cn")
	cnConf.AddConfigPath(".")
	cnConf.SetConfigType("yaml")
	err := cnConf.ReadInConfig()
	if err != nil {
		logger.Log.Errorln("加载server配置错误 path:config/server.toml")
		panic(err.Error())
	}
	cn = cnConf.GetStringMapString("language")
}

func languageEn() {
	cnConf := viper.New()
	cnConf.SetConfigName("config/language_en")
	cnConf.AddConfigPath(".")
	cnConf.SetConfigType("yaml")
	err := cnConf.ReadInConfig()
	if err != nil {
		logger.Log.Errorln("加载server配置错误 path:config/server.toml")
		panic(err.Error())
	}
	en = cnConf.GetStringMapString("language")
}

func GetCodeMsg(languageCode string, code int32) string {
	key := fmt.Sprintf("code_%d", code)
	switch languageCode {
	case LanguageCodeEn:
		msg, ok := en[key]
		if !ok {
			return ""
		}
		return msg
	case LanguageCodeCn:
		msg, ok := cn[key]
		if !ok {
			return ""
		}
		return msg
	}
	//默认英文
	msg, ok := en[key]
	if !ok {
		return ""
	}
	return msg
}
