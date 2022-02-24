package login

import (
	"fmt"
	"net/http"
	"query_api/client"
	"query_api/conf"
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
	id := u.User_id
	pwd := u.User_pwd

	info, err := client.UserOptionClient.GetUserByUserID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": conf.ERROR_LOGIN_CODE,
			"msg": conf.RPC_FAILED_MESSAGE,
		})
	}
	// 密码错误
	if (info.UserPwd != pwd) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": conf.ERROR_LOGIN_CODE,
			"msg": conf.LOGIN_ERROR_MESSAGE,
		})
	}
	// TODO 介入日志
	fmt.Printf("user: %+v", u)
	// 返回成功 
	c.JSON(http.StatusNotAcceptable, gin.H{
		"code": conf.HTTP_SUCCESS_CODE,
		"msg": conf.SUCCESS_RESP_MESSAGE,
		"data": info, 
	})
}