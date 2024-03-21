package models

import (
	"go-admin/common/models"
)

type BusSplit struct {
	models.Model

	TaskName   string `json:"taskName" gorm:"type:varchar(128);comment:任务名称"`
	FilePath   string `json:"filePath" gorm:"type:varchar(255);comment:源文件地址"`
	FileSize   string `json:"fileSize" gorm:"type:double(20,0);comment:文件大小"`
	RstPath    string `json:"rstPath" gorm:"type:varchar(255);comment:目标文件地址"`
	OutputName string `json:"outputName" gorm:"type:varchar(128);comment:拆分后名称"`
	SplitRow   string `json:"splitRow" gorm:"type:tinyint(3);comment:拆分列号"`
	Status     string `json:"status" gorm:"type:tinyint(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (BusSplit) TableName() string {
	return "bus_split"
}

func (e *BusSplit) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusSplit) GetId() interface{} {
	return e.Id
}
