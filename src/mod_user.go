package src

import (
	"time"
)

type ModUser struct {
	Id           uint64    `redis:"id"`
	UserName     string    `redis:"username" form:"userName"`
	Password     string    `redis:"password" form:"password"`
	Avatar       string    `redis:"avatar" form:"avatar"`
	Gender       uint8     `redis:"gender" form:"gender"`
	Birthday     string    `redis:"birthday" form:"birthday"`
	PublishCount uint64    `redis:"publish_count" form:"publishCount"`
	CreateTime   time.Time `redis:"create_time" form:"createTime"`
}

func (this *ModUser) SetUserName(name string) {
	this.UserName = name
}
