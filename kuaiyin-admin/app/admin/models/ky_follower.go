package models

import (

	"go-admin/common/models"

)

type KyFollower struct {
    models.Model
    
    UserId int64 `json:"userId" gorm:"type:int(20);comment:用户id"` 
    FollowerId int64 `json:"followerId" gorm:"type:int(20);comment:关注id"` 
    models.ModelTime
    models.ControlBy
}

func (KyFollower) TableName() string {
    return "ky_follower"
}

func (e *KyFollower) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *KyFollower) GetId() interface{} {
	return e.Id
}