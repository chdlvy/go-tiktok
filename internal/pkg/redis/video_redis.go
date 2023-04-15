package redis

import (
	"fmt"
	"strconv"

	o "tiktok/orm"

	"github.com/garyburd/redigo/redis"
)

// 判断视频是否存在
func IsExistvideo(key string) bool {
	rdb := GetredisConn()
	defer rdb.Close()
	rdb.Send("multi")
	rdb.Send("select", 1)
	rdb.Send("exists", key)
	rdb.Send("select", 0)
	res, err := redis.Values(rdb.Do("exec"))
	if err != nil {
		fmt.Println(err)
		return false
	}
	if exist, _ := redis.Int(res[1], nil); exist == 0 {
		return false
	} else {
		return true
	}
}

// 保存视频到redis中
func SaveVideoToRedis(uid uint64, v map[string]interface{}) {
	rdb := GetredisConn()
	defer rdb.Close()
	// videoId, _ := strconv.ParseUint(user["publish_count"], 10, 64)
	userk := fmt.Sprintf("user:%d:detail", uid)

	// 获取视频总数
	rdb.Send("multi")
	rdb.Send("select", "1")
	rdb.Send("incr", "video:max_id")
	rdb.Send("get", "video:max_id")
	rdb.Send("select", "0")
	res1, err := redis.Values(rdb.Do("exec"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取video的最大数量
	maxId, _ := redis.String(res1[2], nil)
	m := fmt.Sprintf("user:%d:video:%s", uid, maxId)

	rdb.Send("multi")

	// 作者的发布数量自增1
	rdb.Send("hincrby", userk, "publish_count", "1")

	// 将视频信息保存到数据库1中
	rdb.Send("select", "1")

	// 将视频信息保存到有序集合中，分值为点赞数，默认为0
	rdb.Send("ZADD", "videos", 0, m)

	rdb.Send("hmset", m, "id", maxId, "writerId", v["writerId"],
		"videoSrc", v["videoSrc"], "collectCount", v["collectCount"],
		"shareCount", v["shareCount"], "commentCount", v["commentCount"],
		"title", v["title"], "description", v["description"],
		"cover", v["cover"], "duration", v["duration"], "isLike", v["isLike"],
		"playCount", v["playCount"], "likeCount", v["likeCount"])

	// 回到0数据库
	rdb.Send("select", "0")
	// 执行事务
	_, err = rdb.Do("exec")
	if err != nil {
		fmt.Println(err)
		return
	}
	//持久化到数据库
	go func() {
		err := o.AddVideoToSql(v)
		if err != nil {
			panic(err)
		}
	}()

}
func CheckIsLike(uid uint64, videoId uint64) int {
	rdb := GetredisConn()
	defer rdb.Close()
	isLikeKey := fmt.Sprintf("user:%d:like:video:%d", uid, videoId)
	rdb.Do("select", 2)
	// 点赞了为true，反之false
	res, _ := redis.Int(rdb.Do("get", isLikeKey))
	rdb.Do("select", 0)
	return res
}

func GetVideoList(uid uint64, start int64, end int64) []map[string]string {
	rdb := GetredisConn()
	defer rdb.Close()
	rdb.Do("select", 1)
	// 获取前十条视频数据
	results, err := redis.Strings(rdb.Do("ZREVRANGE", "videos", start, end))
	if err != nil {
		panic(err)
	}

	var videos []map[string]string
	for _, result := range results {
		// 获取视频详细信息
		fields, err := redis.Strings(rdb.Do("HGETALL", result))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
		fmt.Println(fields)
		video := make(map[string]string)
		for i := 0; i < len(fields); i += 2 {
			video[fields[i]] = fields[i+1]
		}
		//向map中添加一些作者信息
		writerId, _ := strconv.ParseUint(video["writerId"], 10, 64)
		writerMsg, _ := GetUser(writerId)
		video["writerAvatar"] = writerMsg["avatar"]
		video["writerName"] = writerMsg["username"]
		videoId, _ := strconv.ParseUint(video["id"], 10, 64)
		// 检查用户和视频的点赞关系
		isLike := CheckIsLike(uid, videoId)
		video["isLike"] = strconv.Itoa(isLike)

		videos = append(videos, video)
	}
	rdb.Do("select", 0)
	return videos
}

// ========================================================
// likeStatus   uid:vid :
// func GetVideoWithLikeStatus(videoId int, uid int) {
// 	rdb := GetredisConn()
// 	defer rdb.Close()

// 	key := fmt.Sprintf("user:%d:liked", uid)
// 	likedVideo := fmt.Sprintf("user:%d:video:%d", uid, videoId)
// 	liked, err := redis.Int(rdb.Do("hget", key, likedVideo))
// }
