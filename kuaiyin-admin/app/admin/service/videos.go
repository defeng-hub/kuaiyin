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

type Videos struct {
	service.Service
}

// GetPage 获取Videos列表
func (e *Videos) GetPage(c *dto.VideosGetPageReq, p *actions.DataPermission, list *[]models.Videos, count *int64) error {
	var err error
	var data models.Videos

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("VideosService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Videos对象
func (e *Videos) Get(d *dto.VideosGetReq, p *actions.DataPermission, model *models.Videos) error {
	var data models.Videos

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetVideos error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Videos对象
func (e *Videos) Insert(c *dto.VideosInsertReq) error {
    var err error
    var data models.Videos
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("VideosService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Videos对象
func (e *Videos) Update(c *dto.VideosUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Videos{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("VideosService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Videos
func (e *Videos) Remove(d *dto.VideosDeleteReq, p *actions.DataPermission) error {
	var data models.Videos

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveVideos error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
