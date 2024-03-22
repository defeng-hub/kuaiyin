package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusDocGenGetPageReq struct {
	dto.Pagination     `search:"-"`
    CaseId string `form:"caseId"  search:"type:exact;column:case_id;table:bus_doc_gen" comment:"案件ID"`
    UnitId string `form:"unitId"  search:"type:exact;column:unit_id;table:bus_doc_gen" comment:"单位ID"`
    ReviewCode string `form:"reviewCode"  search:"type:exact;column:review_code;table:bus_doc_gen" comment:"审核员编码"`
    Status string `form:"status"  search:"type:exact;column:status;table:bus_doc_gen" comment:"审批状态"`
    ProcStatus string `form:"procStatus"  search:"type:exact;column:proc_status;table:bus_doc_gen" comment:"处理状态"`
    BusDocGenOrder
}

type BusDocGenOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_doc_gen"`
    CaseId string `form:"caseIdOrder"  search:"type:order;column:case_id;table:bus_doc_gen"`
    UnitId string `form:"unitIdOrder"  search:"type:order;column:unit_id;table:bus_doc_gen"`
    StartNo string `form:"startNoOrder"  search:"type:order;column:start_no;table:bus_doc_gen"`
    ReviewCode string `form:"reviewCodeOrder"  search:"type:order;column:review_code;table:bus_doc_gen"`
    FilePath string `form:"filePathOrder"  search:"type:order;column:file_path;table:bus_doc_gen"`
    RstPath string `form:"rstPathOrder"  search:"type:order;column:rst_path;table:bus_doc_gen"`
    InvolveRow string `form:"involveRowOrder"  search:"type:order;column:involve_row;table:bus_doc_gen"`
    TargetRow string `form:"targetRowOrder"  search:"type:order;column:target_row;table:bus_doc_gen"`
    BankRow string `form:"bankRowOrder"  search:"type:order;column:bank_row;table:bus_doc_gen"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:bus_doc_gen"`
    ProcStatus string `form:"procStatusOrder"  search:"type:order;column:proc_status;table:bus_doc_gen"`
    Reason string `form:"reasonOrder"  search:"type:order;column:reason;table:bus_doc_gen"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:bus_doc_gen"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_doc_gen"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_doc_gen"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_doc_gen"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_doc_gen"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_doc_gen"`
    
}

func (m *BusDocGenGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusDocGenInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    CaseId string `json:"caseId" comment:"案件ID"`
    UnitId string `json:"unitId" comment:"单位ID"`
    StartNo string `json:"startNo" comment:"通知序列号"`
    ReviewCode string `json:"reviewCode" comment:"审核员编码"`
    FilePath string `json:"filePath" comment:"上传文件地址"`
    RstPath string `json:"rstPath" comment:"生成文件地址"`
    InvolveRow string `json:"involveRow" comment:"涉案对象列号"`
    TargetRow string `json:"targetRow" comment:"调查目标列号"`
    BankRow string `json:"bankRow" comment:"所查银行列号"`
    Status string `json:"status" comment:"审批状态"`
    ProcStatus string `json:"procStatus" comment:"处理状态"`
    Reason string `json:"reason" comment:"审批说明"`
    Remark string `json:"remark" comment:"备注"`
    common.ControlBy
}

func (s *BusDocGenInsertReq) Generate(model *models.BusDocGen)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CaseId = s.CaseId
    model.UnitId = s.UnitId
    model.StartNo = s.StartNo
    model.ReviewCode = s.ReviewCode
    model.FilePath = s.FilePath
    model.RstPath = s.RstPath
    model.InvolveRow = s.InvolveRow
    model.TargetRow = s.TargetRow
    model.BankRow = s.BankRow
    model.Status = s.Status
    model.ProcStatus = s.ProcStatus
    model.Reason = s.Reason
    model.Remark = s.Remark
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusDocGenInsertReq) GetId() interface{} {
	return s.Id
}

type BusDocGenUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    CaseId string `json:"caseId" comment:"案件ID"`
    UnitId string `json:"unitId" comment:"单位ID"`
    StartNo string `json:"startNo" comment:"通知序列号"`
    ReviewCode string `json:"reviewCode" comment:"审核员编码"`
    FilePath string `json:"filePath" comment:"上传文件地址"`
    RstPath string `json:"rstPath" comment:"生成文件地址"`
    InvolveRow string `json:"involveRow" comment:"涉案对象列号"`
    TargetRow string `json:"targetRow" comment:"调查目标列号"`
    BankRow string `json:"bankRow" comment:"所查银行列号"`
    Status string `json:"status" comment:"审批状态"`
    ProcStatus string `json:"procStatus" comment:"处理状态"`
    Reason string `json:"reason" comment:"审批说明"`
    Remark string `json:"remark" comment:"备注"`
    common.ControlBy
}

func (s *BusDocGenUpdateReq) Generate(model *models.BusDocGen)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CaseId = s.CaseId
    model.UnitId = s.UnitId
    model.StartNo = s.StartNo
    model.ReviewCode = s.ReviewCode
    model.FilePath = s.FilePath
    model.RstPath = s.RstPath
    model.InvolveRow = s.InvolveRow
    model.TargetRow = s.TargetRow
    model.BankRow = s.BankRow
    model.Status = s.Status
    model.ProcStatus = s.ProcStatus
    model.Reason = s.Reason
    model.Remark = s.Remark
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusDocGenUpdateReq) GetId() interface{} {
	return s.Id
}

// BusDocGenGetReq 功能获取请求参数
type BusDocGenGetReq struct {
     Id int `uri:"id"`
}
func (s *BusDocGenGetReq) GetId() interface{} {
	return s.Id
}

// BusDocGenDeleteReq 功能删除请求参数
type BusDocGenDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusDocGenDeleteReq) GetId() interface{} {
	return s.Ids
}
