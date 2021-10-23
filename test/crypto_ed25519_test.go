package harpocrates_test

import (
	"sync"
	"testing"

	harpocrates "github.com/TunnelWork/Harpocrates"
)

// TestEd25519Key() evaluates if the function creates consistent results.
func TestEd25519Key(t *testing.T) {
	key1 := harpocrates.Ed25519Key("Hello World")
	key2 := harpocrates.Ed25519Key("Hello World")
	key3 := harpocrates.Ed25519Key("")

	if key1 == nil || !key1.Equal(key2) {
		t.Errorf("Ed25519Key() creates inconsistent key or nil key for valid seeds")
	}

	if key3 != nil {
		t.Errorf("Ed25519Key() creates non-nil key for empty seed")
	}
}

// TestEd25519PubKey() evaluates if the function creates consistent results.
func TestEd25519PubKey(t *testing.T) {
	key1 := harpocrates.Ed25519PubKey("Hello World")
	key2 := harpocrates.Ed25519PubKey("Hello World")
	key3 := harpocrates.Ed25519PubKey("")

	if key1 == nil || !key1.Equal(key2) {
		t.Errorf("Ed25519PubKey() creates inconsistent key or nil key for valid seeds")
	}

	if key3 != nil {
		t.Errorf("Ed25519PubKey() creates non-nil key for empty seed")
	}
}

// TestEd25519SignWithSeed() evaluates if the Signature created is consistent.
func TestEd25519SignWithSeed(t *testing.T) {
	sig1, err1 := harpocrates.Ed25519SignWithSeed("My Message!", "Hello World")
	sig2, err2 := harpocrates.Ed25519SignWithSeed("My Message!", "Hello World")
	sig3, err3 := harpocrates.Ed25519SignWithSeed("My Message!", "")
	sig4, err4 := harpocrates.Ed25519SignWithSeed("", "Hello World")

	if sig1 != sig2 || sig1 == "" || err1 != nil || err2 != nil {
		t.Errorf("Ed25519SignWithSeed() can't create consistent signature for valid seeds")
	}

	if sig3 != "" || err3 == nil {
		t.Errorf("Ed25519SignWithSeed() didn't fail properly for invalid seed input")
	}

	if sig4 == "" || err4 != nil {
		t.Errorf("Ed25519SignWithSeed() failed on empty msg")
	}
}

func TestEd25519VerifyWithSeed(t *testing.T) {
	goodseed := "Hello World"

	msg1 := "My Message"
	msg2 := "" // empty msg to fail?
	msg3 := "Some Other Msg"

	sig1, _ := harpocrates.Ed25519SignWithSeed(msg1, goodseed)
	sig2, _ := harpocrates.Ed25519SignWithSeed(msg2, goodseed)
	sig3 := sig2

	if err := harpocrates.Ed25519VerifyWithSeed(msg1, sig1, goodseed); err != nil {
		t.Errorf("Signature 1 should be valid, but errored with %s", err)
	}
	if err := harpocrates.Ed25519VerifyWithSeed(msg2, sig2, goodseed); err != nil {
		t.Errorf("Signature 2 should be valid, but errored with %s", err)
	}
	if harpocrates.Ed25519VerifyWithSeed(msg3, sig3, goodseed) == nil {
		t.Errorf("Signature 3 should not be valid.")
	}
}

func BenchmarkEd25519ParallelSignThenVerify(b *testing.B) {
	seed := "Signed By Gaukas Wang"
	msg := "Gaukas Wang: Hello World."

	wg := sync.WaitGroup{}
	wg.Add(b.N)

	bench := func() {
		sig, _ := harpocrates.Ed25519SignWithSeed(msg, seed)
		harpocrates.Ed25519VerifyWithSeed(msg, sig, seed)
		wg.Done()
	}

	for n := 0; n < b.N; n++ {
		go bench()
		// bench()
	}

	wg.Wait()
}

func BenchmarkEd25519ParallelSign(b *testing.B) {
	seed := "Signed By Gaukas Wang"
	msg := "Gaukas Wang: Hello World."

	wg := sync.WaitGroup{}
	wg.Add(b.N)

	bench := func() {
		harpocrates.Ed25519SignWithSeed(msg, seed)
		wg.Done()
	}

	for n := 0; n < b.N; n++ {
		go bench()
		// bench()
	}

	wg.Wait()
}
