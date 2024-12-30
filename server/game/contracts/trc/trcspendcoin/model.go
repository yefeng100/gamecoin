package trcspendcoin

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/topfreegames/pitaya/v2/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"project/common/utilc"
)

var (
	url             = "grpc.nile.trongrid.io:50051"
	contractAddress = "TP28eyP4q4juXjAasWf8Y6pZkFRni7sbps"
	owner           = "TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a"
)

var (
	tronCli *client.GrpcClient
	tronAbi *core.SmartContract_ABI
)

// TriggerSmartContractRequest 定义请求结构
type TriggerSmartContractRequest struct {
	OwnerAddress     string `json:"owner_address"`     // 调用者地址
	ContractAddress  string `json:"contract_address"`  // 合约地址
	FunctionSelector string `json:"function_selector"` // 方法签名
	Parameter        string `json:"parameter"`         // ABI 编码的参数
	FeeLimit         int64  `json:"fee_limit"`         // 最大手续费（Sun）
	Visible          bool   `json:"visible"`           //
}

// TriggerSmartContractResponse 定义 TriggerSmartContract 响应结构
type TriggerSmartContractResponse struct {
	Transaction struct {
		RawDataHex string `json:"raw_data_hex"`
		TxID       string `json:"txID"`
	} `json:"transaction"`
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
}

type RawDataStruct struct {
	Contract []struct {
		Parameter struct {
			Value struct {
				Data            string `json:"data"`
				OwnerAddress    string `json:"owner_address"`
				ContractAddress string `json:"contract_address"`
			} `json:"value"`
			TypeUrl string `json:"type_url"`
		} `json:"parameter"`
		Type string `json:"type"`
	} `json:"contract"`
	RefBlockBytes string `json:"ref_block_bytes"`
	RefBlockHash  string `json:"ref_block_hash"`
	Expiration    int64  `json:"expiration"`
	FeeLimit      int    `json:"fee_limit"`
	Timestamp     int64  `json:"timestamp"`
}

// SignedTransaction 定义签名后的交易结构
type SignedTransaction struct {
	RawDataHex string   `json:"raw_data_hex"`
	Signature  []string `json:"signature"`
}

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
	if tronAbi == nil {
		InitAbi()
	}
	return tronAbi
}

// CollectTokens 提取代币到合约中
func CollectTokens(amount int64) {
	//获取grpc
	cli := GetTronCli()
	//调用方法
	methodName := "collectTokens"
	methodNameFull := methodName + "(address,address,uint256)"
	//调用参数
	num := utilc.ScoreToCoin(amount)
	numStr := num.String()
	jsonPara := fmt.Sprintf("[{\"address\": \"%s\"},{\"address\": \"%s\"},{\"uint256\": \"%s\"}]",
		"TF17BgPaZYbz8oxbjhriubPDsA7ArKoLX3", "TJ694J6ZBR22tLgPnJA9CktJjhqULPKh6C", numStr)
	// 调用智能合约(没有交易的接口)
	tx, err := cli.TriggerContract(owner, contractAddress, methodNameFull, jsonPara, 400000000, 0, "", 0)
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	// 获得keystore与account  "TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a"  "7289e085338fd7598464cb4d73688d3073b1df77356337514ed1a57446839751"
	ks, acct, _ := store.UnlockedKeystore(owner, "Chxf1986")
	// 封装Tx
	ctrlr := transaction.NewController(cli, ks, acct, tx.Transaction)
	// 真正执行Tx，并判断执行结果
	if err = ctrlr.ExecuteTransaction(); err != nil {
		fmt.Println("ExecuteTransaction err:", err.Error())
		return
	}

	// 此时Tx才上链
	fmt.Println("res :", tronCommon.BytesToHexString(tx.GetTxid()))
	fmt.Println("res :", tronCommon.BytesToHexString(tx.GetResult().GetMessage()))
}

// WithdrawTokens 将代币提到owner账号中
func WithdrawTokens() {
	//获取grpc
	cli := GetTronCli()
	//调用方法
	methodName := "withdrawTokens"
	methodNameFull := methodName + "(address)"
	//调用参数
	jsonPara := fmt.Sprintf("[{\"address\": \"%s\"}]", "TF17BgPaZYbz8oxbjhriubPDsA7ArKoLX3")
	// 调用智能合约(没有交易的接口)
	tx, err := cli.TriggerContract(owner, contractAddress, methodNameFull, jsonPara, 400000000, 0, "", 0)
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	// 获得keystore与account  "TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a"  "7289e085338fd7598464cb4d73688d3073b1df77356337514ed1a57446839751"
	ks, acct, _ := store.UnlockedKeystore(owner, "Chxf1986")
	// 封装Tx
	ctrlr := transaction.NewController(cli, ks, acct, tx.Transaction)
	// 真正执行Tx，并判断执行结果
	if err = ctrlr.ExecuteTransaction(); err != nil {
		fmt.Println("ExecuteTransaction err:", err.Error())
		return
	}

	// 此时Tx才上链
	fmt.Println("res :", tronCommon.BytesToHexString(tx.GetTxid()))
	fmt.Println("res :", tronCommon.BytesToHexString(tx.GetResult().GetMessage()))
}

// SignTransaction 使用私钥签名交易
func SignTransaction(rawDataHex string, privateKey string) (SignedTransaction, error) {
	// 将私钥转换为 ECDSA 格式
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return SignedTransaction{}, fmt.Errorf("invalid private key: %w", err)
	}

	// 将 rawDataHex 解码为字节数组
	rawData, err := hex.DecodeString(rawDataHex)
	if err != nil {
		return SignedTransaction{}, fmt.Errorf("failed to decode raw_data_hex: %w", err)
	}

	// 对交易进行哈希
	txHash := crypto.Keccak256(rawData)

	// 使用私钥签名哈希
	signature, err := crypto.Sign(txHash, privKey)
	if err != nil {
		return SignedTransaction{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// 构造签名后的交易
	return SignedTransaction{
		RawDataHex: rawDataHex,
		Signature:  []string{hex.EncodeToString(signature)},
	}, nil
}

// BroadcastTransaction 调用广播接口
func BroadcastTransaction(apiURL string, signedTx SignedTransaction) error {
	// JSON 编码签名后的交易
	body, err := json.Marshal(signedTx)
	if err != nil {
		return fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	// 发送 HTTP POST 请求
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send broadcast request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// 打印响应
	fmt.Printf("Broadcast Response: %s\n", string(respBody))
	return nil
}
