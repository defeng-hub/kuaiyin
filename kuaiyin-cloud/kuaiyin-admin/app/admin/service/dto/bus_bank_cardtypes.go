package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusBankCardtypesGetPageReq struct {
	dto.Pagination     `search:"-"`
    BankCode string `form:"bankCode"  search:"type:exact;column:bank_code;table:bus_bank_cardtypes" comment:"银行编码"`
    BankKh string `form:"bankKh"  search:"type:contains;column:bank_kh;table:bus_bank_cardtypes" comment:"银行前缀卡号"`
    BankName string `form:"bankName"  search:"type:contains;column:bank_name;table:bus_bank_cardtypes" comment:"银行名称"`
    CardName string `form:"cardName"  search:"type:contains;column:card_name;table:bus_bank_cardtypes" comment:"卡名称"`
    BusBankCardtypesOrder
}

type BusBankCardtypesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_bank_cardtypes"`
    BankCode string `form:"bankCodeOrder"  search:"type:order;column:bank_code;table:bus_bank_cardtypes"`
    BankKh string `form:"bankKhOrder"  search:"type:order;column:bank_kh;table:bus_bank_cardtypes"`
    BankName string `form:"bankNameOrder"  search:"type:order;column:bank_name;table:bus_bank_cardtypes"`
    CardName string `form:"cardNameOrder"  search:"type:order;column:card_name;table:bus_bank_cardtypes"`
    CardLength string `form:"cardLengthOrder"  search:"type:order;column:card_length;table:bus_bank_cardtypes"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_bank_cardtypes"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_bank_cardtypes"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_bank_cardtypes"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_bank_cardtypes"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_bank_cardtypes"`
    
}

func (m *BusBankCardtypesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusBankCardtypesInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    BankCode string `json:"bankCode" comment:"银行编码"`
    BankKh string `json:"bankKh" comment:"银行前缀卡号"`
    BankName string `json:"bankName" comment:"银行名称"`
    CardName string `json:"cardName" comment:"卡名称"`
    CardLength string `json:"cardLength" comment:"卡号长度"`
    common.ControlBy
}

func (s *BusBankCardtypesInsertReq) Generate(model *models.BusBankCardtypes)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.BankKh = s.BankKh
    model.BankName = s.BankName
    model.CardName = s.CardName
    model.CardLength = s.CardLength
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusBankCardtypesInsertReq) GetId() interface{} {
	return s.Id
}

type BusBankCardtypesUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    BankCode string `json:"bankCode" comment:"银行编码"`
    BankKh string `json:"bankKh" comment:"银行前缀卡号"`
    BankName string `json:"bankName" comment:"银行名称"`
    CardName string `json:"cardName" comment:"卡名称"`
    CardLength string `json:"cardLength" comment:"卡号长度"`
    common.ControlBy
}

func (s *BusBankCardtypesUpdateReq) Generate(model *models.BusBankCardtypes)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.BankKh = s.BankKh
    model.BankName = s.BankName
    model.CardName = s.CardName
    model.CardLength = s.CardLength
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusBankCardtypesUpdateReq) GetId() interface{} {
	return s.Id
}

// BusBankCardtypesGetReq 功能获取请求参数
type BusBankCardtypesGetReq struct {
     Id int `uri:"id"`
}
func (s *BusBankCardtypesGetReq) GetId() interface{} {
	return s.Id
}

// BusBankCardtypesDeleteReq 功能删除请求参数
type BusBankCardtypesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusBankCardtypesDeleteReq) GetId() interface{} {
	return s.Ids
}
