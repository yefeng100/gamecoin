package trcspendcoin

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/topfreegames/pitaya/v2/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	url             = "grpc.nile.trongrid.io:50051"
	contractAddress = "TNC8CttPhMhFg9ojddgm2Y51SCus3Ghz8P"
)

var (
	tronCli *client.GrpcClient
	tronAbi *core.SmartContract_ABI
)

func InitTronClient() {
	// 创建 TRON 客户端
	tronCli = client.NewGrpcClient(url)
	// 开始
	err := tronCli.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Errorf("tronClient.Start err. err:%v", err.Error())
		return
	}
	//defer tronClient.Stop()
}

func GetTronCli() *client.GrpcClient {
	if tronCli == nil {
		InitTronClient()
		InitAbi()
	}
	return tronCli
}

// InitAbi 初始化ABI
func InitAbi() {
	//获取abi
	a, err := GetTronCli().GetContractABI(contractAddress)
	if err != nil {
		return
	}
	tronAbi = a
}

// GetABI 获取abi
func GetABI() *core.SmartContract_ABI {
	return tronAbi
}
