package models

import (

	"go-admin/common/models"

)

type BusBankBincode struct {
    models.Model
    
    BankCode string `json:"bankCode" gorm:"type:varchar(128);comment:银行编码"` 
    BankKh string `json:"bankKh" gorm:"type:int(10);comment:银行前缀卡号"` 
    BankName string `json:"bankName" gorm:"type:varchar(128);comment:银行名称"` 
    BankAbbr string `json:"bankAbbr" gorm:"type:varchar(128);comment:银行名称缩写"` 
    CardType string `json:"cardType" gorm:"type:varchar(128);comment:卡类型"` 
    models.ModelTime
    models.ControlBy
}

func (BusBankBincode) TableName() string {
    return "bus_bank_bincode"
}

func (e *BusBankBincode) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BusBankBincode) GetId() interface{} {
	return e.Id
}