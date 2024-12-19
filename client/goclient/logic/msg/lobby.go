package msg

import (
	"framework/log"
	"go.uber.org/zap"
)

func HandlerData(data []byte) {
	log.Info("收到消息", zap.String("data", string(data)))
}
