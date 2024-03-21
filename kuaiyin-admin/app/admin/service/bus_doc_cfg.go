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

type BusDocCfg struct {
	service.Service
}

// GetPage 获取BusDocCfg列表
func (e *BusDocCfg) GetPage(c *dto.BusDocCfgGetPageReq, p *actions.DataPermission, list *[]models.BusDocCfg, count *int64) error {
	var err error
	var data models.BusDocCfg

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusDocCfgService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusDocCfg对象
func (e *BusDocCfg) Get(d *dto.BusDocCfgGetReq, p *actions.DataPermission, model *models.BusDocCfg) error {
	var data models.BusDocCfg

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusDocCfg error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusDocCfg对象
func (e *BusDocCfg) Insert(c *dto.BusDocCfgInsertReq) error {
    var err error
    var data models.BusDocCfg
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusDocCfgService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusDocCfg对象
func (e *BusDocCfg) Update(c *dto.BusDocCfgUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusDocCfg{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusDocCfgService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusDocCfg
func (e *BusDocCfg) Remove(d *dto.BusDocCfgDeleteReq, p *actions.DataPermission) error {
	var data models.BusDocCfg

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusDocCfg error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
