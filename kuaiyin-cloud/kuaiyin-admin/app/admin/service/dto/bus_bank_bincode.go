package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusBankBincodeGetPageReq struct {
	dto.Pagination     `search:"-"`
    BankCode string `form:"bankCode"  search:"type:exact;column:bank_code;table:bus_bank_bincode" comment:"银行编码"`
    BankKh string `form:"bankKh"  search:"type:contains;column:bank_kh;table:bus_bank_bincode" comment:"银行前缀卡号"`
    BankName string `form:"bankName"  search:"type:contains;column:bank_name;table:bus_bank_bincode" comment:"银行名称"`
    CardType string `form:"cardType"  search:"type:exact;column:card_type;table:bus_bank_bincode" comment:"卡类型"`
    BusBankBincodeOrder
}

type BusBankBincodeOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_bank_bincode"`
    BankCode string `form:"bankCodeOrder"  search:"type:order;column:bank_code;table:bus_bank_bincode"`
    BankKh string `form:"bankKhOrder"  search:"type:order;column:bank_kh;table:bus_bank_bincode"`
    BankName string `form:"bankNameOrder"  search:"type:order;column:bank_name;table:bus_bank_bincode"`
    BankAbbr string `form:"bankAbbrOrder"  search:"type:order;column:bank_abbr;table:bus_bank_bincode"`
    CardType string `form:"cardTypeOrder"  search:"type:order;column:card_type;table:bus_bank_bincode"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_bank_bincode"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_bank_bincode"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_bank_bincode"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_bank_bincode"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_bank_bincode"`
    
}

func (m *BusBankBincodeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusBankBincodeInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    BankCode string `json:"bankCode" comment:"银行编码"`
    BankKh string `json:"bankKh" comment:"银行前缀卡号"`
    BankName string `json:"bankName" comment:"银行名称"`
    BankAbbr string `json:"bankAbbr" comment:"银行名称缩写"`
    CardType string `json:"cardType" comment:"卡类型"`
    common.ControlBy
}

func (s *BusBankBincodeInsertReq) Generate(model *models.BusBankBincode)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.BankKh = s.BankKh
    model.BankName = s.BankName
    model.BankAbbr = s.BankAbbr
    model.CardType = s.CardType
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusBankBincodeInsertReq) GetId() interface{} {
	return s.Id
}

type BusBankBincodeUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    BankCode string `json:"bankCode" comment:"银行编码"`
    BankKh string `json:"bankKh" comment:"银行前缀卡号"`
    BankName string `json:"bankName" comment:"银行名称"`
    BankAbbr string `json:"bankAbbr" comment:"银行名称缩写"`
    CardType string `json:"cardType" comment:"卡类型"`
    common.ControlBy
}

func (s *BusBankBincodeUpdateReq) Generate(model *models.BusBankBincode)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.BankKh = s.BankKh
    model.BankName = s.BankName
    model.BankAbbr = s.BankAbbr
    model.CardType = s.CardType
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusBankBincodeUpdateReq) GetId() interface{} {
	return s.Id
}

// BusBankBincodeGetReq 功能获取请求参数
type BusBankBincodeGetReq struct {
     Id int `uri:"id"`
}
func (s *BusBankBincodeGetReq) GetId() interface{} {
	return s.Id
}

// BusBankBincodeDeleteReq 功能删除请求参数
type BusBankBincodeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusBankBincodeDeleteReq) GetId() interface{} {
	return s.Ids
}
