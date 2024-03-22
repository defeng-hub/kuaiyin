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

type BusBankCardtypes struct {
	api.Api
}

// GetPage 获取银行详细数据列表
// @Summary 获取银行详细数据列表
// @Description 获取银行详细数据列表
// @Tags 银行详细数据
// @Param bankCode query string false "银行编码"
// @Param bankKh query string false "银行前缀卡号"
// @Param bankName query string false "银行名称"
// @Param cardName query string false "卡名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusBankCardtypes}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-bank-cardtypes [get]
// @Security Bearer
func (e BusBankCardtypes) GetPage(c *gin.Context) {
    req := dto.BusBankCardtypesGetPageReq{}
    s := service.BusBankCardtypes{}
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
	list := make([]models.BusBankCardtypes, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取银行详细数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取银行详细数据
// @Summary 获取银行详细数据
// @Description 获取银行详细数据
// @Tags 银行详细数据
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusBankCardtypes} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-bank-cardtypes/{id} [get]
// @Security Bearer
func (e BusBankCardtypes) Get(c *gin.Context) {
	req := dto.BusBankCardtypesGetReq{}
	s := service.BusBankCardtypes{}
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
	var object models.BusBankCardtypes

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取银行详细数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建银行详细数据
// @Summary 创建银行详细数据
// @Description 创建银行详细数据
// @Tags 银行详细数据
// @Accept application/json
// @Product application/json
// @Param data body dto.BusBankCardtypesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-bank-cardtypes [post]
// @Security Bearer
func (e BusBankCardtypes) Insert(c *gin.Context) {
    req := dto.BusBankCardtypesInsertReq{}
    s := service.BusBankCardtypes{}
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
		e.Error(500, err, fmt.Sprintf("创建银行详细数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改银行详细数据
// @Summary 修改银行详细数据
// @Description 修改银行详细数据
// @Tags 银行详细数据
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusBankCardtypesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-bank-cardtypes/{id} [put]
// @Security Bearer
func (e BusBankCardtypes) Update(c *gin.Context) {
    req := dto.BusBankCardtypesUpdateReq{}
    s := service.BusBankCardtypes{}
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
		e.Error(500, err, fmt.Sprintf("修改银行详细数据失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除银行详细数据
// @Summary 删除银行详细数据
// @Description 删除银行详细数据
// @Tags 银行详细数据
// @Param data body dto.BusBankCardtypesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-bank-cardtypes [delete]
// @Security Bearer
func (e BusBankCardtypes) Delete(c *gin.Context) {
    s := service.BusBankCardtypes{}
    req := dto.BusBankCardtypesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除银行详细数据失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
