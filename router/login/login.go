package login

import (
	"fmt"
	"query_api/models"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// 初始化user struct
	u := models.User{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	err := c.ShouldBind(&u) 
	if err != nil {
		log.SetLevel(log.DebugLevel)
		log.Error("err: %+v", err)
	}
	fmt.Printf("user: %+v", u)
	// http 请求返回一个字符串 
	c.String(200, "Success")
}