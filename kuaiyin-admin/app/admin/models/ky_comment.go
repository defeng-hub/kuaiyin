package models

import (

	"go-admin/common/models"

)

type KyComment struct {
    models.Model
    
    VideoId int64 `json:"videoId" gorm:"type:int(20);comment:视频id"` 
    UserId int64 `json:"userId" gorm:"type:int(11);comment:用户id"` 
    Msg string `json:"msg" gorm:"type:varchar(128);comment:评论内容"` 
    State int64 `json:"state" gorm:"type:int(255);comment:评论状态（01正常 2有风险）"` 
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