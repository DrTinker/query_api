package response

import (
	"net/http"
	"query_api/conf"

	"github.com/gin-gonic/gin"
)

func Success_resp(c *gin.Context) {
	c.JSON(http.StatusOK, conf.SUCCESS_RESP_MESSAGE)
}

