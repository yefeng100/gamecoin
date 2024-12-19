package msgsrv

import (
	"framework/constants"
	"framework/event"
	"framework/log"
	"framework/net"
	"go.uber.org/zap"
)

type MsgSvr struct {
	msgChan *net.WssClientEventChan
	cliConn *net.WssClient
}

func NewMsgSvr(addr string) *MsgSvr {
	t := &MsgSvr{}
	t.msgChan = net.NewWssClientEventChan(constants.MsgChanEvtSize)
	t.cliConn = net.NewWssClient(t.msgChan, addr, true)
	return t
}

func (t *MsgSvr) Start() error {
	err := t.cliConn.Connect()
	if err != nil {
		log.Error("网络连接失败", zap.String("addr", t.cliConn.GetAddr()), zap.Error(err))
		log.Fatal("网络连接失败")
		return err
	}
	//读取网络数据
	go t.chanData()
	return nil
}

func (t *MsgSvr) chanData() {
	for {
		select {
		case conn := <-t.msgChan.ConnectedSig:
			//网络连接成功
			event.NewConnectEvt(constants.MsgEvtClientConn, conn)
		case conn := <-t.msgChan.DisconnectedSig:
			//断开网络连接
			event.NewDisconnectEvt(constants.MsgEvtNameClientDisConn, conn)
			if !t.cliConn.IsConnected() {
				//不重连，退出网络
				goto _exit
			}
		case data := <-t.msgChan.MsgSig:
			msg, err := UnMarshal(data)
			if err != nil {
				log.Error("wss 客户端解析数据错误", zap.String("addr", t.cliConn.GetAddr()), zap.Error(err))
			}
			//收到数据
			event.NewMsgEvt(constants.MsgEvtNameClientMsg, t.cliConn.GetConn(), msg)
		}
	}
_exit:
	log.Error("退出网络")
}
