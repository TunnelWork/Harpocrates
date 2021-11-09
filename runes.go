package harpocrates

var (
	// Pure-blood runes
	runesLowerCase []rune = []rune("abcdefghijklmnopqrstuvwxyz")
	runesUpperCase []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	runesNumber    []rune = []rune("0123456789")
	runesHEX       []rune = []rune("0123456789ABCEDF")
	runesBASE32    []rune = []rune("ABCDEFGHIJKLMNOPQRSTUV1234567890")
	runesSymbol    []rune = []rune(`~!@#$%^&*_+-=[]{};:,.<>?`) // non-escaping and non-ambiguous symbols

	// Combined runes
	runesMixedCase    []rune = append(runesLowerCase, runesUpperCase...) // Both lowercase & UPPERCASE
	runesAlphaNumeric []rune = append(runesMixedCase, runesNumber...)    // AllCases and numbers
	runesComplete     []rune = append(runesAlphaNumeric, runesSymbol...)
)
