package web

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"project/common/cfg"
	"project/constants"
	"project/logic/baseconfig/configvip"
	"strings"
)

func Start(app pitaya.Pitaya) {
	hr := NewHandlerRemote(app)
	app.RegisterRemote(hr, component.WithName(constants.HandlerModuleWeb), component.WithNameFunc(strings.ToLower))

	onOff := cfg.GetContIns().GetConf().GetInt("bsc.onoff")
	if onOff == 1 {
		//初始化合约
		//bscspendcoin.GetIns() //授权
		//test
		//bscspendcoin.SpendUserCoin("0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "0x47B8EcBA4fdaaF1D47cd60920B9834B333548E51", "0x057188BEe8C920b2F6FB677F035C4EaF19157D31", 200000)
		//bscspendcoin.Owner()
		//erc20.SubscribeContract("0x337610d27c682E347C9cD60BD4b3b107C9d34dDd")
		//trcspendcoin.CollectTokens(20000)
		//trcspendcoin.WithdrawTokens()
	}
	//初始化vip
	configvip.Ins()
}
