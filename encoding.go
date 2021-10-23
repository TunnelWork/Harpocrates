package harpocrates

import (
	b64 "encoding/base64"
	"errors"
)

var (
	ErrDecodeBase64 = errors.New("harpocrates: error decoding base64")
)

func Base64Encoding(plaintext string) string {
	return b64.StdEncoding.EncodeToString([]byte(plaintext))
}

func Base64Decoding(encoded string) string {
	decoded, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "ERR_DECODE_BASE64"
	}
	return string(decoded)
}

func IsBase64(sus string) bool {
	_, err := b64.StdEncoding.DecodeString(sus)
	return err == nil
}
