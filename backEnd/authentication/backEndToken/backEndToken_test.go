package backEndToken_test

import (
	"testing"
	"github.com/zicongmei/angularTest/backEnd/authentication/backEndToken"
	"fmt"
)

func TestBuildToken(t *testing.T) {
	str, err := backEndToken.BuildToken("user1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(str)
}
