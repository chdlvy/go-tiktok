package router

import (
	"github.com/gin-gonic/gin"

	middleware "tiktok/middleware"
)

type Router struct {
}

func (this *Router) InitRouter(r *gin.Engine) {
	r.POST("/login", this.login)
	r.POST("/sign", this.sign)
	r.Use(middleware.VerifyTokenMiddleWare())
}

func (this *Router) login(c *gin.Context) {

}

func (this *Router) sign(c *gin.Context) {

}
