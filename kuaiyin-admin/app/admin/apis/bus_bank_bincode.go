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

type BusBankBincode struct {
	api.Api
}

// GetPage 获取银行主数据列表
// @Summary 获取银行主数据列表
// @Description 获取银行主数据列表
// @Tags 银行主数据
// @Param bankCode query string false "银行编码"
// @Param bankKh query string false "银行前缀卡号"
// @Param bankName query string false "银行名称"
// @Param cardType query string false "卡类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BusBankBincode}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-bank-bincode [get]
// @Security Bearer
func (e BusBankBincode) GetPage(c *gin.Context) {
    req := dto.BusBankBincodeGetPageReq{}
    s := service.BusBankBincode{}
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
	list := make([]models.BusBankBincode, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取银行主数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取银行主数据
// @Summary 获取银行主数据
// @Description 获取银行主数据
// @Tags 银行主数据
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BusBankBincode} "{"code": 200, "data": [...]}"
// @Router /api/v1/bus-bank-bincode/{id} [get]
// @Security Bearer
func (e BusBankBincode) Get(c *gin.Context) {
	req := dto.BusBankBincodeGetReq{}
	s := service.BusBankBincode{}
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
	var object models.BusBankBincode

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取银行主数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建银行主数据
// @Summary 创建银行主数据
// @Description 创建银行主数据
// @Tags 银行主数据
// @Accept application/json
// @Product application/json
// @Param data body dto.BusBankBincodeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bus-bank-bincode [post]
// @Security Bearer
func (e BusBankBincode) Insert(c *gin.Context) {
    req := dto.BusBankBincodeInsertReq{}
    s := service.BusBankBincode{}
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
		e.Error(500, err, fmt.Sprintf("创建银行主数据失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改银行主数据
// @Summary 修改银行主数据
// @Description 修改银行主数据
// @Tags 银行主数据
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BusBankBincodeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bus-bank-bincode/{id} [put]
// @Security Bearer
func (e BusBankBincode) Update(c *gin.Context) {
    req := dto.BusBankBincodeUpdateReq{}
    s := service.BusBankBincode{}
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
		e.Error(500, err, fmt.Sprintf("修改银行主数据失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除银行主数据
// @Summary 删除银行主数据
// @Description 删除银行主数据
// @Tags 银行主数据
// @Param data body dto.BusBankBincodeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bus-bank-bincode [delete]
// @Security Bearer
func (e BusBankBincode) Delete(c *gin.Context) {
    s := service.BusBankBincode{}
    req := dto.BusBankBincodeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除银行主数据失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
