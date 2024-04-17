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

type KyLike struct {
	service.Service
}

// GetPage 获取KyLike列表
func (e *KyLike) GetPage(c *dto.KyLikeGetPageReq, p *actions.DataPermission, list *[]models.KyLike, count *int64) error {
	var err error
	var data models.KyLike

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("KyLikeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取KyLike对象
func (e *KyLike) Get(d *dto.KyLikeGetReq, p *actions.DataPermission, model *models.KyLike) error {
	var data models.KyLike

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetKyLike error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建KyLike对象
func (e *KyLike) Insert(c *dto.KyLikeInsertReq) error {
    var err error
    var data models.KyLike
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("KyLikeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改KyLike对象
func (e *KyLike) Update(c *dto.KyLikeUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.KyLike{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("KyLikeService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除KyLike
func (e *KyLike) Remove(d *dto.KyLikeDeleteReq, p *actions.DataPermission) error {
	var data models.KyLike

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveKyLike error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
