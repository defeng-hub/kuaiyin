package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type KyVideoGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:contains;column:title;table:ky_video" comment:"视频标题"`
	Msg            string `form:"msg"  search:"type:contains;column:msg;table:ky_video" comment:"视频信息"`
	UserId         int64  `form:"userId"  search:"type:exact;column:user_id;table:ky_video" comment:"快音号"`
	KyVideoOrder
}

type KyVideoOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:ky_video"`
	Title     string `form:"titleOrder"  search:"type:order;column:title;table:ky_video"`
	Msg       string `form:"msgOrder"  search:"type:order;column:msg;table:ky_video"`
	Like      string `form:"likeOrder"  search:"type:order;column:like;table:ky_video"`
	Sms       string `form:"smsOrder"  search:"type:order;column:sms;table:ky_video"`
	Src       string `form:"srcOrder"  search:"type:order;column:src;table:ky_video"`
	Href      string `form:"hrefOrder"  search:"type:order;column:href;table:ky_video"`
	UserId    string `form:"userIdOrder"  search:"type:order;column:user_id;table:ky_video"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ky_video"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ky_video"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ky_video"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:ky_video"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:ky_video"`
}

func (m *KyVideoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type KyVideoInsertReq struct {
	Id     int    `json:"-" comment:"主键编码"` // 主键编码
	Title  string `json:"title" comment:"视频标题"`
	Msg    string `json:"msg" comment:"视频信息"`
	Like   int64  `json:"like" comment:"喜欢"`
	Sms    int64  `json:"sms" comment:"评论数"`
	Src    string `json:"src" comment:"视频地址"`
	Href   string `json:"href" comment:"图片"`
	UserId int64  `json:"userId,string" comment:"快音号"`
	common.ControlBy
}

func (s *KyVideoInsertReq) Generate(model *models.KyVideo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Msg = s.Msg
	model.Like = s.Like
	model.Sms = s.Sms
	model.Src = s.Src
	model.Href = s.Href
	model.UserId = s.UserId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *KyVideoInsertReq) GetId() interface{} {
	return s.Id
}

type KyVideoUpdateReq struct {
	Id     int    `uri:"id" comment:"主键编码"` // 主键编码
	Title  string `json:"title" comment:"视频标题"`
	Msg    string `json:"msg" comment:"视频信息"`
	Like   int64  `json:"like" comment:"喜欢"`
	Sms    int64  `json:"sms" comment:"评论数"`
	Src    string `json:"src" comment:"视频地址"`
	Href   string `json:"href" comment:"图片"`
	UserId int64  `json:"userId" comment:"快音号"`
	common.ControlBy
}

func (s *KyVideoUpdateReq) Generate(model *models.KyVideo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Msg = s.Msg
	model.Like = s.Like
	model.Sms = s.Sms
	model.Src = s.Src
	model.Href = s.Href
	model.UserId = s.UserId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *KyVideoUpdateReq) GetId() interface{} {
	return s.Id
}

// KyVideoGetReq 功能获取请求参数
type KyVideoGetReq struct {
	Id int `uri:"id"`
}

func (s *KyVideoGetReq) GetId() interface{} {
	return s.Id
}

// KyVideoDeleteReq 功能删除请求参数
type KyVideoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *KyVideoDeleteReq) GetId() interface{} {
	return s.Ids
}
