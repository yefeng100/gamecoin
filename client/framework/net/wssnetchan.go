package net

// WssClientEventChan 客户端消息通道
type WssClientEventChan struct {
	MsgSig          chan []byte   //完整数据包消息通道
	ConnectedSig    chan *WssConn //连接成功消息通道
	DisconnectedSig chan *WssConn //连接断开消息通道
}

// NewWssClientEventChan 新建客户端消息通道
func NewWssClientEventChan(sigSize int) *WssClientEventChan {
	evt := new(WssClientEventChan)

	evt.MsgSig = make(chan []byte, sigSize)
	evt.ConnectedSig = make(chan *WssConn, sigSize)
	evt.DisconnectedSig = make(chan *WssConn, sigSize)

	return evt
}
