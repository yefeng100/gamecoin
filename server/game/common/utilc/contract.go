package utilc

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
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
