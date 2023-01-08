package accesstoken

import "testing"

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")

	}

	if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}
}