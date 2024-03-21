package models

import (

	"go-admin/common/models"

)

type KyComment struct {
    models.Model
    
    VideoId string `json:"videoId" gorm:"type:int(10);comment:视频id"` 
    Content string `json:"content" gorm:"type:varchar(128);comment:评论内容"` 
    State string `json:"state" gorm:"type:int(10);comment:评论状态"` 
    models.ModelTime
    models.ControlBy
}

func (KyComment) TableName() string {
    return "ky_comment"
}

func (e *KyComment) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *KyComment) GetId() interface{} {
	return e.Id
}