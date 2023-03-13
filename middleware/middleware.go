package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证是否登录过、是否有token
func VerifyTokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.Request.Header.Get("authorization")
		if authHandler == "" {
			c.JSON(200, gin.H{"code": 200, "msg": "请求头部auth为空，登录过期，重新登录"})
			c.Redirect(http.StatusMovedPermanently, "127.0.0.1:8888/login")
			c.Abort()
			return
		}

	}
}
