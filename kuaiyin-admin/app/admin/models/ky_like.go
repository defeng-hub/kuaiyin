package models

import (

	"go-admin/common/models"

)

type KyLike struct {
    models.Model
    
    UserId int64 `json:"userId" gorm:"type:bigint(20);comment:用户id"` 
    VideoId int64 `json:"videoId" gorm:"type:bigint(20);comment:视频id"` 
    models.ModelTime
    models.ControlBy
}

func (KyLike) TableName() string {
    return "ky_like"
}

func (e *KyLike) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *KyLike) GetId() interface{} {
	return e.Id
}