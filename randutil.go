package harpocrates

import (
	crand "crypto/rand"
	"math/big"
)

// Credits:
// - https://github.com/pion/randutil

func GetRandomString(n uint, runes []rune) (string, error) {
	b := make([]rune, n)
	for i := range b {
		v, err := crand.Int(crand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			return "", err
		}
		b[i] = runes[v.Int64()]
	}
	return string(b), nil
}
