package erc20

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/topfreegames/pitaya/v2/logger"
	"math/big"
	"project/common/cfg"
	"strings"
)

const FilePathErc20ABI = "config/contract/erc20.abi" //erc20abi文件目录

var erc20ABI string

func SubscribeContract(erc20ContAddr string) {
	erc20ABI = cfg.ReadAllTxt(FilePathErc20ABI)
	if erc20ABI == "" {
		logger.Log.Errorf("SubscribeContract abi file err, filepath=%s", FilePathErc20ABI)
		return
	}
	go start(erc20ContAddr)
}

func start(erc20ContAddr string) {
	// 连接到以太坊节点
	url := cfg.GetContIns().GetConf().GetString("contract.url")
	client, err := ethclient.Dial(url)
	if err != nil {
		logger.Log.Errorf("SubscribeContract ethclient.Dial err, filepath=%s", FilePathErc20ABI)
		return
	}
	//合约地址，如：usdt,btc,eth
	contractAddr := common.HexToAddress(erc20ContAddr)

	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		logger.Log.Errorf("SubscribeContract abi.JSON err, err=%s, erc20ABI:%s", err, erc20ABI)
		return
	}

	// 创建过滤器查询
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},                     // 监听特定合约地址
		Topics:    [][]common.Hash{{parsedABI.Events["Approval"].ID}}, // 监听 Approval 事件
	}

	// 创建事件订阅
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		//不支持通知
		logger.Log.Errorf("SubscribeContract Failed to subscribe to logs, err=%s", err)
		return
	}
	// 处理事件
	for {
		select {
		case err := <-sub.Err():
			logger.Log.Errorf("SubscribeContract Error while listening for logs, err=%s", err)

		case vLog := <-logs:
			// 解码事件
			var approvalEvent struct {
				Owner   common.Address
				Spender common.Address
				Value   *big.Int
			}

			err := parsedABI.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				logger.Log.Errorf("SubscribeContract Failed to unpack log data, err=%s", err)
				continue
			}

			// 获取 indexed 参数
			approvalEvent.Owner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			// 打印事件信息
			fmt.Printf("Approval Event:\n")
			fmt.Printf("  Owner: %s\n", approvalEvent.Owner.Hex())
			fmt.Printf("  Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("  Value: %s\n", approvalEvent.Value.String())
		}
	}
}
