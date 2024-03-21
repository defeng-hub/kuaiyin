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

type BusDataClean struct {
	service.Service
}

// GetPage 获取BusDataClean列表
func (e *BusDataClean) GetPage(c *dto.BusDataCleanGetPageReq, p *actions.DataPermission, list *[]models.BusDataClean, count *int64) error {
	var err error
	var data models.BusDataClean

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusDataCleanService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusDataClean对象
func (e *BusDataClean) Get(d *dto.BusDataCleanGetReq, p *actions.DataPermission, model *models.BusDataClean) error {
	var data models.BusDataClean

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusDataClean error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusDataClean对象
func (e *BusDataClean) Insert(c *dto.BusDataCleanInsertReq) error {
    var err error
    var data models.BusDataClean
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusDataCleanService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusDataClean对象
func (e *BusDataClean) Update(c *dto.BusDataCleanUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusDataClean{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusDataCleanService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusDataClean
func (e *BusDataClean) Remove(d *dto.BusDataCleanDeleteReq, p *actions.DataPermission) error {
	var data models.BusDataClean

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusDataClean error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
