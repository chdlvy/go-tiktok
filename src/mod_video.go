package src

import (
	"time"
)

type ModVideo struct {
	Id           uint64    `json:"id" form:"id"`                     //视频id
	WriterId     uint64    `json:"writerId" form:"writerId"`         //作者id
	VideoSrc     string    `json:"videoSrc" form:"videoSrc"`         //视频链接
	Title        string    `json:"title" form:"title"`               //标题
	Description  string    `json:"description" form:"description"`   //描述
	Cover        string    `json:"cover" form:"cover"`               //封面
	Duration     time.Time `json:"duration" form:"duration"`         //时长
	PlayCount    uint64    `json:"playCount" form:"playCount"`       //播放量
	LikeCount    uint64    `json:"likeCount" form:"likeCount"`       //点赞数
	CollectCount uint64    `json:"collectCount" form:"collectCount"` //收藏数
	ShareCount   uint64    `json:"shareCount" form:"shareCount"`     //分享数
	CommentCount uint64    `json:"commentCount" form:"commentCount"` //评论数
	IsLike       bool      `json:"isLike" form:"isLike"`             //检查用户是否点赞过
}

func (this *ModVideo) Publish(payLoad *ModVideo) {
	// str, _ := json.Marshal(payLoad)
}

func (this *ModVideo) delPublish(videoId uint64) {

}
