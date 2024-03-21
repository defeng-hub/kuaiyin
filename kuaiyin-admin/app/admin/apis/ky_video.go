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

type KyVideo struct {
	api.Api
}

// GetPage 获取KyVideo列表
// @Summary 获取KyVideo列表
// @Description 获取KyVideo列表
// @Tags KyVideo
// @Param title query string false "标题"
// @Param userId query int64 false "用户id"
// @Param username query string false "用户名"
// @Param msg query string false "描述"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.KyVideo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-video [get]
// @Security Bearer
func (e KyVideo) GetPage(c *gin.Context) {
    req := dto.KyVideoGetPageReq{}
    s := service.KyVideo{}
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
	list := make([]models.KyVideo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyVideo失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取KyVideo
// @Summary 获取KyVideo
// @Description 获取KyVideo
// @Tags KyVideo
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.KyVideo} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-video/{id} [get]
// @Security Bearer
func (e KyVideo) Get(c *gin.Context) {
	req := dto.KyVideoGetReq{}
	s := service.KyVideo{}
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
	var object models.KyVideo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyVideo失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建KyVideo
// @Summary 创建KyVideo
// @Description 创建KyVideo
// @Tags KyVideo
// @Accept application/json
// @Product application/json
// @Param data body dto.KyVideoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ky-video [post]
// @Security Bearer
func (e KyVideo) Insert(c *gin.Context) {
    req := dto.KyVideoInsertReq{}
    s := service.KyVideo{}
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
		e.Error(500, err, fmt.Sprintf("创建KyVideo失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改KyVideo
// @Summary 修改KyVideo
// @Description 修改KyVideo
// @Tags KyVideo
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.KyVideoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ky-video/{id} [put]
// @Security Bearer
func (e KyVideo) Update(c *gin.Context) {
    req := dto.KyVideoUpdateReq{}
    s := service.KyVideo{}
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
		e.Error(500, err, fmt.Sprintf("修改KyVideo失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除KyVideo
// @Summary 删除KyVideo
// @Description 删除KyVideo
// @Tags KyVideo
// @Param data body dto.KyVideoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ky-video [delete]
// @Security Bearer
func (e KyVideo) Delete(c *gin.Context) {
    s := service.KyVideo{}
    req := dto.KyVideoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除KyVideo失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
