package src

import "time"

type ModComment struct {
	id         uint64
	UerId      uint64
	videoId    uint64
	Content    string
	createTime time.Time
	Reply      *ModReply
}
