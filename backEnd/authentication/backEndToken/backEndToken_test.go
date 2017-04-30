package backEndToken_test

import (
	"github.com/zicongmei/angularTest/backEnd/authentication/backEndToken"
	"net/http"
	"testing"
)

func TestBuildToken(t *testing.T) {
	_, err := backEndToken.BuildToken("user1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckToken(t *testing.T) {
	testUsername := "user1"
	token, err := backEndToken.BuildToken(testUsername)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Add("Authorization", token)
	claim, err := backEndToken.CheckToken(req)
	if err != nil {
		t.Fatal(err)
	}
	if content, ok := claim["user"]; !ok {
		t.Fatal("user not in claim")
	} else if username, ok := content.(string); !ok {
		t.Fatal("user not in string")
	} else if username != testUsername {
		t.Fatal("user not match")
	}
}
