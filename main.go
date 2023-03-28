package main

import (
	rt "tiktok/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := new(rt.Router)
	go router.InitRouter(r)
	r.Run(":8000")
}
