package aes

import (
	"crypto/aes"
	"encoding/base64"
)

// AesEncryptECB Aes ECB加密
func AesEncryptECB(origData []byte, key []byte) string {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	encryptString := base64.StdEncoding.EncodeToString([]byte(encrypted))
	return encryptString
}

// AesDecryptECB Aes ECB解密
func AesDecryptECB(encryptedStr string, key []byte) (decrypted []byte) {
	encrypted, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return []byte("")
	}
	cipher, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return []byte("")
	}
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	if trim < 0 {
		return []byte("")
	}
	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
