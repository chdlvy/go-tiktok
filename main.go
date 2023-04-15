package main

import (
	"tiktok/internal/pkg/redis"
	sql "tiktok/orm"
	rt "tiktok/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := new(rt.Router)
	go router.InitRouter(r)
	go redis.InitRedis()
	go sql.InitDB()
	r.Run(":8000")

	// test()
}

func test() {

}
