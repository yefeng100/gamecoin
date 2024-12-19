package lobby

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"project/constants"
	"project/logic/baseconfig/configvip"
	"strings"
)

func Start(app pitaya.Pitaya) {
	h := NewHandlerUser(app)
	app.Register(h, component.WithName(constants.HandlerModuleLobby), component.WithNameFunc(strings.ToLower))

	configvip.Ins()
}
