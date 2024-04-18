package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type VideosGetPageReq struct {
	dto.Pagination     `search:"-"`
    AuthorId int64 `form:"authorId"  search:"type:exact;column:author_id;table:videos" comment:"作者id"`
    Title string `form:"title"  search:"type:contains;column:title;table:videos" comment:"视频标题"`
    VideosOrder
}

type VideosOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:videos"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:videos"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:videos"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:videos"`
    AuthorId string `form:"authorIdOrder"  search:"type:order;column:author_id;table:videos"`
    PlayUrl string `form:"playUrlOrder"  search:"type:order;column:play_url;table:videos"`
    CoverUrl string `form:"coverUrlOrder"  search:"type:order;column:cover_url;table:videos"`
    FavoriteCount string `form:"favoriteCountOrder"  search:"type:order;column:favorite_count;table:videos"`
    CommentCount string `form:"commentCountOrder"  search:"type:order;column:comment_count;table:videos"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:videos"`
    
}

func (m *VideosGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type VideosInsertReq struct {
    Id int `json:"-" comment:""` // 
    AuthorId int64 `json:"authorId" comment:"作者id"`
    PlayUrl string `json:"playUrl" comment:"视频地址"`
    CoverUrl string `json:"coverUrl" comment:"图片地址"`
    FavoriteCount int64 `json:"favoriteCount" comment:"点赞数量"`
    CommentCount int64 `json:"commentCount" comment:"评论数量"`
    Title string `json:"title" comment:"视频标题"`
    common.ControlBy
}

func (s *VideosInsertReq) Generate(model *models.Videos)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.AuthorId = s.AuthorId
    model.PlayUrl = s.PlayUrl
    model.CoverUrl = s.CoverUrl
    model.FavoriteCount = s.FavoriteCount
    model.CommentCount = s.CommentCount
    model.Title = s.Title
}

func (s *VideosInsertReq) GetId() interface{} {
	return s.Id
}

type VideosUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    AuthorId int64 `json:"authorId" comment:"作者id"`
    PlayUrl string `json:"playUrl" comment:"视频地址"`
    CoverUrl string `json:"coverUrl" comment:"图片地址"`
    FavoriteCount int64 `json:"favoriteCount" comment:"点赞数量"`
    CommentCount int64 `json:"commentCount" comment:"评论数量"`
    Title string `json:"title" comment:"视频标题"`
    common.ControlBy
}

func (s *VideosUpdateReq) Generate(model *models.Videos)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.AuthorId = s.AuthorId
    model.PlayUrl = s.PlayUrl
    model.CoverUrl = s.CoverUrl
    model.FavoriteCount = s.FavoriteCount
    model.CommentCount = s.CommentCount
    model.Title = s.Title
}

func (s *VideosUpdateReq) GetId() interface{} {
	return s.Id
}

// VideosGetReq 功能获取请求参数
type VideosGetReq struct {
     Id int `uri:"id"`
}
func (s *VideosGetReq) GetId() interface{} {
	return s.Id
}

// VideosDeleteReq 功能删除请求参数
type VideosDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *VideosDeleteReq) GetId() interface{} {
	return s.Ids
}
