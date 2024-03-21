package models

import (

	"go-admin/common/models"

)

type TbSyncLog struct {
    models.Model
    
    Env string `json:"env" gorm:"type:varchar(128);comment:环境"` 
    LogName string `json:"logName" gorm:"type:varchar(128);comment:文件名"` 
    Status int64 `json:"status" gorm:"type:tinyint(1);comment:同步状态"` 
    models.ModelTime
    models.ControlBy
}

func (TbSyncLog) TableName() string {
    return "tb_sync_log"
}

func (e *TbSyncLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbSyncLog) GetId() interface{} {
	return e.Id
}