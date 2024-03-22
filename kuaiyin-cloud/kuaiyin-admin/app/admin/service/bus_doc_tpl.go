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

type BusDocTpl struct {
	service.Service
}

// GetPage 获取BusDocTpl列表
func (e *BusDocTpl) GetPage(c *dto.BusDocTplGetPageReq, p *actions.DataPermission, list *[]models.BusDocTpl, count *int64) error {
	var err error
	var data models.BusDocTpl

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BusDocTplService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BusDocTpl对象
func (e *BusDocTpl) Get(d *dto.BusDocTplGetReq, p *actions.DataPermission, model *models.BusDocTpl) error {
	var data models.BusDocTpl

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBusDocTpl error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BusDocTpl对象
func (e *BusDocTpl) Insert(c *dto.BusDocTplInsertReq) error {
    var err error
    var data models.BusDocTpl
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BusDocTplService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BusDocTpl对象
func (e *BusDocTpl) Update(c *dto.BusDocTplUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BusDocTpl{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BusDocTplService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BusDocTpl
func (e *BusDocTpl) Remove(d *dto.BusDocTplDeleteReq, p *actions.DataPermission) error {
	var data models.BusDocTpl

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBusDocTpl error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
