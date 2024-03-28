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

type KyFollower struct {
	api.Api
}

// GetPage 获取KyFollower列表
// @Summary 获取KyFollower列表
// @Description 获取KyFollower列表
// @Tags KyFollower
// @Param userId query int64 false "用户id"
// @Param followerId query int64 false "关注id"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.KyFollower}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-follower [get]
// @Security Bearer
func (e KyFollower) GetPage(c *gin.Context) {
    req := dto.KyFollowerGetPageReq{}
    s := service.KyFollower{}
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
	list := make([]models.KyFollower, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyFollower失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取KyFollower
// @Summary 获取KyFollower
// @Description 获取KyFollower
// @Tags KyFollower
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.KyFollower} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-follower/{id} [get]
// @Security Bearer
func (e KyFollower) Get(c *gin.Context) {
	req := dto.KyFollowerGetReq{}
	s := service.KyFollower{}
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
	var object models.KyFollower

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyFollower失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建KyFollower
// @Summary 创建KyFollower
// @Description 创建KyFollower
// @Tags KyFollower
// @Accept application/json
// @Product application/json
// @Param data body dto.KyFollowerInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ky-follower [post]
// @Security Bearer
func (e KyFollower) Insert(c *gin.Context) {
    req := dto.KyFollowerInsertReq{}
    s := service.KyFollower{}
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
		e.Error(500, err, fmt.Sprintf("创建KyFollower失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改KyFollower
// @Summary 修改KyFollower
// @Description 修改KyFollower
// @Tags KyFollower
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.KyFollowerUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ky-follower/{id} [put]
// @Security Bearer
func (e KyFollower) Update(c *gin.Context) {
    req := dto.KyFollowerUpdateReq{}
    s := service.KyFollower{}
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
		e.Error(500, err, fmt.Sprintf("修改KyFollower失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除KyFollower
// @Summary 删除KyFollower
// @Description 删除KyFollower
// @Tags KyFollower
// @Param data body dto.KyFollowerDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ky-follower [delete]
// @Security Bearer
func (e KyFollower) Delete(c *gin.Context) {
    s := service.KyFollower{}
    req := dto.KyFollowerDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除KyFollower失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
