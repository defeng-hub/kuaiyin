package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type KyCommentGetPageReq struct {
	dto.Pagination     `search:"-"`
    VideoId string `form:"videoId"  search:"type:exact;column:video_id;table:ky_comment" comment:"视频id"`
    Content string `form:"content"  search:"type:contains;column:content;table:ky_comment" comment:"评论内容"`
    State string `form:"state"  search:"type:exact;column:state;table:ky_comment" comment:"评论状态"`
    KyCommentOrder
}

type KyCommentOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ky_comment"`
    VideoId string `form:"videoIdOrder"  search:"type:order;column:video_id;table:ky_comment"`
    Content string `form:"contentOrder"  search:"type:order;column:content;table:ky_comment"`
    State string `form:"stateOrder"  search:"type:order;column:state;table:ky_comment"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ky_comment"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ky_comment"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ky_comment"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ky_comment"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ky_comment"`
    
}

func (m *KyCommentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type KyCommentInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    VideoId string `json:"videoId" comment:"视频id"`
    Content string `json:"content" comment:"评论内容"`
    State string `json:"state" comment:"评论状态"`
    common.ControlBy
}

func (s *KyCommentInsertReq) Generate(model *models.KyComment)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.VideoId = s.VideoId
    model.Content = s.Content
    model.State = s.State
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *KyCommentInsertReq) GetId() interface{} {
	return s.Id
}

type KyCommentUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    VideoId string `json:"videoId" comment:"视频id"`
    Content string `json:"content" comment:"评论内容"`
    State string `json:"state" comment:"评论状态"`
    common.ControlBy
}

func (s *KyCommentUpdateReq) Generate(model *models.KyComment)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.VideoId = s.VideoId
    model.Content = s.Content
    model.State = s.State
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *KyCommentUpdateReq) GetId() interface{} {
	return s.Id
}

// KyCommentGetReq 功能获取请求参数
type KyCommentGetReq struct {
     Id int `uri:"id"`
}
func (s *KyCommentGetReq) GetId() interface{} {
	return s.Id
}

// KyCommentDeleteReq 功能删除请求参数
type KyCommentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *KyCommentDeleteReq) GetId() interface{} {
	return s.Ids
}
