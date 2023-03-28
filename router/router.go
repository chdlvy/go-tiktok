package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"tiktok/src"
	"tiktok/utils"

	"tiktok/internal/pkg/redis"

	redisgo "github.com/garyburd/redigo/redis"
)

type Router struct {
}

var userMap map[uint64]*src.User

func (this *Router) InitRouter(r *gin.Engine) {
	r.POST("/login", this.login)
	r.POST("/sign", this.sign)
	r.POST("/publish", this.publish)
	r.GET("/getVideoList", this.getVideoList)
	// r.Use(middleware.VerifyTokenMiddleWare())
}

func (this *Router) login(c *gin.Context) {

}

func (this *Router) sign(c *gin.Context) {
	var newUser src.ModUser
	c.ShouldBind(&newUser)
	user := src.InitUser(&newUser)
	go src.GetUserById(user.ModUser.Id)

}

func (this *Router) publish(c *gin.Context) {
	var newVideo src.ModVideo
	c.ShouldBind(&newVideo)

	user := new(src.ModUser)
	// 根据id获取对应的user信息
	res, err := redis.GetUser(newVideo.WriterId)
	if err != nil {
		panic(err)
	}
	// map转结构体
	utils.MapToStruct(res, user)
	// 用户发布的作品数量+1
	user.PublishCount += 1

	// videoId := user.ModUser.PublishCount

	// 将视频信息存入redis
	// video为hmap名，用来存放所有视频信息
	// key：user:1:video:1
	// value：json序列化的视频信息
	value, _ := json.Marshal(newVideo)
	redis.SaveVideoToRedis(user.Id, string(value))

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func (this *Router) getVideoList(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	start := page * 10
	end := page*10 + 10

	list := getVideo(start, end)
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}

// 从redis中拿数据，每次返回10条
func getVideo(start int64, end int64) []map[string]interface{} {
	rdb := redis.GetredisConn()
	defer rdb.Close()
	res, err := redisgo.Strings(rdb.Do("hvals", "video"))
	if err != nil {
		fmt.Println("获取数据失败")
		panic(err)
	}
	// 防止end溢出
	if end > int64(len(res)) {
		end = int64(len(res))
	}

	// 将获取的数据解析成正常的json格式
	var result []map[string]interface{}
	for i := start; i < end; i++ {
		var v map[string]interface{}
		err = json.Unmarshal([]byte(res[i]), &v)
		if err != nil {
			fmt.Println("解析json失败")
			panic(err)
		}
		result = append(result, v)
	}

	return result
}
