package harpocrates

// GetNewWeakPassword() is deprecated and should not be used
// without a good reason.
func GetNewWeakPassword(n uint) (string, error) {
	return GetRandomString(n, runesAlphaNumeric)
}

func GetNewStrongPassword(n uint) (string, error) {
	return GetRandomString(n, runesComplete)
}
