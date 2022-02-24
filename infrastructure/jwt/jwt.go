package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"user_name"`
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func Encode(c Claims, keys []byte, expire int64) (string, error) {
	if c.ExpiresAt == 0 {
		c.ExpiresAt = time.Now().Unix() + expire
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    token, err := tokenClaims.SignedString(keys)

	return token, err
}

func Decode(token string, keys []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return keys, nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}