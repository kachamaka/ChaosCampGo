package tokens_test

import (
	"testing"

	"github.com/kachamaka/chaosgo/tokens"
)

var (
	TOKEN_SECRET = tokens.GenerateSecret(32)
	TEST_ID      = "63ea64ef2e1c7d21f929d50d"
	TOKEN_STRING = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2M2VhNjRlZjJlMWM3ZDIxZjkyOWQ1MGQifQ.EQMbISHypQddPriTIHuwOdX_3V5lIk1lE04LlKawkM8"
)

func TestTokens(t *testing.T) {
	tokenString, err := tokens.GenerateToken(TEST_ID)
	if err != nil {
		t.Error("generate token error:", err)
	}
	if tokenString != TOKEN_STRING {
		t.Error("diff token string")
	}
	claims, err := tokens.DecryptToken(tokenString)
	if err != nil {
		t.Error("error decrypting token:", err)
	}
	id, ok := claims["_id"]
	if !ok {
		t.Error("no id in token")
	}
	if id != TEST_ID {
		t.Error("ids don't match")
	}
}
