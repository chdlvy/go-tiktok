package orm

import "time"

type Video struct {
	Id           uint64    `gorm:"id;primaryKey"` //视频id
	WriterId     uint64    `gorm:"writerId"`      //作者id
	VideoSrc     string    `gorm:"videoSrc"`      //视频链接
	Title        string    `gorm:"title" `        //标题
	Description  string    `gorm:"description" `  //描述
	Cover        string    `gorm:"cover" `        //封面
	Duration     time.Time `gorm:"duration" `     //时长
	PlayCount    uint64    `gorm:"playCount" `    //播放量
	LikeCount    uint64    `gorm:"likeCount" `    //点赞数
	CollectCount uint64    `gorm:"collectCount" ` //收藏数
	ShareCount   uint64    `gorm:"shareCount" `   //分享数
	CommentCount uint64    `gorm:"commentCount" ` //评论数
	IsLike       uint8     `gorm:"isLike"`        //检查用户是否点赞过
}

func AddVideoToSql(v map[string]interface{}) error {
	db := GetSqlConn()
	video := Video{
		WriterId:    v["writerId"].(uint64),
		VideoSrc:    v["videoSrc"].(string),
		Title:       v["title"].(string),
		Description: v["description"].(string),
		Cover:       v["cover"].(string),
		Duration:    v["duration"].(time.Time),
	}
	err := db.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
