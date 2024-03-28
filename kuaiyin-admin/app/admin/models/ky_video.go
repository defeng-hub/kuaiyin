package models

import (

	"go-admin/common/models"

)

type KyVideo struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(255);comment:视频标题"` 
    Msg string `json:"msg" gorm:"type:varchar(255);comment:视频信息"` 
    Like int64 `json:"like" gorm:"type:varchar(255);comment:喜欢"` 
    Sms int64 `json:"sms" gorm:"type:varchar(255);comment:评论数"` 
    Src string `json:"src" gorm:"type:varchar(255);comment:视频地址"` 
    Href string `json:"href" gorm:"type:varchar(255);comment:图片"` 
    UserId int64 `json:"userId" gorm:"type:int(20);comment:快音号"` 
    models.ModelTime
    models.ControlBy
}

func (KyVideo) TableName() string {
    return "ky_video"
}

func (e *KyVideo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *KyVideo) GetId() interface{} {
	return e.Id
}