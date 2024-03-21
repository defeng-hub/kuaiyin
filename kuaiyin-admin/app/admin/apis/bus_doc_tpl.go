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

type BusDocTpl struct {
	api.Api
}

// GetPage 获取文书模板列表
// @Summary 获取文书模板列表
// @Description 获取文书模板列表
// @Tags 文书模板
// @Param unitName query string false "单位名称"
// @Param applyName query string false "申请模板名称"
// @Param announceName query string false "通知模板名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusDocTpl}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-tpl [get]
// @Security Bearer
func (e BusDocTpl) GetPage(c *gin.Context) {
    req := dto.BusDocTplGetPageReq{}
    s := service.BusDocTpl{}
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
	list := make([]models.BusDocTpl, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文书模板失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取文书模板
// @Summary 获取文书模板
// @Description 获取文书模板
// @Tags 文书模板
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusDocTpl} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-tpl/{id} [get]
// @Security Bearer
func (e BusDocTpl) Get(c *gin.Context) {
	req := dto.BusDocTplGetReq{}
	s := service.BusDocTpl{}
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
	var object models.BusDocTpl

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文书模板失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建文书模板
// @Summary 创建文书模板
// @Description 创建文书模板
// @Tags 文书模板
// @Accept application/json
// @Product application/json
// @Param data body dto.BusDocTplInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-doc-tpl [post]
// @Security Bearer
func (e BusDocTpl) Insert(c *gin.Context) {
    req := dto.BusDocTplInsertReq{}
    s := service.BusDocTpl{}
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
		e.Error(500, err, fmt.Sprintf("创建文书模板失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改文书模板
// @Summary 修改文书模板
// @Description 修改文书模板
// @Tags 文书模板
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusDocTplUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-doc-tpl/{id} [put]
// @Security Bearer
func (e BusDocTpl) Update(c *gin.Context) {
    req := dto.BusDocTplUpdateReq{}
    s := service.BusDocTpl{}
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
		e.Error(500, err, fmt.Sprintf("修改文书模板失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除文书模板
// @Summary 删除文书模板
// @Description 删除文书模板
// @Tags 文书模板
// @Param data body dto.BusDocTplDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-doc-tpl [delete]
// @Security Bearer
func (e BusDocTpl) Delete(c *gin.Context) {
    s := service.BusDocTpl{}
    req := dto.BusDocTplDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除文书模板失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
