package net

import (
	"errors"
	"fmt"
	"framework/log"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net"
	"time"
)

// WssConn  wss连接
type WssConn struct {
	conn     *websocket.Conn
	connAddr string
	nowTime  int64
}

func NewWsConn(conn *websocket.Conn, addr string) *WssConn {
	t := new(WssConn)
	t.conn = conn
	t.connAddr = addr
	t.nowTime = time.Now().Unix()
	return t
}

func (t *WssConn) GetConn() *websocket.Conn {
	return t.conn
}

func (t *WssConn) ReadMsg() (returnData []byte, returnErr error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("wsConn read msg error", zap.Any("err", err))

			returnData, returnErr = nil, errors.New(fmt.Sprintf("%v", err))
			return
		}
	}()

	_, message, err := t.conn.ReadMessage()
	if err != nil {
		log.Error("wss conn read msg error", zap.Any("err", err))
		returnData, returnErr = nil, err
		return
	}
	returnData, returnErr = message, nil
	return
}

func (t *WssConn) WriteMsg(messageType int, data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	return t.conn.WriteMessage(messageType, data)
}

func (t *WssConn) LocalAddr() net.Addr {
	return t.conn.LocalAddr()
}

func (t *WssConn) RemoteAddr() net.Addr {
	return t.conn.RemoteAddr()
}

func (t *WssConn) Close() {
	_ = t.conn.Close()
}
