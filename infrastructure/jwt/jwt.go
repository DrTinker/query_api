package jwt

import (
	"query_api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserPwd string `json:"user_pwd"`
	UserID  int32  `json:"user_id"`
	jwt.StandardClaims
}

func Encode(t models.Token, keys []byte, expire int64) (string, error) {
	c := &Claims{}
	// 拼接claims
	if t.Expire == 0 {
		c.ExpiresAt = time.Now().Unix() + expire
	}
	c.UserID = t.UserID
	c.UserPwd = t.UserPwd

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := tokenClaims.SignedString(keys)

	return token, err
}

func Decode(token string, keys []byte) (*models.Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return keys, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return &models.Token{
				UserPwd: claims.UserPwd,
				UserID:  claims.UserID,
				Expire:  claims.ExpiresAt,
			}, nil
		}
	}

	return nil, err
}
