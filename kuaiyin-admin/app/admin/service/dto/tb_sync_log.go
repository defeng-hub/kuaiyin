package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbSyncLogGetPageReq struct {
	dto.Pagination `search:"-"`
	Env            string `form:"env"  search:"type:exact;column:env;table:tb_sync_log" comment:"环境"`
	LogName        string `form:"logName"  search:"type:exact;column:log_name;table:tb_sync_log" comment:"文件名"`
	Status         int64  `form:"status"  search:"type:exact;column:status;table:tb_sync_log" comment:"同步状态"`
	TbSyncLogOrder
}

type TbSyncLogOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:tb_sync_log"`
	Env       string `form:"envOrder"  search:"type:order;column:env;table:tb_sync_log"`
	LogName   string `form:"logNameOrder"  search:"type:order;column:log_name;table:tb_sync_log"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:tb_sync_log"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_sync_log"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_sync_log"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_sync_log"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_sync_log"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_sync_log"`
}

func (m *TbSyncLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbSyncLogInsertReq struct {
	Id      int    `json:"-" comment:"主键编码"` // 主键编码
	Env     string `json:"env" comment:"环境"`
	LogName string `json:"logName" comment:"文件名"`
	Status  int64  `json:"status" comment:"同步状态"`
	common.ControlBy
}

func (s *TbSyncLogInsertReq) Generate(model *models.TbSyncLog) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Env = s.Env
	model.LogName = s.LogName
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *TbSyncLogInsertReq) GetId() interface{} {
	return s.Id
}

type TbSyncLogUpdateReq struct {
	Id      int    `uri:"id" comment:"主键编码"` // 主键编码
	Env     string `json:"env" comment:"环境"`
	LogName string `json:"logName" comment:"文件名"`
	Status  int64  `json:"status" comment:"同步状态"`
	common.ControlBy
}

func (s *TbSyncLogUpdateReq) Generate(model *models.TbSyncLog) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Env = s.Env
	model.LogName = s.LogName
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *TbSyncLogUpdateReq) GetId() interface{} {
	return s.Id
}

// TbSyncLogGetReq 功能获取请求参数
type TbSyncLogGetReq struct {
	Id int `uri:"id"`
}

func (s *TbSyncLogGetReq) GetId() interface{} {
	return s.Id
}

// TbSyncLogDeleteReq 功能删除请求参数
type TbSyncLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbSyncLogDeleteReq) GetId() interface{} {
	return s.Ids
}
