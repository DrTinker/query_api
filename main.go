package main

// 导入gin包
import (
	"fmt"
	"query_api/client"
	"query_api/start"

	"github.com/gin-gonic/gin"
)

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	start.RegisterRouter(r)

	cfg, err := client.GetConfigClient().GetHttpConfig()
	if err != nil {
		panic(err)
	}

	r.Run(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port))
}
