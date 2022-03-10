package query

import (
	"net/http"
	"query_api/client"
	"query_api/conf"
	"query_api/grpc_gen/query"
	"query_api/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateQueryHandler(c *gin.Context) {
	q := models.Query{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	err := c.ShouldBind(&q)
	if err != nil {
		logrus.Error("CreateQueryHandler err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": conf.HTTP_INVALID_PARAMS_CODE,
			"msg":  conf.HTTP_INVALID_PARAMS_MESSAGE,
		})
		return
	}
	// 从登录态获取id
	var id int32
	if data, ok := c.Get(conf.User_ID); ok {
		id, _ = data.(int32)
	}
	query := &query.Query{
		QueryId:    q.Query_ID,
		QueryName:  q.Query_Name,
		State:      q.State,
		Remark:     q.Remark,
		StartTime:  q.Start_Time,
		EndTime:    q.End_Time,
		EndMethod:  q.End_Method,
		Background: q.Background,
		CreateTime: q.Create_Time,
		Creator:    q.Creator,
	}
	err = client.GetQueryClient().CreateQuery(c, int32(id), query)
	if err != nil {
		logrus.Error("CreateQueryHandler err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": conf.RPC_FAILED_CODE,
			"msg":  conf.RPC_FAILED_MESSAGE,
		})
		return
	}
	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code": conf.HTTP_SUCCESS_CODE,
		"msg":  conf.SUCCESS_RESP_MESSAGE,
	})
	logrus.Error("CreateQueryHandler err: ", err)
}
