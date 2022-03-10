package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"query_api/client"
	"query_api/conf"
)

// 解析jwt，flag为true标识拦截不带token的请求
func JWT(flag bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = conf.HTTP_SUCCESS_CODE
		token := c.GetHeader(conf.JWTHeader)
		// 未携带token
		if token == "" {
			// 拦截
			if flag {
				c.Abort()
				return
			}
			// 标识未携带token登录
			c.Set(conf.JWTFlag, false)
			c.Next()
			return
		}
		// 解析token
		t, err := client.EncryptionClient.JwtDecode(token)
		if err != nil {
			code = conf.ERROR_AUTH_CHECK_TOKEN_FAIL_CODE
		} else if time.Now().Unix() > t.Expire {
			code = conf.ERROR_AUTH_CHECK_TOKEN_TIMEOUT_CODE
		}
		// token无效
		if code != conf.HTTP_SUCCESS_CODE {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  conf.JWT_ERROR_MESSAGE,
			})
			c.Abort()
			return
		}
		// 标识携带token登录
		c.Set(conf.User_ID, t.UserID)
		c.Set(conf.User_PWD, t.UserPwd)
		c.Set(conf.JWTFlag, true)
		c.Next()
	}
}
