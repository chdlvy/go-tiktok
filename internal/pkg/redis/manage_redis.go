package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type ManageRedis struct {
}

const (
	HOST = "127.0.0.1:6379"
	// userid:phone的映射
	IDPHONE = "phone:uid"
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

func InitRedis() {
	rdb := GetredisConn()
	defer rdb.Close()
	rdb.Send("multi")
	rdb.Send("flushall")
	rdb.Send("set", "user:max_id", "0")
	rdb.Send("select", "1")
	rdb.Send("set", "video:max_id", "0")
	rdb.Send("set", "comment:max_id", "0")
	rdb.Send("select", "0")
	rdb.Send("exec")
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

// 用户是否存在
func ExistUser(userId uint64) bool {
	rdb := GetredisConn()
	defer rdb.Close()
	key := fmt.Sprintf("user:%d:detail", userId)
	res, _ := redis.Int64(rdb.Do("exists", key))
	return res == 1
}

// 添加user信息到redis
func AddModUserToRedis(userId uint64, userName string, password string, phone string, avatar string, gender uint8, birthday string, publishCount uint64) error {
	rdb := GetredisConn()
	defer rdb.Close()

	// 创建一个新的Hash存储用户信息
	key := fmt.Sprintf("user:%d:detail", userId)
	rdb.Send("multi")
	rdb.Send("hmset", key, "id", userId, "username", userName, "password",
		password, "phone", phone, "avatar", avatar, "gender", gender,
		"birthday", birthday, "publish_count", publishCount,
		"create_time", time.Now().Unix())
	rdb.Send("hmset", IDPHONE, phone, userId)
	_, err := rdb.Do("exec")
	if err != nil {
		return err
	}
	return nil
}

// 获取user信息
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
