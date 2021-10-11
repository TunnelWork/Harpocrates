package harpocrates

import (
	"bytes"
	"crypto/aes"
	"sync"
)

var (
	cryptoParamsLock = sync.RWMutex{}
	AESCipherIV      = []byte("0123456789ABCDEF")
)

func NewAESCipherIV(iv []byte) {
	cryptoParamsLock.Lock()
	defer cryptoParamsLock.Unlock()
	AESCipherIV = iv
}

func trimAESKey(input []byte) []byte {
	var output []byte

	if len(input) > 32 {
		output = input[:32]
	} else {
		keyPadding := make([]byte, 32-len(input))
		output = append(input, keyPadding...)
	}

	return output
}

// Credits: https://gist.github.com/yingray/57fdc3264b1927ef0f984b533d63abab
func Encrypt(src string, aeskey []byte) ([]byte, error) {
	cryptoParamsLock.RLock()
	defer cryptoParamsLock.RUnlock()
	// Truncate or Concatenate, until the key is 32-byte long
	key := trimAESKey(aeskey)

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
