package models

import (

	"go-admin/common/models"

)

type BusDataClean struct {
    models.Model
    
    TaskName string `json:"taskName" gorm:"type:varchar(128);comment:任务名称"` 
    FileName string `json:"fileName" gorm:"type:varchar(255);comment:源文件名称"` 
    FilePath string `json:"filePath" gorm:"type:varchar(255);comment:源文件地址"` 
    RstPath string `json:"rstPath" gorm:"type:varchar(255);comment:目标文件地址"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态"` 
    models.ModelTime
    models.ControlBy
}

func (BusDataClean) TableName() string {
    return "bus_data_clean"
}

func (e *BusDataClean) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusDataClean) GetId() interface{} {
	return e.Id
}