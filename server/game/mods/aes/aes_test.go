package aes

import (
	"encoding/json"
	"testing"
)

func TestAes(t *testing.T) {
	type nameStr struct {
		MsgId int32
		Name  string
	}
	n := &nameStr{}
	n.MsgId = 111
	n.Name = "sdf"
	b, _ := json.Marshal(n)

	enBuf := AesEncryptECB(b, []byte("1111"))

	deBuf := AesDecryptECB(enBuf, []byte("1111"))

	n2 := &nameStr{}
	_ = json.Unmarshal(deBuf, n2)
	println(n2)

	enBuf = AesEncryptECB([]byte("123456"), []byte("Key-123^456_80307"))

	deBuf = AesDecryptECB(enBuf, []byte("1111"))
	println(deBuf)

}
