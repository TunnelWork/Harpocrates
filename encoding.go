package harpocrates

import (
	b64 "encoding/base64"
)

func Base64Encoding(plaintext string) string {
	return b64.StdEncoding.EncodeToString([]byte(plaintext))
}

func Base64Decoding(encoded string) string {
	decoded, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(decoded)
}
