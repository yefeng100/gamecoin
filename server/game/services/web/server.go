package web

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"project/constants"
	"project/logic/baseconfig/configvip"
	"strings"
)

func Start(app pitaya.Pitaya) {
	hr := NewHandlerRemote(app)
	app.RegisterRemote(hr, component.WithName(constants.HandlerModuleWeb), component.WithNameFunc(strings.ToLower))

	configvip.Ins()
}
