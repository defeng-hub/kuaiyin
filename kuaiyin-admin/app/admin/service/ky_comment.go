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

type KyComment struct {
	service.Service
}

// GetPage 获取KyComment列表
func (e *KyComment) GetPage(c *dto.KyCommentGetPageReq, p *actions.DataPermission, list *[]models.KyComment, count *int64) error {
	var err error
	var data models.KyComment

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("KyCommentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取KyComment对象
func (e *KyComment) Get(d *dto.KyCommentGetReq, p *actions.DataPermission, model *models.KyComment) error {
	var data models.KyComment

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetKyComment error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建KyComment对象
func (e *KyComment) Insert(c *dto.KyCommentInsertReq) error {
    var err error
    var data models.KyComment
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("KyCommentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改KyComment对象
func (e *KyComment) Update(c *dto.KyCommentUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.KyComment{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("KyCommentService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除KyComment
func (e *KyComment) Remove(d *dto.KyCommentDeleteReq, p *actions.DataPermission) error {
	var data models.KyComment

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveKyComment error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
