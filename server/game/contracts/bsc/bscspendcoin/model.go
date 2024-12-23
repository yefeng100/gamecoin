package bscspendcoin

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
	"math/big"
	"project/common/cfg"
	"project/common/utilc"
	"project/contracts/bsc/keystore"
	"sync"
)

var c *Contract = nil
var mu sync.RWMutex

// Contract 合约连接结构
type Contract struct {
	Client    *ethclient.Client
	AN        *Bscspendcoin
	NetworkID *big.Int
}

// GetIns 获取合约信息
func GetIns() *Contract {
	mu.Lock()
	defer mu.Unlock()
	if c == nil {
		url := cfg.GetContIns().GetConf().GetString("contract.url")
		addr := cfg.GetContIns().GetConf().GetString("bsc.spendcoinaddr")
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
	instance, err := NewBscspendcoin(address, client)
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

// SpendUserCoin 授权转账, （usdt地址，玩家钱包地址，目标钱包地址，额度）
func SpendUserCoin(coinAddr, fromAddr, toAddr string, amount int64) (*types.Transaction, error) {
	ins := GetIns()
	if ins == nil {
		logger.Log.Error("SpendUserCoin ins is nil err")
		return nil, errors.New("SpendUserCoin ins is nil err")
	}
	privateKey, err := keystore.GenOptsByChainIdPrv(ins.NetworkID)
	if err != nil {
		logger.Log.Error("SpendUserCoin privateKey err", zap.Error(err))
		return nil, err
	}
	contractAddress := common.HexToAddress(coinAddr)
	f := common.HexToAddress(fromAddr)
	t := common.HexToAddress(toAddr)
	a := utilc.ScoreToCoin(amount)
	tx, err := ins.AN.SpendUserCoin(privateKey, contractAddress, f, t, a)
	if err != nil {
		logger.Log.Error("SpendUserCoin err", zap.Error(err))
		return nil, err
	}
	logger.Log.Infof("SpendUserCoin hash:%s", tx.Hash().String())
	return tx, nil
}

func Owner() string {
	ins := GetIns()
	if ins == nil {
		logger.Log.Error("SpendUserCoin ins is nil err")
		return ""
	}
	addr, err := ins.AN.Owner(nil)
	if err != nil {
		logger.Log.Error("SpendUserCoin err", zap.Error(err))
		return ""
	}
	strAddr := addr.String()
	return strAddr
}
