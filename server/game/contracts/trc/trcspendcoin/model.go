package trcspendcoin

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"project/common/utilc"
	"strings"
)

func GetOwner() {
	url := "https://api.shasta.trongrid.io"
	contractAddress := "TC2igFMeiP5s7xnHppsj92zaboUe8v9gMu"
	privateKey := "7289e085338fd7598464cb4d73688d3073b1df77356337514ed1a57446839751"
	// 创建 TRON 客户端
	tronClient := client.NewGrpcClient(url)
	err := tronClient.Start()
	if err != nil {
		fmt.Println("err.", err.Error())
		return
	}
	defer tronClient.Stop()

	// ABI 编码 collectTokens 方法及其参数
	methodSignature := `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"address","name":"from","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"collectTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"}],"name":"withdrawTokens","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(methodSignature))
	if err != nil {
		fmt.Println("解析失败，err.", err.Error())
		return
	}
	amount := utilc.ScoreToCoin(10000)
	// 编码方法和参数
	data, err := parsedABI.Pack("collectTokens", common.HexToAddress("TF17BgPaZYbz8oxbjhriubPDsA7ArKoLX3"), common.HexToAddress("TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a"), amount)
	if err != nil {
		fmt.Println("ABI 编码失败", err.Error())
		return
	}

	// 转换为十六进制字符串
	dataHex := hex.EncodeToString(data)
	// 构造交易
	tx, err := tronClient.TriggerContract(contractAddress, dataHex, methodSignature, privateKey)
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}

	// 广播交易
	txID, err := tronClient.BroadcastTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println("广播交易失败", err.Error())
	}

	fmt.Println("交易成功，交易哈希", txID)
}
