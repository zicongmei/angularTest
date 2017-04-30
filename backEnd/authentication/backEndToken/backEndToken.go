package backEndToken

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"net/http"
	"time"
)

var rsaKey *rsa.PrivateKey
var rsaPublicKey *rsa.PublicKey

func init() {
	reader := rand.Reader
	bitSize := 2048
	var err error
	rsaKey, err = rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic("Can't generate RSA key: " + err.Error())
	}
	rsaPublicKey = &rsaKey.PublicKey
}

func BuildToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"user": userName,
		"nbf":  time.Now().Add(2 * time.Hour),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(rsaKey)
	return tokenString, err
}

func CheckToken(r *http.Request) (jwt.MapClaims, error) {
	token, err := jwtrequest.ParseFromRequest(r, jwtrequest.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return rsaPublicKey, nil
		})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return make(jwt.MapClaims), err
	}
}
