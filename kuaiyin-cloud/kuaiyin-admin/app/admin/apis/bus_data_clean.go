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

type BusDataClean struct {
	api.Api
}

// GetPage 获取数据清洗列表
// @Summary 获取数据清洗列表
// @Description 获取数据清洗列表
// @Tags 数据清洗
// @Param taskName query string false "任务名称"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusDataClean}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-data-clean [get]
// @Security Bearer
func (e BusDataClean) GetPage(c *gin.Context) {
    req := dto.BusDataCleanGetPageReq{}
    s := service.BusDataClean{}
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
	list := make([]models.BusDataClean, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取数据清洗失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取数据清洗
// @Summary 获取数据清洗
// @Description 获取数据清洗
// @Tags 数据清洗
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusDataClean} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-data-clean/{id} [get]
// @Security Bearer
func (e BusDataClean) Get(c *gin.Context) {
	req := dto.BusDataCleanGetReq{}
	s := service.BusDataClean{}
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
	var object models.BusDataClean

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取数据清洗失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建数据清洗
// @Summary 创建数据清洗
// @Description 创建数据清洗
// @Tags 数据清洗
// @Accept application/json
// @Product application/json
// @Param data body dto.BusDataCleanInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-data-clean [post]
// @Security Bearer
func (e BusDataClean) Insert(c *gin.Context) {
    req := dto.BusDataCleanInsertReq{}
    s := service.BusDataClean{}
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
		e.Error(500, err, fmt.Sprintf("创建数据清洗失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改数据清洗
// @Summary 修改数据清洗
// @Description 修改数据清洗
// @Tags 数据清洗
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusDataCleanUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-data-clean/{id} [put]
// @Security Bearer
func (e BusDataClean) Update(c *gin.Context) {
    req := dto.BusDataCleanUpdateReq{}
    s := service.BusDataClean{}
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
		e.Error(500, err, fmt.Sprintf("修改数据清洗失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除数据清洗
// @Summary 删除数据清洗
// @Description 删除数据清洗
// @Tags 数据清洗
// @Param data body dto.BusDataCleanDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-data-clean [delete]
// @Security Bearer
func (e BusDataClean) Delete(c *gin.Context) {
    s := service.BusDataClean{}
    req := dto.BusDataCleanDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除数据清洗失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
