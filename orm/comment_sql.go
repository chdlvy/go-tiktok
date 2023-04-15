package orm

import (
	"strconv"
	"time"

	u "tiktok/utils"

	"gorm.io/gorm"
)

type Comment struct {
	CommentId uint64 `gorm:"comment_id"`
	UserId    uint64 `gorm:"user_id"`
	VideoId   uint64 `gorm:"video_id"`
	Content   string `gorm:"content"`
	LikeCount uint64 `gorm:"like_count"`
	CreateAt  string `gorm:"create_at"`
}

func AddCommentToSql(commentId uint64, userId uint64, videoId uint64, content string) error {
	db := GetSqlConn()
	createAt := strconv.Itoa(int(time.Now().Unix()))
	comment := Comment{CommentId: commentId, UserId: userId, VideoId: videoId, Content: content, LikeCount: 0, CreateAt: createAt}
	err := db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCommentSql(videoId uint64, page uint64) ([]map[string]interface{}, error) {
	db := GetSqlConn()
	c := []Comment{}
	err := db.Offset(int(page)*20).Limit(20).Where("video_id = ?", videoId).Order("like_count desc").Find(&c).Error
	if err != nil {
		return nil, err
	}
	var arr []map[string]interface{}
	for _, v := range c {
		arr = append(arr, u.StructToMap(v))
	}
	return arr, nil

}

func LikeCommentSql(commentId uint64) error {
	db := GetSqlConn()
	err := db.Model(&Comment{}).Where("comment_id = ?", commentId).UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func CancelLikeCommentSql(commentId uint64) error {
	db := GetSqlConn()
	err := db.Model(&Comment{}).Where("comment_id = ?", commentId).UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}
