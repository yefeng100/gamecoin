package net

import (
	"crypto/tls"
	"framework/log"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2/client"
	"go.uber.org/zap"
)

type netCallback func(data []byte) //回调消息

type ConnServer struct {
	pClient  client.PitayaClient
	disConn  chan bool   //关闭连接
	callback netCallback //回调消息
}

func NewConnectServer() *ConnServer {
	t := &ConnServer{}
	return t
}

func (t *ConnServer) SetCallback(cb netCallback) {
	t.callback = cb
}

func (t *ConnServer) ConnectServer(addr string) error {
	if t.pClient != nil && t.pClient.ConnectedStatus() {
		return nil
	}
	t.pClient = client.New(logrus.InfoLevel)
	//连接
	err := t.tryConn(addr)
	if err != nil {
		log.Error("连接失败！", zap.String("addr", addr), zap.Error(err))
		return err
	}
	log.Info("连接成功！", zap.String("addr", addr))

	go t.readMsg()
	return nil
}

func (t *ConnServer) tryConn(addr string) error {
	if err := t.pClient.ConnectToWS(addr, "", &tls.Config{
		InsecureSkipVerify: true,
	}); err != nil {
		if err := t.pClient.ConnectToWS(addr, ""); err != nil {
			if err := t.pClient.ConnectTo(addr, &tls.Config{
				InsecureSkipVerify: true,
			}); err != nil {
				if err := t.pClient.ConnectTo(addr); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// DisConnect 关闭网络
func (t *ConnServer) DisConnect() {
	if t.disConn != nil && t.pClient.ConnectedStatus() {
		t.disConn <- true
		t.pClient.Disconnect()
	}
}

// 读取网络消息
func (t *ConnServer) readMsg() {
	channel := t.pClient.MsgChannel()
	for {
		select {
		case <-t.disConn:
			close(t.disConn)
			break
		case msg := <-channel:
			t.callback(parseData(msg.Data))
		}
	}
}

// NotifyMsg 通知消息
func (t *ConnServer) NotifyMsg(route string, data []byte) {
	err := t.pClient.SendNotify(route, data)
	if err != nil {
		log.Error("发送错误SendNotify", zap.String("route", route), zap.Error(err))
	}
}
