package backEndToken

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"crypto/rand"
)

var rsaKey *rsa.PrivateKey

func init(){
	reader := rand.Reader
	bitSize := 2048
	var err error
	rsaKey, err = rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic ("Can't generate RSA key: " + err.Error())
	}
}

func BuildToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": userName,
		"expires": time.Now().Add(2 * time.Hour),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(rsaKey)
	return tokenString, err
}
