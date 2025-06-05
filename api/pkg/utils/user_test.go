package utils

import (
	"backend/api/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"testing"
)

func TestHashPasswordAndCheck(t *testing.T) {
	hash, err := HashPassword("pass")
	if err != nil {
		t.Fatalf("hash error: %v", err)
	}
	if !CheckPasswordHash("pass", hash) {
		t.Error("password check failed")
	}
}

func TestEncodeDecodeJWT(t *testing.T) {
	config.S.JwtSecret = []byte("secret")
	id := primitive.NewObjectID()
	token, err := EncodeJWT(id)
	if err != nil {
		t.Fatalf("encode error: %v", err)
	}
	got, err := GetUserIdFromToken(token)
	if err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if got != id {
		t.Errorf("expected %s got %s", id.Hex(), got.Hex())
	}
}

func TestGenerateMnemonicAndValidate(t *testing.T) {
	mnemonic, err := GenerateMnemonic()
	if err != nil {
		t.Fatalf("generate error: %v", err)
	}
	if len(mnemonic) != 12 {
		t.Errorf("expected 12 words, got %d", len(mnemonic))
	}
	if !ValidateMnemonic(mnemonic) {
		t.Error("generated mnemonic should be valid")
	}
}

func TestMnemonicToSeed(t *testing.T) {
	phrase := strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about", " ")
	seed := MnemonicToSeed(phrase)
	expected := "62a772f85e4be6226108b56c0b1cf935c2490e434adec864fe47b189f1ed517d"
	if seed != expected {
		t.Errorf("seed mismatch: %s", seed)
	}
}
