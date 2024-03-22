package models

import (

	"go-admin/common/models"

)

type BusBankCardtypes struct {
    models.Model
    
    BankCode string `json:"bankCode" gorm:"type:varchar(128);comment:银行编码"` 
    BankKh string `json:"bankKh" gorm:"type:int(10);comment:银行前缀卡号"` 
    BankName string `json:"bankName" gorm:"type:varchar(128);comment:银行名称"` 
    CardName string `json:"cardName" gorm:"type:varchar(128);comment:卡名称"` 
    CardLength string `json:"cardLength" gorm:"type:int(6);comment:卡号长度"` 
    models.ModelTime
    models.ControlBy
}

func (BusBankCardtypes) TableName() string {
    return "bus_bank_cardtypes"
}

func (e *BusBankCardtypes) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusBankCardtypes) GetId() interface{} {
	return e.Id
}