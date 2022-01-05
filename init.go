package main

import (
	"query_api/conf"
	"query_api/pkg/config"
	"query_api/router/login"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化ip端口配置
	config.Server = config.Server.Load(conf.App).Init()
}

func RegisterRouter(r *gin.Engine) {
	l := r.Group("/user") 
	{
		l.POST("/login", login.LoginHandler)
	}
}