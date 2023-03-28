package src

import (
	"time"
)

type ModVideo struct {
	Id          uint64    `json:"id" form:"id"`                   //视频id
	WriterId    uint64    `json:"writerId" form:"writerId"`       //作者id
	Title       string    `json:"title" form:"title"`             //标题
	Description string    `json:"description" form:"description"` //描述
	Cover       string    `json:"cover" form:"cover"`             //封面
	Duration    time.Time `json:"duration" form:"duration"`       //时长
	PlayCount   uint64    `json:"playCount" form:"playCount"`     //播放量
	LikeCount   uint64    `json:"likeCount" form:"likeCount"`     //点赞数
}

func (this *ModVideo) Publish(payLoad *ModVideo) {
	// str, _ := json.Marshal(payLoad)
}

func (this *ModVideo) delPublish(videoId uint64) {

}
