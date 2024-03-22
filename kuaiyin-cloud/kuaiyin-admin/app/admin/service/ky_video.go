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

type KyVideo struct {
	service.Service
}

// GetPage 获取KyVideo列表
func (e *KyVideo) GetPage(c *dto.KyVideoGetPageReq, p *actions.DataPermission, list *[]models.KyVideo, count *int64) error {
	var err error
	var data models.KyVideo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("KyVideoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取KyVideo对象
func (e *KyVideo) Get(d *dto.KyVideoGetReq, p *actions.DataPermission, model *models.KyVideo) error {
	var data models.KyVideo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetKyVideo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建KyVideo对象
func (e *KyVideo) Insert(c *dto.KyVideoInsertReq) error {
    var err error
    var data models.KyVideo
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("KyVideoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改KyVideo对象
func (e *KyVideo) Update(c *dto.KyVideoUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.KyVideo{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("KyVideoService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除KyVideo
func (e *KyVideo) Remove(d *dto.KyVideoDeleteReq, p *actions.DataPermission) error {
	var data models.KyVideo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveKyVideo error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
