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

type KyLike struct {
	api.Api
}

// GetPage 获取KyLike列表
// @Summary 获取KyLike列表
// @Description 获取KyLike列表
// @Tags KyLike
// @Param userId query int64 false "用户id"
// @Param videoId query int64 false "视频id"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.KyLike}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-like [get]
// @Security Bearer
func (e KyLike) GetPage(c *gin.Context) {
    req := dto.KyLikeGetPageReq{}
    s := service.KyLike{}
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
	list := make([]models.KyLike, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyLike失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取KyLike
// @Summary 获取KyLike
// @Description 获取KyLike
// @Tags KyLike
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.KyLike} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-like/{id} [get]
// @Security Bearer
func (e KyLike) Get(c *gin.Context) {
	req := dto.KyLikeGetReq{}
	s := service.KyLike{}
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
	var object models.KyLike

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyLike失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建KyLike
// @Summary 创建KyLike
// @Description 创建KyLike
// @Tags KyLike
// @Accept application/json
// @Product application/json
// @Param data body dto.KyLikeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ky-like [post]
// @Security Bearer
func (e KyLike) Insert(c *gin.Context) {
    req := dto.KyLikeInsertReq{}
    s := service.KyLike{}
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
		e.Error(500, err, fmt.Sprintf("创建KyLike失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改KyLike
// @Summary 修改KyLike
// @Description 修改KyLike
// @Tags KyLike
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.KyLikeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ky-like/{id} [put]
// @Security Bearer
func (e KyLike) Update(c *gin.Context) {
    req := dto.KyLikeUpdateReq{}
    s := service.KyLike{}
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
		e.Error(500, err, fmt.Sprintf("修改KyLike失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除KyLike
// @Summary 删除KyLike
// @Description 删除KyLike
// @Tags KyLike
// @Param data body dto.KyLikeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ky-like [delete]
// @Security Bearer
func (e KyLike) Delete(c *gin.Context) {
    s := service.KyLike{}
    req := dto.KyLikeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除KyLike失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
