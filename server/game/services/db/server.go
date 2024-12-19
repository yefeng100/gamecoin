package dbserver

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"project/constants"
	"project/logic/baseconfig/configvip"
	"strings"
)

func Start(app pitaya.Pitaya) {
	hr := NewHandlerRemote(app)
	app.RegisterRemote(hr, component.WithName(constants.HandlerModuleDb), component.WithNameFunc(strings.ToLower))

	configvip.Ins()
}
