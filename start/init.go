package start

import (
	"query_api/client"
	"query_api/infrastructure/rpc"
	middleware "query_api/middleware/user"
	"query_api/router/login"
	"query_api/router/query"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置
	initConfig()
	initRPC()
	initJWT()

	// 业务逻辑接口初始化
	client.InitUserClient(rpc.NewUserServiceClientImpl())
	client.InitQueryClient(rpc.NewQueryClientImpl())
}

func RegisterRouter(r *gin.Engine) {
	l := r.Group("/user")
	{
		l.GET("/login", middleware.JWT(false), login.LoginHandler)
		l.POST("/register", login.RegisterHandler)
	}

	q := r.Group("/query")
	{
		q.GET("/get", middleware.JWT(true), query.GetQueryHandler)
		q.POST("/create", middleware.JWT(true), query.CreateQueryHandler)
	}
}
