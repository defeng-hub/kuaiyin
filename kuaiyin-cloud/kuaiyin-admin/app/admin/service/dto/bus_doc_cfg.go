package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusDocCfgGetPageReq struct {
	dto.Pagination     `search:"-"`
    BusDocCfgOrder
}

type BusDocCfgOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_doc_cfg"`
    FillingUnit string `form:"fillingUnitOrder"  search:"type:order;column:filling_unit;table:bus_doc_cfg"`
    CaseUnit string `form:"caseUnitOrder"  search:"type:order;column:case_unit;table:bus_doc_cfg"`
    UnitTitle string `form:"unitTitleOrder"  search:"type:order;column:unit_title;table:bus_doc_cfg"`
    LegalDocInfo string `form:"legalDocInfoOrder"  search:"type:order;column:legal_doc_info;table:bus_doc_cfg"`
    Person1 string `form:"person1Order"  search:"type:order;column:person1;table:bus_doc_cfg"`
    Phone1 string `form:"phone1Order"  search:"type:order;column:phone1;table:bus_doc_cfg"`
    Person2 string `form:"person2Order"  search:"type:order;column:person2;table:bus_doc_cfg"`
    Phone2 string `form:"phone2Order"  search:"type:order;column:phone2;table:bus_doc_cfg"`
    PUnitNum string `form:"pUnitNumOrder"  search:"type:order;column:p_unit_num;table:bus_doc_cfg"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_doc_cfg"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_doc_cfg"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_doc_cfg"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_doc_cfg"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_doc_cfg"`
    
}

func (m *BusDocCfgGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusDocCfgInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    FillingUnit string `json:"fillingUnit" comment:"填表单位"`
    CaseUnit string `json:"caseUnit" comment:"办案单位"`
    UnitTitle string `json:"unitTitle" comment:"文书单位抬头"`
    LegalDocInfo string `json:"legalDocInfo" comment:"法律文书信息"`
    Person1 string `json:"person1" comment:"经办人1"`
    Phone1 string `json:"phone1" comment:"电话1"`
    Person2 string `json:"person2" comment:"经办人2"`
    Phone2 string `json:"phone2" comment:"电话2"`
    PUnitNum string `json:"pUnitNum" comment:"上级单位编号"`
    common.ControlBy
}

func (s *BusDocCfgInsertReq) Generate(model *models.BusDocCfg)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.FillingUnit = s.FillingUnit
    model.CaseUnit = s.CaseUnit
    model.UnitTitle = s.UnitTitle
    model.LegalDocInfo = s.LegalDocInfo
    model.Person1 = s.Person1
    model.Phone1 = s.Phone1
    model.Person2 = s.Person2
    model.Phone2 = s.Phone2
    model.PUnitNum = s.PUnitNum
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusDocCfgInsertReq) GetId() interface{} {
	return s.Id
}

type BusDocCfgUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    FillingUnit string `json:"fillingUnit" comment:"填表单位"`
    CaseUnit string `json:"caseUnit" comment:"办案单位"`
    UnitTitle string `json:"unitTitle" comment:"文书单位抬头"`
    LegalDocInfo string `json:"legalDocInfo" comment:"法律文书信息"`
    Person1 string `json:"person1" comment:"经办人1"`
    Phone1 string `json:"phone1" comment:"电话1"`
    Person2 string `json:"person2" comment:"经办人2"`
    Phone2 string `json:"phone2" comment:"电话2"`
    PUnitNum string `json:"pUnitNum" comment:"上级单位编号"`
    common.ControlBy
}

func (s *BusDocCfgUpdateReq) Generate(model *models.BusDocCfg)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.FillingUnit = s.FillingUnit
    model.CaseUnit = s.CaseUnit
    model.UnitTitle = s.UnitTitle
    model.LegalDocInfo = s.LegalDocInfo
    model.Person1 = s.Person1
    model.Phone1 = s.Phone1
    model.Person2 = s.Person2
    model.Phone2 = s.Phone2
    model.PUnitNum = s.PUnitNum
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusDocCfgUpdateReq) GetId() interface{} {
	return s.Id
}

// BusDocCfgGetReq 功能获取请求参数
type BusDocCfgGetReq struct {
     Id int `uri:"id"`
}
func (s *BusDocCfgGetReq) GetId() interface{} {
	return s.Id
}

// BusDocCfgDeleteReq 功能删除请求参数
type BusDocCfgDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusDocCfgDeleteReq) GetId() interface{} {
	return s.Ids
}
