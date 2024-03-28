package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type KyFollowerGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId int64 `form:"userId"  search:"type:exact;column:user_id;table:ky_follower" comment:"用户id"`
    FollowerId int64 `form:"followerId"  search:"type:exact;column:follower_id;table:ky_follower" comment:"关注id"`
    KyFollowerOrder
}

type KyFollowerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ky_follower"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:ky_follower"`
    FollowerId string `form:"followerIdOrder"  search:"type:order;column:follower_id;table:ky_follower"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ky_follower"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ky_follower"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ky_follower"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ky_follower"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ky_follower"`
    
}

func (m *KyFollowerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type KyFollowerInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户id"`
    FollowerId int64 `json:"followerId" comment:"关注id"`
    common.ControlBy
}

func (s *KyFollowerInsertReq) Generate(model *models.KyFollower)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.FollowerId = s.FollowerId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *KyFollowerInsertReq) GetId() interface{} {
	return s.Id
}

type KyFollowerUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户id"`
    FollowerId int64 `json:"followerId" comment:"关注id"`
    common.ControlBy
}

func (s *KyFollowerUpdateReq) Generate(model *models.KyFollower)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.FollowerId = s.FollowerId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *KyFollowerUpdateReq) GetId() interface{} {
	return s.Id
}

// KyFollowerGetReq 功能获取请求参数
type KyFollowerGetReq struct {
     Id int `uri:"id"`
}
func (s *KyFollowerGetReq) GetId() interface{} {
	return s.Id
}

// KyFollowerDeleteReq 功能删除请求参数
type KyFollowerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *KyFollowerDeleteReq) GetId() interface{} {
	return s.Ids
}
