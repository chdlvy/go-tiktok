package router

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"tiktok/internal/pkg/redis"
	t "tiktok/jwt"
	"tiktok/middleware"
	"tiktok/src"
	utils "tiktok/utils"
)

type Router struct {
}

var userMap map[uint64]*src.User

func (this *Router) InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.POST("/test", this.test)

	r.POST("/login", this.login)
	// r.POST("/sign", this.sign)

	// r.Use(middleware.VerifyTokenMiddleWare())
	r.POST("/publish", this.publish)
	r.GET("/getVideoList", this.getVideoList)
	r.GET("/like", this.like)
	r.GET("/cancelLike", this.cancelLike)
	r.GET("/getLikeList", this.getLikeList)
	r.POST("/upVideo", this.upVideo)
	r.POST("/comment", this.comment)
	r.GET("/likeComment", this.likeComment)
	r.GET("/cancelLikeComment", this.cancelLikeComment)
	r.GET("getComment", this.getComment)
}

func (this *Router) test(c *gin.Context) {
	d := c.PostForm("name")
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": d,
	})
}

func (this *Router) login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	_, err := redis.CheckUserLogin(phone, password)
	if err != nil {
		// 注册账号
		if err.Error() == "账号未注册" {
			var newUser = src.ModUser{
				Phone:        phone,
				UserName:     "用户" + strconv.Itoa(int(time.Now().Unix())),
				Password:     password,
				Avatar:       "defaultImg",
				Gender:       2,
				Birthday:     "",
				PublishCount: 0,
			}
			user, err := src.InitUser(&newUser)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			} else {
				// 注册完毕直接登录
				m := map[string]interface{}{
					"phone":    phone,
					"password": password,
				}
				token := t.GetToken(m)
				res, _ := redis.GetUser(user.ModUser.Id)
				c.JSON(http.StatusOK, gin.H{
					"status": 1,
					"msg":    "登录成功",
					"token":  token,
					"user":   res,
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
	} else {
		// 登录账号并返回token
		m := map[string]interface{}{
			"phone":    phone,
			"password": password,
		}
		token := t.GetToken(m)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "登录成功",
			"token":  token,
		})
	}

}

// func (this *Router) sign(c *gin.Context) {
// 	var newUser src.ModUser
// 	c.ShouldBind(&newUser)
// 	_, err := src.InitUser(&newUser)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": err.Error(),
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": "注册成功",
// 		})
// 	}

// }

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

	// 结构体转map
	value := utils.StructToMap(newVideo)
	redis.SaveVideoToRedis(user.Id, value)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func (this *Router) getVideoList(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	uid, _ := strconv.ParseUint(c.Query("uid"), 10, 64)
	start := page * 10
	end := page*10 + 10

	// list := redis.GetVideo(start, end)
	list := redis.GetVideoList(uid, start, end)

	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}

func (this *Router) like(c *gin.Context) {
	uid, _ := strconv.ParseUint(c.Query("uid"), 10, 64)
	videoId, _ := strconv.ParseUint(c.Query("videoId"), 10, 64)
	writerId, _ := strconv.ParseUint(c.Query("writerId"), 10, 64)
	likeCount, err := redis.AddtoLikeList(uid, videoId, writerId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"likeCount": likeCount,
	})
}

func (this *Router) cancelLike(c *gin.Context) {
	uid, _ := strconv.ParseUint(c.Query("uid"), 10, 64)
	videoId, _ := strconv.ParseUint(c.Query("videoId"), 10, 64)
	writerId, _ := strconv.ParseUint(c.Query("writerId"), 10, 64)
	likeCount, err := redis.CancelLike(uid, videoId, writerId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"likeCount": likeCount,
	})
}

func (this *Router) getLikeList(c *gin.Context) {
	uid, _ := strconv.ParseUint(c.Query("uid"), 10, 64)
	likeList := redis.GetLikeList(uid)
	c.JSON(http.StatusOK, gin.H{
		"likeList": likeList,
	})
}

func (this *Router) upVideo(c *gin.Context) {
	video, _ := c.FormFile("video")
	timer := strconv.Itoa(int(time.Now().Unix()))
	video.Filename = timer
	dst := "../asserts/video/" + video.Filename + ".mp4"
	c.SaveUploadedFile(video, dst)
}

func (this *Router) comment(c *gin.Context) {
	writerId, _ := strconv.ParseUint(c.PostForm("writerId"), 10, 64)
	videoId, _ := strconv.ParseUint(c.PostForm("videoId"), 10, 64)
	uid, _ := strconv.ParseUint(c.PostForm("uid"), 10, 64)
	content := c.PostForm("content")
	// 添加评论
	err := redis.AddComment(writerId, videoId, uid, content)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "评论成功",
	})
}

func (this *Router) likeComment(c *gin.Context) {
	uid, _ := strconv.ParseUint(c.Query("uid"), 10, 64)
	videoId, _ := strconv.ParseUint(c.Query("videoId"), 10, 64)
	commentId, _ := strconv.ParseUint(c.Query("commentId"), 10, 64)
	err := redis.LikeComment(videoId, uid, commentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"likeStatus": 0,
			"err":        err.Error(),
		})
	}
	// 0表示点赞失败，1表示成功
	c.JSON(http.StatusOK, gin.H{
		"likeStatus": 1,
		"msg":        "点赞成功",
	})
}

func (this *Router) cancelLikeComment(c *gin.Context) {
	commentId, _ := strconv.ParseUint(c.Query("commentId"), 10, 64)
	err := redis.CancelLikeComment(commentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"cancelLikeStatus": 0,
			"err":              err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"cancelLikeStatus": 1,
	})
}

func (this *Router) getComment(c *gin.Context) {
	page, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	videoId, _ := strconv.ParseUint(c.Query("videoId"), 10, 64)
	commentList, count := redis.GetComment(videoId, page)
	c.JSON(http.StatusOK, gin.H{
		"list":  commentList,
		"count": count,
	})
}
