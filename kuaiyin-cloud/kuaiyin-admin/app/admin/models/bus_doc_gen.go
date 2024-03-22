package models

import (

	"go-admin/common/models"

)

type BusDocGen struct {
    models.Model
    
    CaseId string `json:"caseId" gorm:"type:int(20);comment:案件ID"` 
    UnitId string `json:"unitId" gorm:"type:int(20);comment:单位ID"` 
    StartNo string `json:"startNo" gorm:"type:varchar(255);comment:通知序列号"` 
    ReviewCode string `json:"reviewCode" gorm:"type:varchar(255);comment:审核员编码"` 
    FilePath string `json:"filePath" gorm:"type:varchar(255);comment:上传文件地址"` 
    RstPath string `json:"rstPath" gorm:"type:varchar(255);comment:生成文件地址"` 
    InvolveRow string `json:"involveRow" gorm:"type:tinyint(3);comment:涉案对象列号"` 
    TargetRow string `json:"targetRow" gorm:"type:tinyint(3);comment:调查目标列号"` 
    BankRow string `json:"bankRow" gorm:"type:tinyint(3);comment:所查银行列号"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:审批状态"` 
    ProcStatus string `json:"procStatus" gorm:"type:tinyint(1);comment:处理状态"` 
    Reason string `json:"reason" gorm:"type:varchar(255);comment:审批说明"` 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:备注"` 
    models.ModelTime
    models.ControlBy
}

func (BusDocGen) TableName() string {
    return "bus_doc_gen"
}

func (e *BusDocGen) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusDocGen) GetId() interface{} {
	return e.Id
}