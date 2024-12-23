package erc20

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
	"math/big"
	"project/common/cfg"
	"sync"
)

var c *Contract = nil
var mu sync.RWMutex

var usdtAddr = "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"

// Contract 合约连接结构
type Contract struct {
	Client    *ethclient.Client
	AN        *Erc20
	NetworkID *big.Int
}

// GetIns 获取合约信息
func GetIns() *Contract {
	mu.Lock()
	defer mu.Unlock()
	if c == nil {
		url := cfg.GetContIns().GetConf().GetString("contract.url")
		addr := usdtAddr
		if addr == "" || url == "" {
			logger.Log.Error("InitContract1 err", zap.String("url", url), zap.String("addr", addr))
			return nil
		}
		t, err := InitContract(url, addr)
		if err != nil {
			logger.Log.Error("InitContract1 err", zap.String("url", url), zap.String("addr", addr), zap.Error(err))
			return nil
		}
		c = t

		//创建定时器
		NewTimerErc20()
	}

	return c
}

func InitContract(url, addr string) (*Contract, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		logger.Log.Error("InitContract err", zap.String("url", url), zap.String("addr", addr), zap.Error(err))
		return nil, err
	}
	address := common.HexToAddress(addr)
	instance, err := NewErc20(address, client)
	netId, errID := client.NetworkID(context.Background())
	if errID != nil {
		logger.Log.Error("InitContract NetworkID err", zap.String("url", url), zap.String("addr", addr), zap.Error(err))
		return nil, err
	}

	c = &Contract{
		Client:    client,
		AN:        instance,
		NetworkID: netId,
	}
	return c, nil
}
