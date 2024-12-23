package cfg

import (
	"bytes"
	"encoding/json"
	"github.com/topfreegames/pitaya/v2/logger"
	"os"
)

// ReadAllTxt 读取全文本
func ReadAllTxt(name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		logger.Log.Errorf("open file err:%s", err)
		return ""
	}
	dst := &bytes.Buffer{}
	_ = json.Compact(dst, f)
	return dst.String()
}
