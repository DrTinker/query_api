package jwt

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"

	"query_api/conf"
	"query_api/client"
)

// 解析jwt
func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int

        code = conf.HTTP_SUCCESS_CODE
        token := c.Query("token")
        if token == "" {
            code = conf.HTTP_INVALID_PARAMS_CODE
        } else {
			// 解析token
            claims, err := client.EncryptionClient.JwtDecode(token)
            if err != nil {
                code = conf.ERROR_AUTH_CHECK_TOKEN_FAIL_CODE
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = conf.ERROR_AUTH_CHECK_TOKEN_TIMEOUT_CODE
            }
        }

        if code != conf.HTTP_SUCCESS_CODE {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : code,
                "msg" : conf.JWT_ERROR_MESSAGE,
            })

            c.Abort()
            return
        }

        c.Next()
    }
}