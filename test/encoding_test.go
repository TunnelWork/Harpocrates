package harpocrates_test

import (
	"reflect"
	"testing"

	harpocrates "github.com/TunnelWork/Harpocrates"
)

func TestBase64EncodingDecoding(t *testing.T) {
	originalText := "Hello This Is Gaukas Wang!"
	ExpectedB64 := "SGVsbG8gVGhpcyBJcyBHYXVrYXMgV2FuZyE="
	GeneratedB64 := harpocrates.Base64Encoding(originalText)
	DecodedText := harpocrates.Base64Decoding(GeneratedB64)

	if !reflect.DeepEqual(originalText, DecodedText) {
		t.Errorf("Can't decode the encoded result. Original: %s, Decoded: %s\n", originalText, DecodedText)
	}

	if !reflect.DeepEqual(ExpectedB64, GeneratedB64) {
		t.Errorf("Encoding failed. Expected: %s, Got: %s\n", ExpectedB64, GeneratedB64)
	}
}

func TestIsBase64(t *testing.T) {
	SomeB64 := "SGVsbG8gVGhpcyBJcyBHYXVrYXMgV2FuZyE="
	SomethingNotB64 := "SGVsbG8gVGhpcyBJcyBHYXVrYXMg_2FuZyE="

	if !harpocrates.IsBase64(SomeB64) {
		t.Errorf("Can't identify Base64 encoded string correctly.")
	}

	if harpocrates.IsBase64(SomethingNotB64) {
		t.Errorf("Can't identify Non-Base64 encoded string correctly.")
	}
}
