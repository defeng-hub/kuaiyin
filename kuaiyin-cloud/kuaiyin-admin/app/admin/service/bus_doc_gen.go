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

type BusDocGen struct {
	service.Service
}

// GetPage 获取BusDocGen列表
func (e *BusDocGen) GetPage(c *dto.BusDocGenGetPageReq, p *actions.DataPermission, list *[]models.BusDocGen, count *int64) error {
	var err error
	var data models.BusDocGen

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusDocGenService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusDocGen对象
func (e *BusDocGen) Get(d *dto.BusDocGenGetReq, p *actions.DataPermission, model *models.BusDocGen) error {
	var data models.BusDocGen

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusDocGen error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusDocGen对象
func (e *BusDocGen) Insert(c *dto.BusDocGenInsertReq) error {
    var err error
    var data models.BusDocGen
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusDocGenService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusDocGen对象
func (e *BusDocGen) Update(c *dto.BusDocGenUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusDocGen{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusDocGenService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusDocGen
func (e *BusDocGen) Remove(d *dto.BusDocGenDeleteReq, p *actions.DataPermission) error {
	var data models.BusDocGen

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusDocGen error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
