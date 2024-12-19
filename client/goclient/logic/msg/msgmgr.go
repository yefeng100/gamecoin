package msg

import (
	"encoding/json"
	"fmt"
	"framework/log"
	"go.uber.org/zap"
	"goclient/logic/nn"
	"goclient/modules/conf"
	"goclient/modules/net"
	"os"
)

var conn *net.ConnServer

func NewMsgMgr() {
	conn = net.NewConnectServer()
	ip := conf.GetIns().GetString("net.addr")
	port := conf.GetIns().GetInt("net.port")
	addr := fmt.Sprintf("%s:%d", ip, port)
	err := conn.ConnectServer(addr)
	if err != nil {
		log.Error("连接失败", zap.String("addr", addr), zap.Error(err))
		os.Exit(0)
		return
	}
	conn.SetCallback(HandlerData)

	type TestMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	tst := &TestMessage{}
	tst.Content = "aaaaa"
	tst.Name = "bbbbb"
	jt, _ := json.Marshal(tst)
	conn.NotifyMsg("gw.gw.testmessage", jt)
	//
}

func ConnectNN(addr string) error {
	err := conn.ConnectServer(addr)
	if err != nil {
		log.Error("连接失败", zap.String("addr", addr), zap.Error(err))
		return err
	}
	l := nn.NewLogicNN()
	conn.SetCallback(l.HandlerData)
	return nil
}
