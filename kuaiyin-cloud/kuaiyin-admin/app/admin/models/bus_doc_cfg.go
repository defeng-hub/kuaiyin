package models

import (

	"go-admin/common/models"

)

type BusDocCfg struct {
    models.Model
    
    FillingUnit string `json:"fillingUnit" gorm:"type:varchar(255);comment:填表单位"` 
    CaseUnit string `json:"caseUnit" gorm:"type:varchar(255);comment:办案单位"` 
    UnitTitle string `json:"unitTitle" gorm:"type:varchar(255);comment:文书单位抬头"` 
    LegalDocInfo string `json:"legalDocInfo" gorm:"type:varchar(255);comment:法律文书信息"` 
    Person1 string `json:"person1" gorm:"type:varchar(255);comment:经办人1"` 
    Phone1 string `json:"phone1" gorm:"type:varchar(255);comment:电话1"` 
    Person2 string `json:"person2" gorm:"type:varchar(255);comment:经办人2"` 
    Phone2 string `json:"phone2" gorm:"type:varchar(255);comment:电话2"` 
    PUnitNum string `json:"pUnitNum" gorm:"type:varchar(255);comment:上级单位编号"` 
    models.ModelTime
    models.ControlBy
}

func (BusDocCfg) TableName() string {
    return "bus_doc_cfg"
}

func (e *BusDocCfg) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusDocCfg) GetId() interface{} {
	return e.Id
}