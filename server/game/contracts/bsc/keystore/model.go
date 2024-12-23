package keystore

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"os"
)

var (
	Key        *KeyStore
	KeyPrivate *KeyStore
)

type KeyStore struct {
	Ks      *keystore.KeyStore
	Account accounts.Account
}

func InitKeyStore() *KeyStore {
	password := viper.GetString("Bsc.Pwd")

	dirName := "./tmp"
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		logger.Log.Error("ReadDir error", zap.Error(err))
	}
	var account accounts.Account
	//var err error
	ks := keystore.NewKeyStore(dirName, keystore.StandardScryptN, keystore.StandardScryptP)
	password += "AiseP"
	if len(dir) == 0 {
		account, err = ks.NewAccount(password)
		if err != nil {
			logger.Log.Error("NewAccount error", zap.Error(err))
		}
		err = ks.Unlock(account, password)
		if err != nil {
			logger.Log.Error("Unlock error", zap.Error(err))
		}
		Key = &KeyStore{
			Ks:      ks,
			Account: account,
		}
		return Key
	}
	acc := ks.Accounts()
	err = ks.Unlock(acc[0], password)
	if err != nil {
		logger.Log.Error("Unlock error", zap.Error(err))
	}
	Key = &KeyStore{
		Ks:      ks,
		Account: acc[0],
	}
	return Key
}

// KeyStoreConvertPrivate 密钥库转私钥
func KeyStoreConvertPrivate() {
	pwd := "0712!@#$qwerAiseP"
	keyjson, err := os.ReadFile("tmp/UTC--2024-07-15T09-54-03.191043900Z--620d2c6ddea04a7c98a25f17d36080ba902ba4d1")
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
		return
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, pwd)
	if err != nil {
		return
	}
	pKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	fmt.Println("%v, %v", pKey, addr)
	//InitKeyStoreByPrivate(pwd, addr.String(), pKey)
}

func GenOptsByChainId(chainId *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyStoreTransactorWithChainID(Key.Ks, Key.Account, chainId)
	return opts, err
}
