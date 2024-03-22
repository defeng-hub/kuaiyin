package models

import (

	"go-admin/common/models"

)

type BusDocTpl struct {
    models.Model
    
    UnitName string `json:"unitName" gorm:"type:int(20);comment:单位名称"` 
    ApplyName string `json:"applyName" gorm:"type:varchar(128);comment:申请模板名称"` 
    ApplyPath string `json:"applyPath" gorm:"type:varchar(255);comment:申请模板路径"` 
    AnnounceName string `json:"announceName" gorm:"type:varchar(128);comment:通知模板名称"` 
    AnnouncePath string `json:"announcePath" gorm:"type:varchar(255);comment:通知模版路径"` 
    models.ModelTime
    models.ControlBy
}

func (BusDocTpl) TableName() string {
    return "bus_doc_tpl"
}

func (e *BusDocTpl) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusDocTpl) GetId() interface{} {
	return e.Id
}