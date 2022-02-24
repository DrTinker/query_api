package main

// 导入gin包
import (
	"fmt"
	"query_api/pkg/config"

	"github.com/gin-gonic/gin"
)

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	RegisterRouter(r)

	r.Run(fmt.Sprintf("%s:%d", config.HttpConfig.Address, config.HttpConfig.Port))
}
