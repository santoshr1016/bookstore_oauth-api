package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("Expiration token should be  24")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Access token should not be expired, its just created ")
	}
	if at.AccessToken != "" {
		t.Error("Access token should not have defined token id")
	}
	if at.UserId != 0 {
		t.Error("Access token should not have defined user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired(){
		t.Error("empty access token, expired")
	}

	at.Expires = time.Now().UTC().Add(3*time.Hour).Unix()
	if at.IsExpired(){
		t.Error("Token created with the expiry 3 hr")
	}
}
