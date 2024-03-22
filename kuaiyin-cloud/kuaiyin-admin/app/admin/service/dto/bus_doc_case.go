package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusDocCaseGetPageReq struct {
	dto.Pagination     `search:"-"`
    CaseName string `form:"caseName"  search:"type:contains;column:case_name;table:bus_doc_case" comment:"案件名称"`
    CaseCode string `form:"caseCode"  search:"type:exact;column:case_code;table:bus_doc_case" comment:"案件编号"`
    CaseType string `form:"caseType"  search:"type:exact;column:case_type;table:bus_doc_case" comment:"案件类别"`
    GetType string `form:"getType"  search:"type:exact;column:get_type;table:bus_doc_case" comment:"调取种类"`
    BusDocCaseOrder
}

type BusDocCaseOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_doc_case"`
    CaseName string `form:"caseNameOrder"  search:"type:order;column:case_name;table:bus_doc_case"`
    CaseCode string `form:"caseCodeOrder"  search:"type:order;column:case_code;table:bus_doc_case"`
    CaseType string `form:"caseTypeOrder"  search:"type:order;column:case_type;table:bus_doc_case"`
    CaseInfo string `form:"caseInfoOrder"  search:"type:order;column:case_info;table:bus_doc_case"`
    SuspectsName string `form:"suspectsNameOrder"  search:"type:order;column:suspects_name;table:bus_doc_case"`
    SuspectsGender string `form:"suspectsGenderOrder"  search:"type:order;column:suspects_gender;table:bus_doc_case"`
    SuspectsBirthday string `form:"suspectsBirthdayOrder"  search:"type:order;column:suspects_birthday;table:bus_doc_case"`
    GetType string `form:"getTypeOrder"  search:"type:order;column:get_type;table:bus_doc_case"`
    Person1 string `form:"person1Order"  search:"type:order;column:person1;table:bus_doc_case"`
    Phone1 string `form:"phone1Order"  search:"type:order;column:phone1;table:bus_doc_case"`
    Person2 string `form:"person2Order"  search:"type:order;column:person2;table:bus_doc_case"`
    Phone2 string `form:"phone2Order"  search:"type:order;column:phone2;table:bus_doc_case"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_doc_case"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_doc_case"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_doc_case"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_doc_case"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_doc_case"`
    
}

func (m *BusDocCaseGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusDocCaseInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    CaseName string `json:"caseName" comment:"案件名称"`
    CaseCode string `json:"caseCode" comment:"案件编号"`
    CaseType string `json:"caseType" comment:"案件类别"`
    CaseInfo string `json:"caseInfo" comment:"简要案情"`
    SuspectsName string `json:"suspectsName" comment:"嫌疑人姓名"`
    SuspectsGender string `json:"suspectsGender" comment:"嫌疑人性别"`
    SuspectsBirthday string `json:"suspectsBirthday" comment:"嫌疑人生日"`
    GetType string `json:"getType" comment:"调取种类"`
    Person1 string `json:"person1" comment:"经办人1"`
    Phone1 string `json:"phone1" comment:"电话1"`
    Person2 string `json:"person2" comment:"经办人2"`
    Phone2 string `json:"phone2" comment:"电话2"`
    common.ControlBy
}

func (s *BusDocCaseInsertReq) Generate(model *models.BusDocCase)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CaseName = s.CaseName
    model.CaseCode = s.CaseCode
    model.CaseType = s.CaseType
    model.CaseInfo = s.CaseInfo
    model.SuspectsName = s.SuspectsName
    model.SuspectsGender = s.SuspectsGender
    model.SuspectsBirthday = s.SuspectsBirthday
    model.GetType = s.GetType
    model.Person1 = s.Person1
    model.Phone1 = s.Phone1
    model.Person2 = s.Person2
    model.Phone2 = s.Phone2
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusDocCaseInsertReq) GetId() interface{} {
	return s.Id
}

type BusDocCaseUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    CaseName string `json:"caseName" comment:"案件名称"`
    CaseCode string `json:"caseCode" comment:"案件编号"`
    CaseType string `json:"caseType" comment:"案件类别"`
    CaseInfo string `json:"caseInfo" comment:"简要案情"`
    SuspectsName string `json:"suspectsName" comment:"嫌疑人姓名"`
    SuspectsGender string `json:"suspectsGender" comment:"嫌疑人性别"`
    SuspectsBirthday string `json:"suspectsBirthday" comment:"嫌疑人生日"`
    GetType string `json:"getType" comment:"调取种类"`
    Person1 string `json:"person1" comment:"经办人1"`
    Phone1 string `json:"phone1" comment:"电话1"`
    Person2 string `json:"person2" comment:"经办人2"`
    Phone2 string `json:"phone2" comment:"电话2"`
    common.ControlBy
}

func (s *BusDocCaseUpdateReq) Generate(model *models.BusDocCase)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CaseName = s.CaseName
    model.CaseCode = s.CaseCode
    model.CaseType = s.CaseType
    model.CaseInfo = s.CaseInfo
    model.SuspectsName = s.SuspectsName
    model.SuspectsGender = s.SuspectsGender
    model.SuspectsBirthday = s.SuspectsBirthday
    model.GetType = s.GetType
    model.Person1 = s.Person1
    model.Phone1 = s.Phone1
    model.Person2 = s.Person2
    model.Phone2 = s.Phone2
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusDocCaseUpdateReq) GetId() interface{} {
	return s.Id
}

// BusDocCaseGetReq 功能获取请求参数
type BusDocCaseGetReq struct {
     Id int `uri:"id"`
}
func (s *BusDocCaseGetReq) GetId() interface{} {
	return s.Id
}

// BusDocCaseDeleteReq 功能删除请求参数
type BusDocCaseDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusDocCaseDeleteReq) GetId() interface{} {
	return s.Ids
}
