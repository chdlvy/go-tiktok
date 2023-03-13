package src

import "time"

type ModVideo struct {
	id          uint64    //视频id
	writerId    uint64    //作者id
	Title       string    //标题
	Description string    //描述
	Cover       string    //封面
	Duration    time.Time //时长
	PlayCount   uint64    //播放量
	LikeCount   uint64    //点赞数
}

func (this *ModVideo) AddVideo() {

}
