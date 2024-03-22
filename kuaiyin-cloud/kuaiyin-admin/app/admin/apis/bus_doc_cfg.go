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

type BusDocCfg struct {
	api.Api
}

// GetPage 获取文书设置列表
// @Summary 获取文书设置列表
// @Description 获取文书设置列表
// @Tags 文书设置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusDocCfg}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-cfg [get]
// @Security Bearer
func (e BusDocCfg) GetPage(c *gin.Context) {
	req := dto.BusDocCfgGetPageReq{}
	s := service.BusDocCfg{}
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
	list := make([]models.BusDocCfg, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文书设置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取文书设置
// @Summary 获取文书设置
// @Description 获取文书设置
// @Tags 文书设置
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusDocCfg} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-cfg/{id} [get]
// @Security Bearer
func (e BusDocCfg) Get(c *gin.Context) {
	req := dto.BusDocCfgGetReq{}
	s := service.BusDocCfg{}
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
	var object models.BusDocCfg

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文书设置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建文书设置
// @Summary 创建文书设置
// @Description 创建文书设置
// @Tags 文书设置
// @Accept application/json
// @Product application/json
// @Param data body dto.BusDocCfgInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-doc-cfg [post]
// @Security Bearer
func (e BusDocCfg) Insert(c *gin.Context) {
	req := dto.BusDocCfgInsertReq{}
	s := service.BusDocCfg{}
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
		e.Error(500, err, fmt.Sprintf("创建文书设置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改文书设置
// @Summary 修改文书设置
// @Description 修改文书设置
// @Tags 文书设置
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusDocCfgUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-doc-cfg/{id} [put]
// @Security Bearer
func (e BusDocCfg) Update(c *gin.Context) {
	req := dto.BusDocCfgUpdateReq{}
	s := service.BusDocCfg{}
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
		e.Error(500, err, fmt.Sprintf("修改文书设置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除文书设置
// @Summary 删除文书设置
// @Description 删除文书设置
// @Tags 文书设置
// @Param data body dto.BusDocCfgDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-doc-cfg [delete]
// @Security Bearer
func (e BusDocCfg) Delete(c *gin.Context) {
	s := service.BusDocCfg{}
	req := dto.BusDocCfgDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除文书设置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
