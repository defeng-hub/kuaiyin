package models

import (
	"time"

	"go-admin/common/models"
)

type Comments struct {
	models.Model

	CreatedTime time.Time `json:"createdTime" gorm:"type:datetime(3);comment:创建时间"`
	VideoId     string    `json:"videoId" gorm:"type:bigint;comment:视频id"`
	UserId      string    `json:"userId" gorm:"type:bigint;comment:用户id"`
	Content     string    `json:"content" gorm:"type:varchar(255);comment:评论内容"`
	//models.ModelTime
	//models.ControlBy
}

func (Comments) TableName() string {
	return "comments"
}

//func (e *Comments) Generate() models.ActiveRecord {
//	o := *e
//	return &o
//}

func (e *Comments) GetId() interface{} {
	return e.Id
}
