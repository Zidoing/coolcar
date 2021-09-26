package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issue      string
	nowFunc    func() time.Time
}

func NewJWTTokenGen(issuer string, private *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issue:      issuer,
		nowFunc:    time.Now,
		privateKey: private,
	}

}

func (t *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	now := t.nowFunc().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.issue,
		IssuedAt:  now,
		ExpiresAt: now + int64(expire.Seconds()),
		Subject:   accountID,
	})

	fmt.Println(expire)

	return token.SignedString(t.privateKey)

}
