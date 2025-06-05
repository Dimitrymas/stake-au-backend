package utils

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 2, 2, 3, 3, 3, 4}
	expected := []int{1, 2, 3, 4}
	result := RemoveDuplicates(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGenerateKeyPairAndLoad(t *testing.T) {
	priv, _, err := GenerateKeyPair(1024)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	enc := PrivateKeyToString(priv)
	got, err := LoadPrivateKey(enc)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if priv.N.Cmp(got.N) != 0 {
		t.Error("loaded key mismatch")
	}
}

func TestPublicKeyToString(t *testing.T) {
	priv, pub, err := GenerateKeyPair(1024)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	s, err := PublicKeyToString(pub)
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	if s == "" {
		t.Error("empty string")
	}
	// ensure LoadPrivateKey of priv doesn't affect
	_ = priv
}
