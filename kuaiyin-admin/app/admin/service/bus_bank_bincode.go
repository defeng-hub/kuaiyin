package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type BusBankBincode struct {
	service.Service
}

// GetPage 获取BusBankBincode列表
func (e *BusBankBincode) GetPage(c *dto.BusBankBincodeGetPageReq, p *actions.DataPermission, list *[]models.BusBankBincode, count *int64) error {
	var err error
	var data models.BusBankBincode

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusBankBincodeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusBankBincode对象
func (e *BusBankBincode) Get(d *dto.BusBankBincodeGetReq, p *actions.DataPermission, model *models.BusBankBincode) error {
	var data models.BusBankBincode

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusBankBincode error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusBankBincode对象
func (e *BusBankBincode) Insert(c *dto.BusBankBincodeInsertReq) error {
    var err error
    var data models.BusBankBincode
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusBankBincodeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusBankBincode对象
func (e *BusBankBincode) Update(c *dto.BusBankBincodeUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusBankBincode{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusBankBincodeService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusBankBincode
func (e *BusBankBincode) Remove(d *dto.BusBankBincodeDeleteReq, p *actions.DataPermission) error {
	var data models.BusBankBincode

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusBankBincode error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
