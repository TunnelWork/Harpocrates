package harpocrates

import "crypto/ed25519"

func Ed25519Key(seed []byte) ed25519.PrivateKey {
	if len(seed) == 0 {
		return nil
	}

	for len(seed) < 32 {
		seed = append(seed, seed...)
	}

	if len(seed) > 32 {
		seed = seed[:32]
	}

	return ed25519.NewKeyFromSeed(seed)
}
