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

type BusDocCase struct {
	api.Api
}

// GetPage 获取案件信息列表
// @Summary 获取案件信息列表
// @Description 获取案件信息列表
// @Tags 案件信息
// @Param caseName query string false "案件名称"
// @Param caseCode query string false "案件编号"
// @Param caseType query string false "案件类别"
// @Param getType query string false "调取种类"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusDocCase}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-case [get]
// @Security Bearer
func (e BusDocCase) GetPage(c *gin.Context) {
    req := dto.BusDocCaseGetPageReq{}
    s := service.BusDocCase{}
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
	list := make([]models.BusDocCase, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取案件信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取案件信息
// @Summary 获取案件信息
// @Description 获取案件信息
// @Tags 案件信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusDocCase} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-case/{id} [get]
// @Security Bearer
func (e BusDocCase) Get(c *gin.Context) {
	req := dto.BusDocCaseGetReq{}
	s := service.BusDocCase{}
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
	var object models.BusDocCase

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取案件信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建案件信息
// @Summary 创建案件信息
// @Description 创建案件信息
// @Tags 案件信息
// @Accept application/json
// @Product application/json
// @Param data body dto.BusDocCaseInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-doc-case [post]
// @Security Bearer
func (e BusDocCase) Insert(c *gin.Context) {
    req := dto.BusDocCaseInsertReq{}
    s := service.BusDocCase{}
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
		e.Error(500, err, fmt.Sprintf("创建案件信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改案件信息
// @Summary 修改案件信息
// @Description 修改案件信息
// @Tags 案件信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusDocCaseUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-doc-case/{id} [put]
// @Security Bearer
func (e BusDocCase) Update(c *gin.Context) {
    req := dto.BusDocCaseUpdateReq{}
    s := service.BusDocCase{}
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
		e.Error(500, err, fmt.Sprintf("修改案件信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除案件信息
// @Summary 删除案件信息
// @Description 删除案件信息
// @Tags 案件信息
// @Param data body dto.BusDocCaseDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-doc-case [delete]
// @Security Bearer
func (e BusDocCase) Delete(c *gin.Context) {
    s := service.BusDocCase{}
    req := dto.BusDocCaseDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除案件信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
