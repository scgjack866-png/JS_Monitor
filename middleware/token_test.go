package middleware

import "testing"

func TestParseTokenWithBareToken(t *testing.T) {
	token := CreateToken(map[string]string{"uuid": "u1", "exp": "9999999999"})
	claims, ok := ParseToken(token)
	if !ok {
		t.Fatal("expected bare token to parse successfully")
	}
	if claims["uuid"] != "u1" {
		t.Fatalf("unexpected uuid: %v", claims["uuid"])
	}
}

func TestParseTokenWithBearerToken(t *testing.T) {
	token := CreateToken(map[string]string{"uuid": "u2", "exp": "9999999999"})
	claims, ok := ParseToken("Bearer " + token)
	if !ok {
		t.Fatal("expected bearer token to parse successfully")
	}
	if claims["uuid"] != "u2" {
		t.Fatalf("unexpected uuid: %v", claims["uuid"])
	}
}

func TestParseTokenWithInvalidHeader(t *testing.T) {
	if claims, ok := ParseToken("Bearer"); ok || claims != nil {
		t.Fatal("expected invalid authorization header to fail")
	}
}
