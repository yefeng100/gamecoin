package net

import (
	"errors"
	"framework/log"
	"framework/util"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net"
	"net/url"
	"sync"
	"time"
)

type WssClient struct {
	mu          *sync.RWMutex       //锁
	evtChan     *WssClientEventChan //消息通道
	isConnected bool                //是否已经连接
	conn        *WssConn            //客户端连接
	connAddr    string              //连接地址
	isReConn    bool                //是否重连
}

// NewWssClient 创建客户端wss网络
func NewWssClient(evtChan *WssClientEventChan, addr string, isReConn bool) *WssClient {
	clt := &WssClient{}
	clt.mu = &sync.RWMutex{}
	clt.evtChan = evtChan
	clt.isConnected = false
	clt.connAddr = addr
	clt.isReConn = isReConn
	return clt
}

func (t *WssClient) Connect() error {
	if t.isConnected {
		log.Error("wss已经建立连接")
		return errors.New("wss已经建立连接")
	}
	u, err := url.Parse(t.connAddr)
	if err != nil {
		log.Error("WssClient Parse addr error", zap.Error(err))
		return err
	}
	log.Info("wss连接", zap.String("addr", u.String()))
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error("WssClient Dial error", zap.Error(err))
		return err
	}
	t.conn = NewWsConn(conn, u.String())
	t.isConnected = true
	t.evtChan.ConnectedSig <- t.conn
	//重连检查
	t.reConn()
	//读数据
	go t.readMsg()

	return nil
}

// 读数据
func (t *WssClient) readMsg() {
	defer util.PanicErrStack()
	log.Info("wss client 开始读数据", zap.String("addr", t.conn.connAddr))
	for {
		msg, err := t.conn.ReadMsg()
		if err != nil {
			log.Error("wss client read msg error", zap.Error(err))
			t.conn.Close()
			break
		}
		if msg == nil || len(msg) == 0 {
			log.Error("wss client read msg nil", zap.String("addr", t.conn.connAddr))
			t.conn.Close()
			break
		}
		t.evtChan.MsgSig <- msg
	}
}

// WriteMsg 写消息
func (t *WssClient) WriteMsg(byData []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.isConnected {
		return errors.New("wss not connect")
	}

	if byData == nil || len(byData) == 0 {
		log.Error("wss client write msg nil", zap.String("addr", t.conn.connAddr))
		return nil
	}
	err := t.conn.WriteMsg(websocket.TextMessage, byData)
	if err != nil {
		log.Error("wss client write msg error", zap.Error(err))
		go t.Close()
	}
	return nil
}

func (t *WssClient) Close() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.isConnected {
		log.Error("wss 已经被关闭")
		return
	}
	t.isConnected = false
	t.conn.Close()
	t.evtChan.DisconnectedSig <- t.conn
}

// IsConnected 是否连接
func (t *WssClient) IsConnected() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.isConnected
}

func (t *WssClient) LocalAddr() net.Addr {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.conn.GetConn().LocalAddr()
}

func (t *WssClient) RemoteAddr() net.Addr {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.conn.GetConn().RemoteAddr()
}

// 断开重连
func (t *WssClient) reConn() {
	if !t.isReConn {
		return
	}
	go func() {
		defer util.PanicErrStack()
		for {
			time.Sleep(time.Second * 2)
			if !t.IsConnected() {
				log.Error("开始重连", zap.String("addr", t.connAddr))
				_ = t.Connect()
			}
		}
	}()
}

// GetAddr 当前地址
func (t *WssClient) GetAddr() string {
	return t.connAddr
}

func (t *WssClient) GetConn() *WssConn {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.conn
}
