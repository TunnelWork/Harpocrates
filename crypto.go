package harpocrates

import "crypto/ed25519"

func Ed25519Key(seed string) ed25519.PrivateKey {
	if len(seed) == 0 {
		return nil
	}

	var byteSeed []byte = []byte(seed)

	for len(byteSeed) < 32 {
		byteSeed = append(byteSeed, byteSeed...)
	}

	if len(byteSeed) > 32 {
		byteSeed = byteSeed[:32]
	}

	return ed25519.NewKeyFromSeed(byteSeed)
}
