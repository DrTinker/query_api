package query

import (
	"net/http"
	"query_api/client"
	"query_api/conf"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 获取指定问卷信息
func GetQueryHandler(c *gin.Context) {
	data, _ := c.GetQuery("id")
	id, err := strconv.Atoi(data)
	if err != nil {
		logrus.Error("GetQueryHandler err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": conf.HTTP_INVALID_PARAMS_CODE,
			"msg":  conf.HTTP_INVALID_PARAMS_MESSAGE,
		})
		return
	}
	q, err := client.GetQueryClient().GetQueryByID(c, int32(id))
	if err != nil {
		logrus.Error("CreateQueryHandler err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": conf.RPC_FAILED_CODE,
			"msg":  conf.RPC_FAILED_MESSAGE,
		})
		return
	}
	// 返回成功
	logrus.Error("LoginHandler success: %v", q.QueryId)
	c.JSON(http.StatusOK, gin.H{
		"code": conf.HTTP_SUCCESS_CODE,
		"msg":  conf.SUCCESS_RESP_MESSAGE,
		"data": q,
	})
}

// 获取问卷及全部问题，用于问卷编辑界面
func GetQueryQuestionHandler(c *gin.Context) {

}
