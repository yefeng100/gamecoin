package msgsrv

import (
	"bytes"
	"encoding/binary"
)

// Message 网络消息对象定义
type Message struct {
	BClassID int32       //一级协议号
	SClassID int32       //二级协议号
	MsgData  interface{} //消息数据
}

func Marshal(msg *Message) ([]byte, error) {
	b1 := IntToBytes(msg.BClassID, true)
	b2 := IntToBytes(msg.SClassID, true)
	bh := append(b1, b2...)

	var msgData []byte
	if by, ok := msg.MsgData.([]byte); ok {
		msgData = by
	}
	ret := append(bh, msgData...)
	return ret, nil
}

func UnMarshal(data []byte) (*Message, error) {
	m1 := new(Message)
	m1.BClassID = BytesToInt(data[:4], true)
	m1.SClassID = BytesToInt(data[4:8], true)
	m1.MsgData = data[8:]
	return m1, nil
}

// IntToBytes int转换为bytes
func IntToBytes(n int32, bigEndian bool) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	order = binary.BigEndian
	if !bigEndian {
		order = binary.LittleEndian
	}

	binary.Write(bytesBuffer, order, tmp)
	return bytesBuffer.Bytes()
}

// BytesToInt bytes转换为int
func BytesToInt(data []byte, bigEndian bool) int32 {
	bytesBuffer := bytes.NewBuffer(data)

	var order binary.ByteOrder
	order = binary.BigEndian
	if !bigEndian {
		order = binary.LittleEndian
	}

	var x int32
	binary.Read(bytesBuffer, order, &x)

	return x
}

// StructToBytes struct转换为bytes
func StructToBytes(st interface{}, bigEndian bool) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	order = binary.BigEndian
	if !bigEndian {
		order = binary.LittleEndian
	}

	binary.Write(bytesBuffer, order, st)
	return bytesBuffer.Bytes()
}
