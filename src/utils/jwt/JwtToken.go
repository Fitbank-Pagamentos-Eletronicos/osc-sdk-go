package jwt

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

const mySigningKey = "S42@874smy@@ox0ypj6g"

type claims struct {
	Tag_name  string   `json:"tag_name"`
	Scopes    []string `json:"scopes"`
	jwt.StandardClaims
}

type Token struct {
	Access_token string `json:"access_token"`
	Expire_at    string `json:"expire_at"`
}

func GenerateToken(tag_name string, scopes []string) Token {

	expiresAt := time.Now().Add(24 * time.Hour)
	claims := claims{
		tag_name,
		scopes,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	singnedToken, _ := token.SignedString([]byte(mySigningKey))

	return Token{
		Access_token: singnedToken,
		Expire_at:    expiresAt.String(),
	}
}
