package keystore

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/topfreegames/pitaya/v2/logger"
	"math/big"
	"project/common/cfg"
)

var privateKey *ecdsa.PrivateKey = nil

// initPrivateKey 初始化私钥
func initPrivateKey() {
	priKey := cfg.GetContIns().GetConf().GetString("contract.privatekey")
	pkHex, err := crypto.HexToECDSA(priKey)
	if err != nil {
		logger.Log.Errorf("InitPrivateKey err. %v", err)
		return
	}
	privateKey = pkHex
}

// GenOptsByChainIdPrv 获取私钥
func GenOptsByChainIdPrv(chainId *big.Int) (*bind.TransactOpts, error) {
	if privateKey == nil {
		initPrivateKey()
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	return opts, err
}
