package rediscache

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2/logger"
	"project/base/cache"
	"project/modules/redisstorage"
)

func RGetNonce(ctx context.Context, userKey string) string {
	key := fmt.Sprintf(cache.RKeyUserNonce, userKey)
	conn := redisstorage.GetConn()
	if conn == nil {
		logger.Log.Errorf("RGetNonce. redis connection error")
		return ""
	}
	randNum, err := conn.Get(ctx, key).Result()
	if err != nil {
		logger.Log.Errorf("RGetNonce. redis get err. key:%s, err:%s", userKey, err)
		return ""
	}
	return randNum
}

func RSetNonce(ctx context.Context, userKey, nonce string) {
	key := fmt.Sprintf(cache.RKeyUserNonce, userKey)
	conn := redisstorage.GetConn()
	if conn == nil {
		logger.Log.Errorf("RGetNonce. redis connection error")
		return
	}
	conn.Set(ctx, key, nonce, cache.RRExpirationSecond30)
}
