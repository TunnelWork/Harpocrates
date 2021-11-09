package harpocrates

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
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

func GetRandomNumber(min, max int) int {
	c, err := crand.Int(crand.Reader, big.NewInt(int64(max-min)))
	if err == nil {
		return int(c.Int64()) + min
	}
	return min + rand.Intn(max-min) // skipcq: GSC-G404
}

// GetRandomHex generates a Hex (as string)
// with maximum value bound by bytes as max < 2^(bytes*8)
func GetRandomHex(bytes uint) (string, error) {
	return GetRandomString(bytes*2, runesHEX)
}

// GetRandomBase32 generates a Base32 (as string)
func GetRandomBase32() (string, error) {
	return GetRandomString(32, runesBASE32)
}
