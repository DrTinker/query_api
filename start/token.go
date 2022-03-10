package start

import (
	"query_api/client"
	"query_api/conf"
)

func initJWT() {
	// 初始化jwt
	client.EncryptionClient.JWTInit(conf.JWTExpireValue, []byte(conf.JWTKeyValue))
}
