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

type KyFollower struct {
	service.Service
}

// GetPage 获取KyFollower列表
func (e *KyFollower) GetPage(c *dto.KyFollowerGetPageReq, p *actions.DataPermission, list *[]models.KyFollower, count *int64) error {
	var err error
	var data models.KyFollower

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("KyFollowerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取KyFollower对象
func (e *KyFollower) Get(d *dto.KyFollowerGetReq, p *actions.DataPermission, model *models.KyFollower) error {
	var data models.KyFollower

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetKyFollower error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建KyFollower对象
func (e *KyFollower) Insert(c *dto.KyFollowerInsertReq) error {
    var err error
    var data models.KyFollower
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("KyFollowerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改KyFollower对象
func (e *KyFollower) Update(c *dto.KyFollowerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.KyFollower{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("KyFollowerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除KyFollower
func (e *KyFollower) Remove(d *dto.KyFollowerDeleteReq, p *actions.DataPermission) error {
	var data models.KyFollower

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveKyFollower error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
