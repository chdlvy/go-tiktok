package redis

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 根据电话获取用户id
func GetUserIdByPhone(phone string) uint64 {
	rdb := GetredisConn()
	defer rdb.Close()
	res, _ := redis.Uint64(rdb.Do("hget", IDPHONE, phone))
	if res > 0 {
		return res
	} else {
		return 0
	}
}

// 检查账号密码判断是否登录成功
func CheckUserLogin(phone string, password string) (bool, error) {
	uid := GetUserIdByPhone(phone)
	if uid == 0 {
		return false, errors.New("账号未注册")
	}
	rdb := GetredisConn()
	defer rdb.Close()
	table := fmt.Sprintf("user:%d:detail", uid)
	psw, _ := redis.String(rdb.Do("hget", table, "password"))
	if psw == password {
		return true, nil
	}
	return false, errors.New("账号或者密码错误")

}

func HasRejistered(phone string) error {
	res := GetUserIdByPhone(phone)
	if res > 0 {
		return errors.New("账号已经被注册")
	}
	return nil
}
