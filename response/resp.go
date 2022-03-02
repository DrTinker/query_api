package response

import (
	"net/http"
	"query_api/conf"

	"github.com/gin-gonic/gin"
)

// 通用成功响应体
func Success_resp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": conf.HTTP_SUCCESS_CODE,
		"msg":  conf.SUCCESS_RESP_MESSAGE,
	})
}
