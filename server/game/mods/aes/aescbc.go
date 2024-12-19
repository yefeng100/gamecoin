package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// AesEncryptCBC 加密
func AesEncryptCBC(plaintext []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err=", err)
		return "", errors.New("invalid decrypt key")
	}
	blockSize := block.BlockSize()
	plaintext = pKCS5Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	encryptString := base64.RawURLEncoding.EncodeToString([]byte(ciphertext))
	return encryptString, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	b := []byte{byte(padding)}
	padtext := bytes.Repeat(b, padding)
	return append(ciphertext, padtext...)
}

// AesDecryptCBC 解密
func AesDecryptCBC(ciphertext string, key []byte, iv []byte) (string, error) {
	decode_data, err := base64.RawURLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", errors.New("invalid decrypt data")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err=", err)
		return "", errors.New("invalid decrypt key")
	}
	blockModel := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(decode_data))
	blockModel.CryptBlocks(plaintext, decode_data)
	plaintext = pKCS5UnPadding(plaintext)
	return string(plaintext), nil
}

func pKCS5UnPadding(ciphertext []byte) []byte {
	length := len(ciphertext)
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
