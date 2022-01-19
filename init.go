package main

import (
	client "query_api/client"
	"query_api/conf"
	"query_api/pkg/config"
	"query_api/router/login"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func init() {
	// 初始化http服务端ip端口配置
	config.HttpServer = config.HttpServer.Load(conf.App).InitByAddress(conf.HttpServer, conf.DefaultIp, conf.DefaultHttpPort)
	// 初始化RPC连接
	config.RpcServer = config.RpcServer.Load(conf.App).InitByAddress(conf.RpcConnect, conf.DefaultIp, conf.DefaultRpcPort)
	client.RpcConn = rpcConnect()
}

func RegisterRouter(r *gin.Engine) {
	l := r.Group("/user") 
	{
		l.POST("/login", login.LoginHandler)
	}
}

func rpcConnect() *grpc.ClientConn{
	conn, err := grpc.Dial(config.RpcServer.Address, grpc.WithInsecure())
	
    if err != nil {
		panic(err)
    }
	return conn
}