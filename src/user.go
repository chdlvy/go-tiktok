package src

import (
	"fmt"
	"strconv"
	redis "tiktok/internal/pkg/redis"
	"time"

	redisgo "github.com/garyburd/redigo/redis"
)

type User struct {
	ModComment *ModComment
	ModReply   *ModReply
	ModUser    *ModUser
	ModVideo   *ModVideo
}

func InitUser(newUser *ModUser) *User {
	user := new(User)
	user.ModComment = new(ModComment)
	user.ModReply = new(ModReply)
	user.ModUser = newUser
	user.AddUser()
	user.ModVideo = new(ModVideo)
	return user
}

func (this *User) AddUser() error {

	// 更新最大ID，将新用户ID设置为当前最大ID加1
	num, err := redis.IncrMaxUserId()
	if err != nil {
		return err
	}
	this.ModUser.Id = uint64(num)

	// rdb := redis.GetredisConn()
	// defer rdb.Close()
	// key := fmt.Sprintf("user:%d:", this.ModUser.Id)
	// str, _ := json.Marshal(this)
	// fmt.Println(str)
	// rdb.Do("hmset", "user", key, string(str))

	// 将用户信息存储到Redis
	err = redis.AddModUserToRedis(this.ModUser.Id, this.ModUser.UserName, this.ModUser.Password, this.ModUser.Avatar, this.ModUser.Gender, this.ModUser.Birthday, this.ModUser.PublishCount)
	if err != nil {
		return err
	}

	return nil
}

// 根据id获取用户信息
func GetUserById(userId uint64) (*ModUser, error) {
	rdb := redis.GetredisConn()
	defer rdb.Close()

	key := "user:" + strconv.FormatUint(userId, 10)
	// 先检查用户是否存在
	isExists, err := redisgo.Int(rdb.Do("exists", key))
	if isExists == 0 {
		fmt.Println("用户不存在")
		return nil, err
	}

	values, err := redisgo.Strings(rdb.Do("hmget", key, "id", "username", "password", "avatar", "gender", "birthday", "publish_count", "create_time"))
	if err != nil {
		return nil, err
	}

	// 获取的数据对应放入结构体中并返回
	var user ModUser
	user.Id, _ = strconv.ParseUint(values[0], 10, 64)
	user.UserName = values[1]
	user.Password = values[2]
	user.Avatar = values[3]

	gender64, _ := strconv.ParseUint(values[4], 10, 8)
	user.Gender = uint8(gender64)

	user.Birthday = values[5]
	user.PublishCount, _ = strconv.ParseUint(values[6], 10, 64)

	createTimeUnix, _ := strconv.ParseInt(values[7], 10, 64)
	user.CreateTime = time.Unix(createTimeUnix, 0)

	return &user, nil
}
