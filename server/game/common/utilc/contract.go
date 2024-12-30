package utilc

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const CoinExchangeRate = 10000 //汇率，

func ScoreToCoin(val int64) *big.Int {
	secPrice := big.NewFloat(float64(val) / float64(CoinExchangeRate))
	secPrice.Mul(secPrice, big.NewFloat(1000000000000000000))
	f := new(big.Int)
	secPrice.Int(f)
	return f
}

func CoinToScore(val *big.Int) int64 {
	f := new(big.Float).SetInt(val)
	baseValue := big.NewFloat(1000000000000000000)
	amountValue := new(big.Float).Quo(f, baseValue)
	amountTem, _ := amountValue.Float64()
	amount := int64(amountTem * CoinExchangeRate)
	return amount
}

// AddrTronToEvm converts a TRON address to an EVM address, in lower case.
func AddrTronToEvm(address string) (string, error) {
	decoded := base58.Decode(address)
	if len(decoded) < 8 {
		return "", errors.New("invalid TRON address")
	}
	// Remove the first byte (0x41 prefix) and last 4 bytes (checksum)
	return "0x" + hex.EncodeToString(decoded[1:len(decoded)-4]), nil
}

// AddrEvmToTron converts an EVM address to a TRON address.
func AddrEvmToTron(address string) (string, error) {
	if len(address) < 2 {
		return "", errors.New("invalid EVM address")
	}
	if address[:2] != "0x" {
		hexAddr := common.HexToAddress(address)
		address = hexAddr.String()
	}
	addr := "41" + address[2:]
	doubleSha1, err := sha256Hex(addr)
	if err != nil {
		return "", fmt.Errorf("error in first SHA-256 hash: %w", err)
	}
	doubleSha2, err := sha256Hex(doubleSha1)
	if err != nil {
		return "", fmt.Errorf("error in second SHA-256 hash: %w", err)
	}
	checkSum := doubleSha2[:8]
	fullAddr := addr + checkSum
	decoded, err := hex.DecodeString(fullAddr)
	if err != nil {
		return "", fmt.Errorf("failed to decode address: %w", err)
	}
	return base58.Encode(decoded), nil
}

// sha256Hex computes the SHA-256 hash of a hexadecimal string and returns the result as a hexadecimal string.
func sha256Hex(msg string) (string, error) {
	bytes, err := hex.DecodeString(msg)
	if err != nil {
		return "", fmt.Errorf("invalid hex input: %w", err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:]), nil
}

// EncodeParameters 将方法签名和参数进行ABI编码
func EncodeParameters(method string, toAddress string, amount uint64) (string, error) {
	abiDefinition := fmt.Sprintf(`[{"name":"%s","type":"function","inputs":[
		{"name":"to","type":"address"},
		{"name":"value","type":"uint256"}]}]`, method)

	parsedABI, err := abi.JSON(bytes.NewReader([]byte(abiDefinition)))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %w", err)
	}

	abiMethod, ok := parsedABI.Methods[method]
	if !ok {
		return "", fmt.Errorf("method %s not found in ABI", method)
	}

	to := common.HexToAddress(toAddress)

	// ABI 编码参数
	packedData, err := abiMethod.Inputs.Pack(to, amount)
	if err != nil {
		return "", fmt.Errorf("failed to pack parameters: %w", err)
	}

	// 前加方法签名 ID
	methodID := abiMethod.ID
	fullData := append(methodID, packedData...)

	return hex.EncodeToString(fullData), nil
}
