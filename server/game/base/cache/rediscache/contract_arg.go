package rediscache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/topfreegames/pitaya/v2/logger"
	"project/modules/redisstorage"
	"strconv"
)

const (
	RKeyErcBlockApprove = "contract:erc:block:approve" //erc20授权快
)

// RGetErcBlockApprove erc20授权快
func RGetErcBlockApprove() int64 {
	key := RKeyErcBlockApprove
	conn := redisstorage.GetConn()
	if conn == nil {
		logger.Log.Errorf("RGetErcBlockApprove. redis connection error")
		return 0
	}
	block, err := conn.Get(context.TODO(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0
		}
		logger.Log.Errorf("RGetErcBlockApprove. redis get err. key:%s, err:%s", key, err)
		return -1
	}
	b, err := strconv.ParseInt(block, 10, 64)
	if err != nil {
		logger.Log.Errorf("RGetErcBlockApprove. atoi err. block:%s, err:%s", block, err)
		return 0
	}
	return b
}

// RSetErcBlockApprove erc20授权快
func RSetErcBlockApprove(block int64) error {
	key := RKeyErcBlockApprove
	conn := redisstorage.GetConn()
	if conn == nil {
		logger.Log.Errorf("RSetErcBlockApprove. redis connection error")
		return errors.New("RSetErcBlockApprove redis connection error")
	}
	err := conn.Set(context.TODO(), key, block, 0).Err()
	if err != nil {
		logger.Log.Errorf("RSetErcBlockApprove. redis get err. key:%s, err:%s", key, err)
		return err
	}
	return nil
}
