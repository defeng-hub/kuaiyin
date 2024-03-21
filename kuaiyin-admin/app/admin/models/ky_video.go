package models

import (

	"go-admin/common/models"

)

type KyVideo struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(128);comment:标题"` 
    Src string `json:"src" gorm:"type:varchar(500);comment:视频地址"` 
    UserId int64 `json:"userId" gorm:"type:int(10);comment:用户id"` 
    Username string `json:"username" gorm:"type:varchar(255);comment:用户名"` 
    Href string `json:"href" gorm:"type:varchar(255);comment:图片"` 
    Msg string `json:"msg" gorm:"type:varchar(500);comment:描述"` 
    Link int64 `json:"link" gorm:"type:int(10);comment:喜欢"` 
    Sms int64 `json:"sms" gorm:"type:int(10);comment:评论数量"` 
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