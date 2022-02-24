package client

import (
	"query_api/infrastructure/jwt"
)

type Encryption interface {
	JWTInit (expire int64, key []byte)
	JwtEncode (c jwt.Claims) (string, error)
	JwtDecode (s string) (*jwt.Claims, error)
}

type encryption struct {
	JWTKey []byte
	JWTExpire int64
}

var EncryptionClient encryption

func (e *encryption) JwtEncode (c jwt.Claims) (string, error){
	token, err := jwt.Encode(c, e.JWTKey, e.JWTExpire)
	return token, err
}

func (e *encryption) JwtDecode (s string) (*jwt.Claims, error){
	claims, err := jwt.Decode(s, e.JWTKey)
	return claims, err
}

func (e *encryption) JWTInit (expire int64, key []byte) {
	e.JWTKey = key
	e.JWTExpire = expire
}
