package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type TbSyncLog struct {
	api.Api
}

// GetPage 获取日志同步列表
// @Summary 获取日志同步列表
// @Description 获取日志同步列表
// @Tags 日志同步
// @Param env query string false "环境"
// @Param logName query string false "文件名"
// @Param status query int64 false "同步状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbSyncLog}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-sync-log [get]
// @Security Bearer
func (e TbSyncLog) GetPage(c *gin.Context) {
    req := dto.TbSyncLogGetPageReq{}
    s := service.TbSyncLog{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.TbSyncLog, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取日志同步失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取日志同步
// @Summary 获取日志同步
// @Description 获取日志同步
// @Tags 日志同步
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbSyncLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-sync-log/{id} [get]
// @Security Bearer
func (e TbSyncLog) Get(c *gin.Context) {
	req := dto.TbSyncLogGetReq{}
	s := service.TbSyncLog{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.TbSyncLog

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取日志同步失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建日志同步
// @Summary 创建日志同步
// @Description 创建日志同步
// @Tags 日志同步
// @Accept application/json
// @Product application/json
// @Param data body dto.TbSyncLogInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-sync-log [post]
// @Security Bearer
func (e TbSyncLog) Insert(c *gin.Context) {
    req := dto.TbSyncLogInsertReq{}
    s := service.TbSyncLog{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建日志同步失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改日志同步
// @Summary 修改日志同步
// @Description 修改日志同步
// @Tags 日志同步
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbSyncLogUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-sync-log/{id} [put]
// @Security Bearer
func (e TbSyncLog) Update(c *gin.Context) {
    req := dto.TbSyncLogUpdateReq{}
    s := service.TbSyncLog{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改日志同步失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除日志同步
// @Summary 删除日志同步
// @Description 删除日志同步
// @Tags 日志同步
// @Param data body dto.TbSyncLogDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-sync-log [delete]
// @Security Bearer
func (e TbSyncLog) Delete(c *gin.Context) {
    s := service.TbSyncLog{}
    req := dto.TbSyncLogDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除日志同步失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
