package login

import (
	"net/http"
	"query_api/client"
	"query_api/conf"
	"query_api/grpc_gen/user"
	"query_api/models"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// TODO 邮箱手机认证
func RegisterHandler(c *gin.Context) {
	// 初始化user struct
	// 不能直接用grpc生成的结构
	u := models.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		log.Error("RegisterHandler err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": conf.HTTP_INVALID_PARAMS_CODE,
			"msg":  conf.HTTP_INVALID_PARAMS_MESSAGE,
		})
		return
	}
	// 调用client
	user := &user.User{
		UserPwd:  u.User_pwd,
		UserName: u.User_name,
		Phone:    int64(u.User_phone),
		UserTag:  u.User_tag,
		Pass:     u.Pass,
		Log:      u.Log,
		Email:    u.User_email,
	}
	err = client.GetUserClient().CreateUser(c, user)
	log.Error("CreateUser err: ", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": conf.RPC_FAILED_CODE,
			"msg":  conf.RPC_FAILED_MESSAGE,
		})
		return
	}
	// 返回成功
	log.Error("RegisterHandler success: %v", u.User_phone)
	c.JSON(http.StatusOK, gin.H{
		"code": conf.HTTP_SUCCESS_CODE,
		"msg":  conf.SUCCESS_RESP_MESSAGE,
	})
}
