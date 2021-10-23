package harpocrates

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type AESMode uint8

const (
	CBC AESMode = iota
	CFB
	CTR
	GCM
	OFB
)

var (
	DefaultInitialVector []byte = []byte("AAAAAAAAAAAAAA")

	ErrNotImplementedAESMode error = errors.New("harpocrates: the selected AES mode is not implemented (or yet)")
)

type AESCipher struct {
	iv   []byte // Initial Vector
	key  []byte
	mode AESMode
}

func InAESKeysize(key []byte) []byte {
	size := len(key)
	if size <= 16 {
		key, _ = ResizeByteArray(key, 16)
	} else if size <= 24 {
		key, _ = ResizeByteArray(key, 24)
	} else {
		key, _ = ResizeByteArray(key, 32)
	}
	return key
}

func NewAESCipher(iv, key []byte, mode AESMode) *AESCipher {
	if len(key) == 0 || len(iv) == 0 {
		return nil
	}

	key = InAESKeysize(key)
	iv, err := ResizeByteArray(iv, len(key))
	if err != nil {
		return nil
	}

	return &AESCipher{
		iv:   iv,
		key:  key,
		mode: mode,
	}
}

func DefaultAESCipher(key []byte) *AESCipher {
	return NewAESCipher(DefaultInitialVector, key, CBC)
}

func (ac *AESCipher) Encrypt(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, ErrZeroLength
	}

	switch ac.mode {
	case CBC:
		return AESEncryptCBC(ac.iv, ac.key, src)
	default: // Unimplemented
		return nil, ErrNotImplementedAESMode
	}
}

func (ac *AESCipher) Decrypt(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, ErrZeroLength
	}

	switch ac.mode {
	case CBC:
		return AESDecryptCBC(ac.iv, ac.key, src)
	default: // Unimplemented
		return nil, ErrNotImplementedAESMode
	}
}

func (ac *AESCipher) HexDigestEncrypt(str string) (string, error) {
	src := []byte(str)

	dst, err := ac.Encrypt(src)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}

func (ac *AESCipher) HexDigestDecrypt(hexstr string) (string, error) {
	src, err := base64.RawStdEncoding.DecodeString(hexstr)
	if err != nil {
		return "", err
	}

	dst, err := ac.Decrypt(src)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}

// Credit: https://medium.com/@thanat.arp/encrypt-decrypt-aes256-cbc-shell-script-golang-node-js-ffb675a05669

// AESEncryptCBC()
func AESEncryptCBC(iv, key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCEncrypter(block, iv)

	padded := PKCS5Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(padded))
	cbc.CryptBlocks(ciphertext, padded)

	return ciphertext, nil
}

func AESDecryptCBC(iv, key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	cbc.CryptBlocks(plaintext, ciphertext)

	return PKCS5Trimming(plaintext), nil
}

func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
