package gw

import (
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"net/http"
	"project/common/cfg"
	"project/constants"
	"strings"
)

func NetAcceptor(builder *pitaya.Builder) {
	//tcp
	tcpPort := cfg.GetIns().GetsSvConf().GetInt32("listen.tcpport")
	if tcpPort > 0 {
		tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", tcpPort))
		builder.AddAcceptor(tcp)
	}
	//wss
	wssPort := cfg.GetIns().GetsSvConf().GetInt32("listen.wssport")
	if wssPort > 0 {
		wss := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", wssPort))
		builder.AddAcceptor(wss)
	}
}

func Start(app pitaya.Pitaya) {
	h := NewHandler(app)
	app.Register(h, component.WithName(constants.HandlerModuleGw), component.WithNameFunc(strings.ToLower))

	httpPort := cfg.GetIns().GetsSvConf().GetInt32("listen.httpport")
	if httpPort > 0 {
		http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
		http.HandleFunc("/web", WebEnter)
		go func() {
			_ = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
		}()
	}
}
