package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type ManageRedis struct {
}

const (
	HOST = "127.0.0.1:6379"
)

var pool *redis.Pool

func GetredisConn() redis.Conn {

	if pool == nil {
		pool = &redis.Pool{
			MaxIdle:     16,
			MaxActive:   0,
			IdleTimeout: 300,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", HOST)
			},
		}
	}
	// 从连接池中拿一个连接出去
	return pool.Get()
}

// 每一个用户注册就添加user到user表中
func AddUsertoRedis() {

}

// 获取当前最大的用户ID
func GetMaxUserId() (int64, error) {
	rdb := GetredisConn()
	defer rdb.Close()

	// 获取user数量，如果为空，则设置为0
	userId, err := redis.Int64(rdb.Do("get", "user:max_id"))
	if err != nil && err != redis.ErrNil {
		return 0, err
	}
	if err == redis.ErrNil {
		userId = 0
	}
	return userId, nil
}

func IncrMaxUserId() (int64, error) {
	rdb := GetredisConn()
	defer rdb.Close()

	// 更新最大ID，将新用户ID设置为当前最大ID加1
	num, err := redis.Int64(rdb.Do("incr", "user:max_id"))
	if err != nil {
		return 0, err
	}
	return num, nil
}

func ExistUser(userId uint64) bool {
	rdb := GetredisConn()
	defer rdb.Close()
	key := fmt.Sprintf("user:%d:detail", userId)
	res, _ := redis.Int64(rdb.Do("exists", key))
	return res == 1
}

func AddModUserToRedis(userId uint64, userName string, password string, avatar string, gender uint8, birthday string, publishCount uint64) error {
	rdb := GetredisConn()
	defer rdb.Close()

	// 创建一个新的Hash存储用户信息
	key := fmt.Sprintf("user:%d:detail", userId)
	_, err := rdb.Do("hmset", key, "id", userId, "username", userName, "password",
		password, "avatar", avatar, "gender", gender,
		"birthday", birthday, "publish_count", publishCount,
		"create_time", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func GetUser(userId uint64) (map[string]string, error) {
	rdb := GetredisConn()
	defer rdb.Close()

	ok := ExistUser(userId)
	if !ok {
		panic("用户不存在")
	}

	key := fmt.Sprintf("user:%d:detail", userId)
	res, err := redis.StringMap(rdb.Do("hgetall", key))
	return res, err
	// redis.ScanStruct(rdb.Do("hgetall", key))
}

func SaveVideoToRedis(uid uint64, v string) {
	rdb := GetredisConn()
	defer rdb.Close()
	// 获取user
	user, _ := GetUser(uid)

	videoId, _ := strconv.ParseUint(user["publish_count"], 10, 64)
	k := fmt.Sprintf("user:%d:video:%d", uid, videoId+1)
	userk := fmt.Sprintf("user:%d:detail", uid)
	// k := "user:" + strconv.FormatUint(uid, 10) + ":video:" + videoId

	rdb.Send("multi")
	rdb.Send("hset", "video", k, v)
	// 作者的发布数量自增1
	rdb.Send("hincrby", userk, "publish_count", "1")

	// 执行事务
	res, err := rdb.Do("exec")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
