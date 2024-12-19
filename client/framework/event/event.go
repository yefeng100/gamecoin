package event

import (
	"framework/constants"
	"framework/msgsrv"
	"framework/net"
)

type Event struct {
	Name string
	Args map[string]interface{}
}

// NewConnectEvt 网络连接事件
func NewConnectEvt(name string, conn *net.WssConn) *Event {
	e := &Event{}
	e.Name = name
	e.Args = make(map[string]interface{})
	e.Args[constants.MsgEvtNameClientConn] = conn
	return e
}

// NewDisconnectEvt 网络断开事件
func NewDisconnectEvt(name string, conn *net.WssConn) *Event {
	e := &Event{}
	e.Name = name
	e.Args = make(map[string]interface{})
	e.Args[constants.MsgEvtNameClientConn] = conn
	return e
}

// NewMsgEvt 网络消息事件
func NewMsgEvt(name string, conn *net.WssConn, msg *msgsrv.Message) *Event {
	e := &Event{}
	e.Name = name
	e.Args = make(map[string]interface{})
	e.Args[constants.MsgEvtNameClientConn] = conn
	e.Args[constants.MsgEvtNameClientMsg] = msg
	return e
}
