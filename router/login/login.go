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
	var token string
	// 处理jwt token
	if c.GetBool(conf.JWTFlag) {
		claims, e := c.Get(conf.JWTClaims)
		if !e {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"code": conf.ErrJWTTokenInvaild,
				"msg":  conf.JWT_ERROR_MESSAGE,
			})
		}
		// 类型推断
		t := claims.(*models.Token)
		u.User_id = t.UserID
		u.User_pwd = t.UserPwd
	} else {
		// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
		err := c.ShouldBind(&u)
		if err != nil {
			log.SetLevel(log.DebugLevel)
			log.Error("err: %+v\n", err)
		}
	}
	id := u.User_id
	pwd := u.User_pwd

	info, err := client.UserOptionClient.GetUserByUserID(c, id)
	fmt.Printf("GetUserByUserID rpc error: %+v\n", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": conf.ERROR_LOGIN_CODE,
			"msg":  conf.RPC_FAILED_MESSAGE,
		})
		return
	}
	// 密码错误
	if info.UserPwd != pwd {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": conf.ERROR_LOGIN_CODE,
			"msg":  conf.LOGIN_ERROR_MESSAGE,
		})
	}
	// TODO 介入日志
	fmt.Printf("user: %+v", u)
	if !c.GetBool(conf.JWTFlag) {
		token, _ = client.EncryptionClient.JwtEncode(models.Token{
			UserID:  id,
			UserPwd: pwd,
			Expire:  0,
		})
	}
	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":  conf.HTTP_SUCCESS_CODE,
		"msg":   conf.SUCCESS_RESP_MESSAGE,
		"data":  info,
		"token": token,
	})
}
