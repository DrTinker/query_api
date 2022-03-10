package login

import (
	"net/http"
	"query_api/client"
	"query_api/conf"
	"query_api/models"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// 初始化user struct
	u := models.Login{}
	var token string
	// 处理jwt token
	if c.GetBool(conf.JWTFlag) {
		if id, ok := c.Get(conf.User_ID); ok && id != nil {
			u.User_id, _ = id.(int32)
		}
		if pwd, ok := c.Get(conf.User_PWD); ok && pwd != nil {
			u.User_pwd, _ = pwd.(string)
		}
	} else {
		err := c.ShouldBind(&u)
		if err != nil {
			log.Error("LoginHandler err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": conf.HTTP_INVALID_PARAMS_CODE,
				"msg":  conf.HTTP_INVALID_PARAMS_MESSAGE,
			})
		}
	}
	id := u.User_id
	pwd := u.User_pwd

	info, err := client.GetUserClient().GetUserByUserID(c, id)
	if err != nil {
		log.Error("GetUserByUserID rpc err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": conf.RPC_FAILED_CODE,
			"msg":  conf.RPC_FAILED_MESSAGE,
		})
		return
	}
	// 密码错误
	if info.UserPwd != pwd {
		log.Info("LoginHandler pwd err: %+v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": conf.ERROR_LOGIN_CODE,
			"msg":  conf.LOGIN_ERROR_MESSAGE,
		})
		return
	}
	if !c.GetBool(conf.JWTFlag) {
		token, _ = client.EncryptionClient.JwtEncode(models.Token{
			UserID:  id,
			UserPwd: pwd,
			Expire:  0,
		})
	}
	// 返回成功
	log.Error("LoginHandler success: %v", u.User_id)
	c.JSON(http.StatusOK, gin.H{
		"code":  conf.HTTP_SUCCESS_CODE,
		"msg":   conf.SUCCESS_RESP_MESSAGE,
		"data":  info,
		"token": token,
	})
}
