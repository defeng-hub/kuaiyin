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

type BusBankCardtypes struct {
	service.Service
}

// GetPage 获取BusBankCardtypes列表
func (e *BusBankCardtypes) GetPage(c *dto.BusBankCardtypesGetPageReq, p *actions.DataPermission, list *[]models.BusBankCardtypes, count *int64) error {
	var err error
	var data models.BusBankCardtypes

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusBankCardtypesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusBankCardtypes对象
func (e *BusBankCardtypes) Get(d *dto.BusBankCardtypesGetReq, p *actions.DataPermission, model *models.BusBankCardtypes) error {
	var data models.BusBankCardtypes

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusBankCardtypes error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusBankCardtypes对象
func (e *BusBankCardtypes) Insert(c *dto.BusBankCardtypesInsertReq) error {
    var err error
    var data models.BusBankCardtypes
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusBankCardtypesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusBankCardtypes对象
func (e *BusBankCardtypes) Update(c *dto.BusBankCardtypesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusBankCardtypes{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusBankCardtypesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusBankCardtypes
func (e *BusBankCardtypes) Remove(d *dto.BusBankCardtypesDeleteReq, p *actions.DataPermission) error {
	var data models.BusBankCardtypes

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusBankCardtypes error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
