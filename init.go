package main

import (
	"query_api/client"
	"query_api/conf"
	"query_api/models"
	"query_api/pkg/config"
	"query_api/router/login"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func init() {
	// 初始化http服务端ip端口配置
	config.HttpConfig = config.HttpConfig.Load(conf.App).InitHttp()
	// 初始化RPC连接
	config.RpcConfig = config.RpcConfig.Load(conf.App).InitRpc()
	models.RpcConn = rpcConnect()
	// 初始化jwt
	client.EncryptionClient.JWTInit(conf.JWTExpireValue, []byte(conf.JWTKeyValue))
}

func RegisterRouter(r *gin.Engine) {
	l := r.Group("/user") 
	{
		l.POST("/login", login.LoginHandler)
	}
}

func rpcConnect() *grpc.ClientConn{
	conn, err := grpc.Dial(config.RpcConfig.Address, grpc.WithInsecure())
	
    if err != nil {
		panic(err)
    }
	return conn
}