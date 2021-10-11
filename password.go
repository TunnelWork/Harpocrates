package harpocrates

// GetNewWeakPassword() is deprecated and should not be used
// without a good reason.
func GetNewWeakPassword(n uint) (string, error) {
	return GetRandomString(n, runesAlphaNumeric)
}

func GetNewStrongPassword(n uint) (string, error) {
	return GetRandomString(n, runesComplete)
}

// GetRandomHex generates a Hex (as string)
// with maximum value bound by bytes as max < 2^(bytes*8)
func GetRandomHex(bytes uint) (string, error) {
	return GetRandomString(bytes*2, runesHEX)
}
