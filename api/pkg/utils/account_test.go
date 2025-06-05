package utils

import (
	"backend/api/pkg/models"
	"testing"
)

func TestBuildAccountProxyString(t *testing.T) {
	acc := &models.Account{ProxyIP: "1.1.1.1", ProxyPort: "8080", ProxyLogin: "u", ProxyPass: "p"}
	if s := BuildAccountProxyString(acc); s != "u:p@1.1.1.1:8080" {
		t.Errorf("unexpected string: %s", s)
	}

	acc.ProxyLogin = ""
	acc.ProxyPass = ""
	if s := BuildAccountProxyString(acc); s != "1.1.1.1:8080" {
		t.Errorf("unexpected string without creds: %s", s)
	}

	acc.ProxyIP = ""
	if s := BuildAccountProxyString(acc); s != "" {
		t.Errorf("expected empty string, got: %s", s)
	}
}
