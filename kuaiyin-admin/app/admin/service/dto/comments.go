package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CommentsGetPageReq struct {
	dto.Pagination     `search:"-"`
    VideoId string `form:"videoId"  search:"type:exact;column:video_id;table:comments" comment:"视频id"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:comments" comment:"用户id"`
    Content string `form:"content"  search:"type:contains;column:content;table:comments" comment:"评论内容"`
    CommentsOrder
}

type CommentsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:comments"`
    CreatedTime string `form:"createdTimeOrder"  search:"type:order;column:created_time;table:comments"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:comments"`
    VideoId string `form:"videoIdOrder"  search:"type:order;column:video_id;table:comments"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:comments"`
    Content string `form:"contentOrder"  search:"type:order;column:content;table:comments"`
    
}

func (m *CommentsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CommentsInsertReq struct {
    Id int `json:"-" comment:""` // 
    CreatedTime time.Time `json:"createdTime" comment:"创建时间"`
    VideoId string `json:"videoId" comment:"视频id"`
    UserId string `json:"userId" comment:"用户id"`
    Content string `json:"content" comment:"评论内容"`
    common.ControlBy
}

func (s *CommentsInsertReq) Generate(model *models.Comments)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreatedTime = s.CreatedTime
    model.VideoId = s.VideoId
    model.UserId = s.UserId
    model.Content = s.Content
}

func (s *CommentsInsertReq) GetId() interface{} {
	return s.Id
}

type CommentsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    CreatedTime time.Time `json:"createdTime" comment:"创建时间"`
    VideoId string `json:"videoId" comment:"视频id"`
    UserId string `json:"userId" comment:"用户id"`
    Content string `json:"content" comment:"评论内容"`
    common.ControlBy
}

func (s *CommentsUpdateReq) Generate(model *models.Comments)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreatedTime = s.CreatedTime
    model.VideoId = s.VideoId
    model.UserId = s.UserId
    model.Content = s.Content
}

func (s *CommentsUpdateReq) GetId() interface{} {
	return s.Id
}

// CommentsGetReq 功能获取请求参数
type CommentsGetReq struct {
     Id int `uri:"id"`
}
func (s *CommentsGetReq) GetId() interface{} {
	return s.Id
}

// CommentsDeleteReq 功能删除请求参数
type CommentsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CommentsDeleteReq) GetId() interface{} {
	return s.Ids
}
