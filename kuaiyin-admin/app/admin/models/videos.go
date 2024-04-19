package models

import (
	"go-admin/common/models"
)

type Videos struct {
	models.Model

	AuthorId      int64  `json:"authorId" gorm:"type:bigint unsigned;comment:作者id"`
	PlayUrl       string `json:"playUrl" gorm:"type:varchar(512);comment:视频地址"`
	CoverUrl      string `json:"coverUrl" gorm:"type:varchar(255);comment:图片地址"`
	FavoriteCount int64  `json:"favoriteCount" gorm:"type:bigint;comment:点赞数量"`
	CommentCount  int64  `json:"commentCount" gorm:"type:bigint;comment:评论数量"`
	Title         string `json:"title" gorm:"type:varchar(50);comment:视频标题"`
	models.ModelTime
	//models.ControlBy
}

func (Videos) TableName() string {
	return "videos"
}

//func (e *Videos) Generate() models.ActiveRecord {
//	o := *e
//	return &o
//}

func (e *Videos) GetId() interface{} {
	return e.Id
}
