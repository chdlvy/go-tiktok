package redis

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 点赞视频添加到点赞表中
func AddtoLikeList(uid uint64, videoId uint64, writerId uint64) (int64, error) {
	// 检查用户是否存在
	hasUser := ExistUser(uid)
	if !hasUser {
		return 0, errors.New("用户不存在")
	}

	rdb := GetredisConn()
	defer rdb.Close()
	likeListKey := fmt.Sprintf("user:%d:likeList", uid)
	key := fmt.Sprintf("user:%d:liked", writerId)
	likedVideo := fmt.Sprintf("user:%d:video:%d", writerId, videoId)
	isLike := fmt.Sprintf("user:%d:like:video:%d", uid, videoId)
	// 检查该视频是否存在
	isExist := IsExistvideo(likedVideo)
	if !isExist {
		fmt.Println("视频不存在")
		return 0, nil
	}

	// 添加zset集合中的对应视频的点赞数，按照点赞多的排序
	rdb.Send("multi")
	rdb.Send("select", 1)
	rdb.Send("zincrby", "videos", 1, likedVideo)
	// 视频点赞数+1
	rdb.Send("hincrby", likedVideo, "likeCount", "1")

	// 在2号数据库添加用户和视频的点赞关系
	rdb.Send("select", 2)
	rdb.Send("set", isLike, 1)

	rdb.Send("select", 0)
	// 视频作者的获赞数+1
	rdb.Send("hincrby", key, "likedNum", "1")

	// 将点赞的视频的id添加到用户喜欢列表中
	rdb.Send("lpush", likeListKey, likedVideo)
	val, err := redis.Values(rdb.Do("exec"))
	if err != nil {
		return 0, err
	}
	// 返回点赞数
	return val[2].(int64), nil

}

func CancelLike(uid uint64, videoId uint64, writerId uint64) (int64, error) {
	// 检查用户是否存在
	hasUser := ExistUser(uid)
	if !hasUser {
		return 0, errors.New("用户不存在")
	}

	rdb := GetredisConn()
	defer rdb.Close()
	likeListKey := fmt.Sprintf("user:%d:likeList", uid)
	key := fmt.Sprintf("user:%d:liked", writerId)
	likedVideo := fmt.Sprintf("user:%d:video:%d", writerId, videoId)
	isLike := fmt.Sprintf("user:%d:like:video:%d", uid, videoId)
	// 检查该视频是否存在
	isExist := IsExistvideo(likedVideo)
	if !isExist {
		fmt.Println("视频不存在")
		return 0, nil
	}

	// 减少zset集合中的对应视频的点赞数
	rdb.Send("multi")
	rdb.Send("select", 1)
	rdb.Send("zincrby", "videos", -1, likedVideo)
	// 视频点赞数-1
	rdb.Send("hincrby", likedVideo, "likeCount", -1)

	// 取消用户和视频的点赞关系
	rdb.Send("select", 2)
	rdb.Send("del", isLike)

	rdb.Send("select", 0)
	// 视频作者的获赞数-1
	rdb.Send("hincrby", key, "likedNum", -1)

	// 将点赞的视频的id添加到用户喜欢列表中
	rdb.Send("lrem", likeListKey, 1, likedVideo)
	val, err := redis.Values(rdb.Do("exec"))
	if err != nil {
		return 0, err
	}
	// 返回点赞数
	return val[2].(int64), nil
}

func GetLikeList(uid uint64) []map[string]string {
	rdb := GetredisConn()
	defer rdb.Close()
	likeListKey := fmt.Sprintf("user:%d:likeList", uid)

	res, err := redis.Strings(rdb.Do("lrange", likeListKey, 0, -1))
	if err != nil {
		panic(err)
	}

	rdb.Do("select", 1)

	// 获取视频
	var videos []map[string]string
	for _, result := range res {
		// 获取视频详细信息
		fields, err := redis.Strings(rdb.Do("HGETALL", result))
		if err != nil {
			fmt.Println(err)
			continue
		}

		video := make(map[string]string)
		for i := 0; i < len(fields); i += 2 {
			video[fields[i]] = fields[i+1]
		}
		videos = append(videos, video)
	}
	rdb.Do("select", 0)
	return videos
}
