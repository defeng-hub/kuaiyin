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

type BusDocCase struct {
	service.Service
}

// GetPage 获取BusDocCase列表
func (e *BusDocCase) GetPage(c *dto.BusDocCaseGetPageReq, p *actions.DataPermission, list *[]models.BusDocCase, count *int64) error {
	var err error
	var data models.BusDocCase

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusDocCaseService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusDocCase对象
func (e *BusDocCase) Get(d *dto.BusDocCaseGetReq, p *actions.DataPermission, model *models.BusDocCase) error {
	var data models.BusDocCase

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusDocCase error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusDocCase对象
func (e *BusDocCase) Insert(c *dto.BusDocCaseInsertReq) error {
    var err error
    var data models.BusDocCase
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusDocCaseService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusDocCase对象
func (e *BusDocCase) Update(c *dto.BusDocCaseUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusDocCase{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusDocCaseService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusDocCase
func (e *BusDocCase) Remove(d *dto.BusDocCaseDeleteReq, p *actions.DataPermission) error {
	var data models.BusDocCase

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusDocCase error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
