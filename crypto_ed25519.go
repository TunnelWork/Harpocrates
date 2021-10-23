package harpocrates

import (
	"crypto/ed25519"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ed25519Master = jwt.SigningMethodEd25519{}

	ErrKeyIsNil error = errors.New("harpocrates: can't sign/verify when key is nil")
)

// Ed25519Key() creates a ed25519.PrivateKey from seed
// if the seed is longer than 32 bytes, it will be truncated.
// otherwise, auto-extended by repeating itself
// May return nil if seed is empty
func Ed25519Key(seed string) ed25519.PrivateKey {
	if len(seed) == 0 {
		return nil
	}

	var byteSeed []byte
	byteSeed, _ = ResizeByteArray([]byte(seed), 32)

	return ed25519.NewKeyFromSeed(byteSeed)
}

// Ed25519PubKey() creates a ed25519.PublicKey from seed
// if the seed is longer than 32 bytes, it will be truncated.
// otherwise, auto-extended by repeating itself
// May return nil if seed is empty
func Ed25519PubKey(seed string) ed25519.PublicKey {
	if key := Ed25519Key(seed); key != nil {
		return key.Public().(ed25519.PublicKey)
	}
	return nil
}

// Ed25519SignWithSeed() sign a msg string with a key created from the seed
func Ed25519SignWithSeed(msg, seed string) (string, error) {
	signingKey := Ed25519Key(seed)
	if signingKey == nil {
		return "", ErrKeyIsNil
	}
	return ed25519Master.Sign(msg, signingKey)
}

// Ed25519SignWithSeed() sign a msg string with a key created from the seed
func Ed25519VerifyWithSeed(msg, signature, seed string) error {
	veriyingKey := Ed25519PubKey(seed)
	if veriyingKey == nil {
		return ErrKeyIsNil
	}
	return ed25519Master.Verify(msg, signature, veriyingKey)
}
