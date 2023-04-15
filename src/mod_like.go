package src

import "tiktok/internal/pkg/redis"

type ModLike struct {
	LikeList []uint64 `redis:"likeList"` // 点赞的视频列表

	LikedNum uint64 `redis:likedNum` //获赞数量
}

func (this *ModLike) Like(uid uint64, videoId uint64, writerId uint64) {
	redis.AddtoLikeList(uid, videoId, writerId)
}

func (this *ModLike) GetLikeList(uid uint64) {
	redis.GetLikeList(uid)
}
