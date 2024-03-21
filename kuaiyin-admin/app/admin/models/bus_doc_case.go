package models

import (

	"go-admin/common/models"

)

type BusDocCase struct {
    models.Model
    
    CaseName string `json:"caseName" gorm:"type:varchar(128);comment:案件名称"` 
    CaseCode string `json:"caseCode" gorm:"type:varchar(64);comment:案件编号"` 
    CaseType string `json:"caseType" gorm:"type:varchar(16);comment:案件类别"` 
    CaseInfo string `json:"caseInfo" gorm:"type:varchar(255);comment:简要案情"` 
    SuspectsName string `json:"suspectsName" gorm:"type:varchar(32);comment:嫌疑人姓名"` 
    SuspectsGender string `json:"suspectsGender" gorm:"type:varchar(8);comment:嫌疑人性别"` 
    SuspectsBirthday string `json:"suspectsBirthday" gorm:"type:varchar(16);comment:嫌疑人生日"` 
    GetType string `json:"getType" gorm:"type:varchar(16);comment:调取种类"` 
    Person1 string `json:"person1" gorm:"type:varchar(255);comment:经办人1"` 
    Phone1 string `json:"phone1" gorm:"type:varchar(255);comment:电话1"` 
    Person2 string `json:"person2" gorm:"type:varchar(255);comment:经办人2"` 
    Phone2 string `json:"phone2" gorm:"type:varchar(255);comment:电话2"` 
    models.ModelTime
    models.ControlBy
}

func (BusDocCase) TableName() string {
    return "bus_doc_case"
}

func (e *BusDocCase) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusDocCase) GetId() interface{} {
	return e.Id
}