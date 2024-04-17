package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type KyLikeGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId int64 `form:"userId"  search:"type:exact;column:user_id;table:ky_like" comment:"用户id"`
    VideoId int64 `form:"videoId"  search:"type:exact;column:video_id;table:ky_like" comment:"视频id"`
    KyLikeOrder
}

type KyLikeOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ky_like"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:ky_like"`
    VideoId string `form:"videoIdOrder"  search:"type:order;column:video_id;table:ky_like"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ky_like"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ky_like"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ky_like"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ky_like"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ky_like"`
    
}

func (m *KyLikeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type KyLikeInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户id"`
    VideoId int64 `json:"videoId" comment:"视频id"`
    common.ControlBy
}

func (s *KyLikeInsertReq) Generate(model *models.KyLike)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.VideoId = s.VideoId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *KyLikeInsertReq) GetId() interface{} {
	return s.Id
}

type KyLikeUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户id"`
    VideoId int64 `json:"videoId" comment:"视频id"`
    common.ControlBy
}

func (s *KyLikeUpdateReq) Generate(model *models.KyLike)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.VideoId = s.VideoId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *KyLikeUpdateReq) GetId() interface{} {
	return s.Id
}

// KyLikeGetReq 功能获取请求参数
type KyLikeGetReq struct {
     Id int `uri:"id"`
}
func (s *KyLikeGetReq) GetId() interface{} {
	return s.Id
}

// KyLikeDeleteReq 功能删除请求参数
type KyLikeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *KyLikeDeleteReq) GetId() interface{} {
	return s.Ids
}
