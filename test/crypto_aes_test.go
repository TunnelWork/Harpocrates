package harpocrates_test

import (
	"reflect"
	"testing"

	harpocrates "github.com/TunnelWork/Harpocrates"
)

func TestInAESKeySize(t *testing.T) {
	key0 := []byte("")                                      // Empty, should get nil
	key1 := []byte("AKeyTooShort")                          // Should get a result with len == 16
	key2 := []byte("LongerKeyShouldBeUsed")                 // Should get len == 24
	key3 := []byte("LongerThanTheLongKeyForBetterSecurity") // Should get len == 32

	res0 := harpocrates.InAESKeysize(key0)
	res1 := harpocrates.InAESKeysize(key1)
	res2 := harpocrates.InAESKeysize(key2)
	res3 := harpocrates.InAESKeysize(key3)

	if res0 != nil {
		t.Errorf("Failed to return nil for input with zero length.")
	}

	if len(res1) != 16 {
		t.Errorf("Failed to extend to 16 byte key, got %d\n", len(res1))
	}

	if len(res2) != 24 {
		t.Errorf("Failed to extend to 24 byte key, got %d\n", len(res2))
	}

	if len(res3) != 32 {
		t.Errorf("Failed to extend to 32 byte key, got %d\n", len(res3))
	}
}

func TestNewAESCipher(t *testing.T) {
	iv := []byte("SomeRandomIV")
	key := []byte("SomeRandomKey")

	aesCipher := harpocrates.NewAESCipher(iv, key, harpocrates.CBC)
	if aesCipher == nil {
		t.Errorf("Can't get new AES cipher.")
	}

	ptext1 := []byte("Hello World! My name is Gaukas Wang.")
	ctext1, err1 := aesCipher.Encrypt(ptext1)
	if err1 != nil {
		t.Errorf("Can't encrypt ptext1, error: %s", err1)
	}
	ptext1r, err1 := aesCipher.Decrypt(ctext1)
	if err1 != nil {
		t.Errorf("Can't decrypt ctext1, error: %s", err1)
	}
	if !reflect.DeepEqual(ptext1, ptext1r) {
		t.Errorf("ptext1r: %s, but ptext1: %s", string(ptext1r), string(ptext1))
	}

	ptext2 := "Gaukas Wang likes the idea of Hex digest."
	ctext2, err2 := aesCipher.HexDigestEncrypt(ptext2)
	if err2 != nil {
		t.Errorf("Can't HexDigestEncrypt ptext2, error: %s", err2)
	}
	ptext2r, err2 := aesCipher.HexDigestDecrypt(ctext2)
	if err2 != nil {
		t.Errorf("Can't HexDigestDecrypt ctext2, error: %s", err2)
	}
	if !reflect.DeepEqual(ptext2, ptext2r) {
		t.Errorf("ptext2r: %s, but ptext2: %s", ptext2r, ptext2)
	}
}
