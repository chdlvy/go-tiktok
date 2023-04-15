package redis

import (
	"errors"
	"fmt"
	"sync"

	sql "tiktok/orm"

	"github.com/garyburd/redigo/redis"
)

func AddComment(writerId uint64, videoId uint64, uid uint64, content string) error {

	// 检查用户是否存在
	hasUser := ExistUser(uid)
	if !hasUser {
		return errors.New("用户不存在")
	}

	rdb := GetredisConn()
	defer rdb.Close()
	commentKey := fmt.Sprintf("comment:video:%d", videoId)

	rdb.Do("select", 1)
	// 拿到评论id

	commenId, _ := redis.Uint64(rdb.Do("incr", "comment:max_id"))
	commentcol := fmt.Sprintf("user:%d:comment:%d", uid, commenId)
	video := fmt.Sprintf("user:%d:video:%d", writerId, videoId)
	// 添加评论
	rdb.Do("zadd", commentKey, 0, commentcol)
	rdb.Do("hincrby", video, "commentCount", "1")
	rdb.Do("select", 0)

	// 添加评论到数据库，评论放redis太占内存
	err := sql.AddCommentToSql(commenId, uid, videoId, content)
	if err != nil {
		return err
	}
	return nil
}

func GetComment(videoId uint64, page uint64) ([]map[string]interface{}, uint64) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var c []map[string]interface{}
	// 获取评论信息
	go func() {
		c, _ = sql.GetCommentSql(videoId, page)
		for _, v := range c {
			uid := v["userId"]
			userMsg, _ := GetUser(uid.(uint64))
			um := map[string]interface{}{
				"avatar":   userMsg["avatar"],
				"username": userMsg["username"],
				"id":       userMsg["id"],
			}
			v["userMsg"] = um
		}
		wg.Done()
	}()

	wg.Add(1)
	tmpcount := make(chan uint64, 1)
	// 开个协程从redis获取评论个数
	go func() {
		defer wg.Done()
		rdb := GetredisConn()
		defer rdb.Close()
		commentKey := fmt.Sprintf("comment:video:%d", videoId)
		rdb.Do("select", 1)
		count, _ := redis.Uint64(rdb.Do("zcard", commentKey))
		tmpcount <- count
		rdb.Do("select", 0)
	}()
	wg.Wait()
	return c, <-tmpcount
}

// 点赞评论=============点赞也要操作数据库
func LikeComment(videoId uint64, uid uint64, commentId uint64) error {
	// 检查用户是否存在
	hasUser := ExistUser(uid)
	if !hasUser {
		return errors.New("用户不存在")
	}

	rdb := GetredisConn()
	defer rdb.Close()
	commentKey := fmt.Sprintf("comment:video:%d", videoId)
	commentcol := fmt.Sprintf("user:%d:comment:%d", uid, videoId)
	rdb.Send("multi")
	rdb.Send("select", 1)
	rdb.Send("zincrby", commentKey, 1, commentcol)
	rdb.Send("select", 0)
	_, err := rdb.Do("exec")
	if err != nil {
		return err
	}
	// 调用函数让数据库中的评论点赞数+1
	err = sql.LikeCommentSql(commentId)
	if err != nil {
		return err
	}
	return nil
}

func CancelLikeComment(commentId uint64) error {
	err := sql.CancelLikeCommentSql(commentId)
	if err != nil {
		return err
	}
	return nil
}
