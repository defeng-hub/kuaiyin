package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusDocTplGetPageReq struct {
	dto.Pagination     `search:"-"`
    UnitName string `form:"unitName"  search:"type:contains;column:unit_name;table:bus_doc_tpl" comment:"单位名称"`
    ApplyName string `form:"applyName"  search:"type:contains;column:apply_name;table:bus_doc_tpl" comment:"申请模板名称"`
    AnnounceName string `form:"announceName"  search:"type:contains;column:announce_name;table:bus_doc_tpl" comment:"通知模板名称"`
    BusDocTplOrder
}

type BusDocTplOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_doc_tpl"`
    UnitName string `form:"unitNameOrder"  search:"type:order;column:unit_name;table:bus_doc_tpl"`
    ApplyName string `form:"applyNameOrder"  search:"type:order;column:apply_name;table:bus_doc_tpl"`
    ApplyPath string `form:"applyPathOrder"  search:"type:order;column:apply_path;table:bus_doc_tpl"`
    AnnounceName string `form:"announceNameOrder"  search:"type:order;column:announce_name;table:bus_doc_tpl"`
    AnnouncePath string `form:"announcePathOrder"  search:"type:order;column:announce_path;table:bus_doc_tpl"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_doc_tpl"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_doc_tpl"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_doc_tpl"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_doc_tpl"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_doc_tpl"`
    
}

func (m *BusDocTplGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusDocTplInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    UnitName string `json:"unitName" comment:"单位名称"`
    ApplyName string `json:"applyName" comment:"申请模板名称"`
    ApplyPath string `json:"applyPath" comment:"申请模板路径"`
    AnnounceName string `json:"announceName" comment:"通知模板名称"`
    AnnouncePath string `json:"announcePath" comment:"通知模版路径"`
    common.ControlBy
}

func (s *BusDocTplInsertReq) Generate(model *models.BusDocTpl)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UnitName = s.UnitName
    model.ApplyName = s.ApplyName
    model.ApplyPath = s.ApplyPath
    model.AnnounceName = s.AnnounceName
    model.AnnouncePath = s.AnnouncePath
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusDocTplInsertReq) GetId() interface{} {
	return s.Id
}

type BusDocTplUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    UnitName string `json:"unitName" comment:"单位名称"`
    ApplyName string `json:"applyName" comment:"申请模板名称"`
    ApplyPath string `json:"applyPath" comment:"申请模板路径"`
    AnnounceName string `json:"announceName" comment:"通知模板名称"`
    AnnouncePath string `json:"announcePath" comment:"通知模版路径"`
    common.ControlBy
}

func (s *BusDocTplUpdateReq) Generate(model *models.BusDocTpl)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UnitName = s.UnitName
    model.ApplyName = s.ApplyName
    model.ApplyPath = s.ApplyPath
    model.AnnounceName = s.AnnounceName
    model.AnnouncePath = s.AnnouncePath
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusDocTplUpdateReq) GetId() interface{} {
	return s.Id
}

// BusDocTplGetReq 功能获取请求参数
type BusDocTplGetReq struct {
     Id int `uri:"id"`
}
func (s *BusDocTplGetReq) GetId() interface{} {
	return s.Id
}

// BusDocTplDeleteReq 功能删除请求参数
type BusDocTplDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusDocTplDeleteReq) GetId() interface{} {
	return s.Ids
}
