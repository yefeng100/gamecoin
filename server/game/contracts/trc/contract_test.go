package trc

import (
	"encoding/hex"
	"fmt"
	eABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	"math/big"
	"project/common/utilc"
	"project/contracts/trc/trcspendcoin"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"strings"
	"testing"
)

var (
	url             = "grpc.nile.trongrid.io:50051"
	contractAddress = "TNC8CttPhMhFg9ojddgm2Y51SCus3Ghz8P"
	owner           = "TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a"
)

func Test(t *testing.T) {
	//a1 := common.HexToAddress("000000000000000000000000669a25f58e2d8189b9d033c5eb041748c96ffbec")
	//a2, _ := utilc.AddrEvmToTron("000000000000000000000000669a25f58e2d8189b9d033c5eb041748c96ffbec")
	//as, _ := utilc.AddrEvmToTron(a1.String())
	//fmt.Println("EVM to Tron", as, a2)
	//TriggerConstantContract()
	setNumber()
	getNumber()
	if false {
		GetOwner()
	}
}

func getNumber() {
	//获取grpc
	cli := trcspendcoin.GetTronCli()
	//调用方法
	methodName := "getNumer"
	methodNameFull := methodName + "()"
	//调用参数
	// 调用智能合约
	tx, err := cli.TriggerConstantContract(owner, contractAddress, methodNameFull, "")
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	//解码方法
	arg, err := abi.GetParser(trcspendcoin.GetABI(), methodName)
	if err != nil {
		fmt.Println("abi.GetParser err.", err.Error())
		return
	}
	//返回参数
	var result []interface{}
	result, err = arg.Unpack(tx.ConstantResult[0])
	fmt.Println(result[0].(*big.Int).Int64())
}

func setNumber() {
	//获取grpc
	cli := trcspendcoin.GetTronCli()
	//调用方法
	methodName := "setNumber"
	methodNameFull := methodName + "(uint256)"
	//调用参数
	num := big.NewInt(10001)
	numStr := num.String()
	jsonPara := fmt.Sprintf("[{\"uint256\": %s}]", numStr)
	// 调用智能合约
	res, err := cli.TriggerConstantContract(owner, contractAddress, methodNameFull, jsonPara)
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	fmt.Println("ok", res)
}

func GetOwner() {
	//获取grpc
	cli := trcspendcoin.GetTronCli() //调用方法
	methodName := "owner"
	methodNameFull := methodName + "()"
	//调用参数
	// 调用智能合约
	tx, err := cli.TriggerConstantContract(owner, contractAddress, methodNameFull, "")
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	//解码方法
	arg, err := abi.GetParser(trcspendcoin.GetABI(), methodName)
	if err != nil {
		fmt.Println("abi.GetParser err.", err.Error())
		return
	}
	//返回参数
	var result []interface{}
	result, err = arg.Unpack(tx.ConstantResult[0])
	addr := result[0].(common.Address)
	oAddr, _ := utilc.AddrEvmToTron(addr.String())
	fmt.Println("owner addr:", oAddr)
}

func GetOwner5() {
	//privateKey := "7289e085338fd7598464cb4d73688d3073b1df77356337514ed1a57446839751"
	// 创建 TRON 客户端
	tronClient := client.NewGrpcClient(url)

	err := tronClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("err.", err.Error())
		return
	}
	defer tronClient.Stop()

	// ABI 编码 collectTokens 方法及其参数
	methodSignature := `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"address","name":"from","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"collectTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"}],"name":"withdrawTokens","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	parsedABI, err := eABI.JSON(strings.NewReader(methodSignature))
	if err != nil {
		fmt.Println("解析失败，err.", err.Error())
		return
	}
	method, _ := parsedABI.Methods["owner"]
	// 构建交易
	//data, _ := method.Inputs.Pack()
	// 调用智能合约
	result, err := tronClient.TriggerConstantContract("TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a", contractAddress, "owner()", "")
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	// 解码返回值
	output, err := method.Outputs.Unpack(result.ConstantResult[0])
	data := tronCommon.BytesToHexString(result.ConstantResult[0])
	fmt.Println("aaaaaaa1", data)
	data2 := tronCommon.ToHex(result.ConstantResult[0])
	fmt.Println("aaaaaaa2", data2)
	data3 := hex.EncodeToString(result.ConstantResult[0])
	fmt.Println("aaaaaaa3", data3)

	switch output[0].(type) {
	case string:
		fmt.Println("string", output[0].(string))
		break
	case int:
		fmt.Println("int", output[0].(int))
		break
	case float64:
		fmt.Println("float64", output[0].(float64))
		break
	case common.Address:
		addr := output[0].(common.Address)
		fmt.Println("获取的数据", addr.String())
		as, _ := utilc.AddrEvmToTron(addr.String())
		fmt.Println("EVM to Tron", as)
		as2, _ := utilc.AddrTronToEvm(as)
		fmt.Println("Tron to EVM", as2)
		break
	}
	fmt.Println("调用结果:", err, output)
}

func getNumber2() {
	//abiStr := cfg.ReadAllTxt("config/contract/trcspendcoin.abi")
	abiStr := "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"collectTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getNumerMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	if abiStr == "" {
		fmt.Println("abiStr is nil:")
		return
	}
	// ABI 编码 collectTokens 方法及其参数
	parsedABI, err := eABI.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Println("解析失败，err.", err.Error())
		return
	}
	//获取grpc
	cli := trcspendcoin.GetTronCli()
	//调用方法
	methodName := "getNumer"
	methodNameFull := methodName + "()"
	//调用参数
	// 调用智能合约
	result, err := cli.TriggerConstantContract(owner, contractAddress, methodNameFull, "")
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}
	//解码方法
	method, _ := parsedABI.Methods[methodName]
	// 解码返回值
	output, err := method.Outputs.Unpack(result.ConstantResult[0])

	switch output[0].(type) {
	case *big.Int:
		b := output[0].(*big.Int)
		num := b.Int64()
		fmt.Println(num)
	}
	fmt.Println(output)
}

func GetOwner1() {

	// 创建 Tron 实例，连接到 Tron 网络
	/*client := tronweb.New("https://api.tronstack.io") // 或使用主网或测试网

	// 使用合约地址和方法名称调用智能合约
	contractAddress := "你的合约地址"
	methodName := "你的智能合约方法"

	// 如果需要，可以添加输入参数
	params := []interface{}{"参数1", "参数2"}

	// 调用合约
	result, err := client.CallContract(contractAddress, methodName, params)*/
}

func GetOwner2() {
	/*url := "https://api.shasta.trongrid.io"
	contractAddress := "TC2igFMeiP5s7xnHppsj92zaboUe8v9gMu"
	//privateKey := "7289e085338fd7598464cb4d73688d3073b1df77356337514ed1a57446839751"
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
	parsedABI, err := eABI.JSON(strings.NewReader(methodSignature))
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
	tx, err := tronClient.TriggerConstantContract("TSuEeaVL5MpPmeZpW2rfd4SrnQobfh73Hv", contractAddress, "collectTokens", methodSignature)
	if err != nil {
		fmt.Println("构造交易失败:", err.Error())
		return
	}

	// 广播交易
	txID, err := tronClient.BroadcastTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println("广播交易失败", err.Error())
	}

	fmt.Println("交易成功，交易哈希", txID)*/
}

func TestGetAccount(t *testing.T) {
	url := "https://api.shasta.trongrid.io/wallet/getaccount"

	payload := strings.NewReader("{\"address\":\"TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestGetAccountBalance(t *testing.T) {

	url := "https://api.shasta.trongrid.io/wallet/getaccountbalance"

	payload := strings.NewReader("{\"account_identifier\":{\"address\":\"TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a\"}, }}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func GetContract() {
	url := "https://api.nileex.io/wallet/getcontract"

	payload := strings.NewReader("{\"value\":\"TC2igFMeiP5s7xnHppsj92zaboUe8v9gMu\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
func GetContractInfo() {
	url := "https://api.nileex.io/wallet/getcontractinfo"

	payload := strings.NewReader("{\"value\":\"TC2igFMeiP5s7xnHppsj92zaboUe8v9gMu\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TriggerConstantContract() {
	url := "https://api.nileex.io/wallet/triggerconstantcontract"

	para := "{\"owner_address\":\"TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a\",\"contract_address\":\"TNC8CttPhMhFg9ojddgm2Y51SCus3Ghz8P\",\"function_selector\":\"owner()\",\"parameter\":\"\",\"visible\":true}"

	payload := strings.NewReader(para)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

/*
var c *Contract = nil
var mu sync.RWMutex

// Contract 合约连接结构
type Contract struct {
	Client    *ethclient.Client
	AN        *Trcspendcoin
	NetworkID *big.Int
}

// GetIns 获取合约信息
func GetIns() *Contract {
	mu.Lock()
	defer mu.Unlock()
	if c == nil {
		//url := cfg.GetContIns().GetConf().GetString("trc.url")
		//addr := cfg.GetContIns().GetConf().GetString("trc.spendcoinaddr")
		url := "https://grpc.nile.trongrid.io:50051"
		addr := "TC2igFMeiP5s7xnHppsj92zaboUe8v9gMu"
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
	instance, err := NewTrcspendcoin(address, client)
//netId, errID := client.NetworkID(context.Background())
	//if errID != nil {
	//	logger.Log.Error("InitContract NetworkID err", zap.String("url", url), zap.String("addr", addr), zap.Error(err))
	//	return nil, err
	//}

	c = &Contract{
		Client: client,
		AN:     instance,
		//NetworkID: netId,
	}
	return c, nil
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
*/
