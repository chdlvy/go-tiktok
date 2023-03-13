package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type ManageRedis struct {
}

func (this *ManageRedis) InitRedis(host string) {
	conn, err := redis.Dial("tcp", host)
	if err != nil {
		fmt.Println("connect redis err：", err)
		return
	}
	defer conn.Close()

}
