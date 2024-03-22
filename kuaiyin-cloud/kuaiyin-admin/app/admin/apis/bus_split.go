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

type BusSplit struct {
	api.Api
}

// GetPage 获取数据文件拆分列表
// @Summary 获取数据文件拆分列表
// @Description 获取数据文件拆分列表
// @Tags 数据文件拆分
// @Param taskName query string false "任务名称"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusSplit}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-split [get]
// @Security Bearer
func (e BusSplit) GetPage(c *gin.Context) {
	req := dto.BusSplitGetPageReq{}
	s := service.BusSplit{}
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
	list := make([]models.BusSplit, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取数据文件拆分失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取数据文件拆分
// @Summary 获取数据文件拆分
// @Description 获取数据文件拆分
// @Tags 数据文件拆分
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusSplit} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-split/{id} [get]
// @Security Bearer
func (e BusSplit) Get(c *gin.Context) {
	req := dto.BusSplitGetReq{}
	s := service.BusSplit{}
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
	var object models.BusSplit

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取数据文件拆分失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建数据文件拆分
// @Summary 创建数据文件拆分
// @Description 创建数据文件拆分
// @Tags 数据文件拆分
// @Accept application/json
// @Product application/json
// @Param data body dto.BusSplitInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-split [post]
// @Security Bearer
func (e BusSplit) Insert(c *gin.Context) {
	req := dto.BusSplitInsertReq{}
	s := service.BusSplit{}
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
		e.Error(500, err, fmt.Sprintf("创建数据文件拆分失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改数据文件拆分
// @Summary 修改数据文件拆分
// @Description 修改数据文件拆分
// @Tags 数据文件拆分
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusSplitUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-split/{id} [put]
// @Security Bearer
func (e BusSplit) Update(c *gin.Context) {
	req := dto.BusSplitUpdateReq{}
	s := service.BusSplit{}
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
		e.Error(500, err, fmt.Sprintf("修改数据文件拆分失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除数据文件拆分
// @Summary 删除数据文件拆分
// @Description 删除数据文件拆分
// @Tags 数据文件拆分
// @Param data body dto.BusSplitDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-split [delete]
// @Security Bearer
func (e BusSplit) Delete(c *gin.Context) {
	s := service.BusSplit{}
	req := dto.BusSplitDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除数据文件拆分失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
