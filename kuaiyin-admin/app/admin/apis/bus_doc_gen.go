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

type BusDocGen struct {
	api.Api
}

// GetPage 获取文书生成列表
// @Summary 获取文书生成列表
// @Description 获取文书生成列表
// @Tags 文书生成
// @Param caseId query string false "案件ID"
// @Param unitId query string false "单位ID"
// @Param reviewCode query string false "审核员编码"
// @Param status query string false "审批状态"
// @Param procStatus query string false "处理状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusDocGen}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-gen [get]
// @Security Bearer
func (e BusDocGen) GetPage(c *gin.Context) {
    req := dto.BusDocGenGetPageReq{}
    s := service.BusDocGen{}
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
	list := make([]models.BusDocGen, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文书生成失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取文书生成
// @Summary 获取文书生成
// @Description 获取文书生成
// @Tags 文书生成
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusDocGen} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-doc-gen/{id} [get]
// @Security Bearer
func (e BusDocGen) Get(c *gin.Context) {
	req := dto.BusDocGenGetReq{}
	s := service.BusDocGen{}
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
	var object models.BusDocGen

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文书生成失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建文书生成
// @Summary 创建文书生成
// @Description 创建文书生成
// @Tags 文书生成
// @Accept application/json
// @Product application/json
// @Param data body dto.BusDocGenInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-doc-gen [post]
// @Security Bearer
func (e BusDocGen) Insert(c *gin.Context) {
    req := dto.BusDocGenInsertReq{}
    s := service.BusDocGen{}
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
		e.Error(500, err, fmt.Sprintf("创建文书生成失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改文书生成
// @Summary 修改文书生成
// @Description 修改文书生成
// @Tags 文书生成
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusDocGenUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-doc-gen/{id} [put]
// @Security Bearer
func (e BusDocGen) Update(c *gin.Context) {
    req := dto.BusDocGenUpdateReq{}
    s := service.BusDocGen{}
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
		e.Error(500, err, fmt.Sprintf("修改文书生成失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除文书生成
// @Summary 删除文书生成
// @Description 删除文书生成
// @Tags 文书生成
// @Param data body dto.BusDocGenDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-doc-gen [delete]
// @Security Bearer
func (e BusDocGen) Delete(c *gin.Context) {
    s := service.BusDocGen{}
    req := dto.BusDocGenDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除文书生成失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
