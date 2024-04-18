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

type Videos struct {
	api.Api
}

// GetPage 获取Videos列表
// @Summary 获取Videos列表
// @Description 获取Videos列表
// @Tags Videos
// @Param authorId query int64 false "作者ID"
// @Param playUrl query string false "视频地址"
// @Param coverUrl query string false "图片地址"
// @Param title query string false "标题"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Videos}} "{"code": 200, "data": [...]}"
// @Router /api/v1/videos [get]
// @Security Bearer
func (e Videos) GetPage(c *gin.Context) {
    req := dto.VideosGetPageReq{}
    s := service.Videos{}
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
	list := make([]models.Videos, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Videos失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Videos
// @Summary 获取Videos
// @Description 获取Videos
// @Tags Videos
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Videos} "{"code": 200, "data": [...]}"
// @Router /api/v1/videos/{id} [get]
// @Security Bearer
func (e Videos) Get(c *gin.Context) {
	req := dto.VideosGetReq{}
	s := service.Videos{}
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
	var object models.Videos

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Videos失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Videos
// @Summary 创建Videos
// @Description 创建Videos
// @Tags Videos
// @Accept application/json
// @Product application/json
// @Param data body dto.VideosInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/videos [post]
// @Security Bearer
func (e Videos) Insert(c *gin.Context) {
    req := dto.VideosInsertReq{}
    s := service.Videos{}
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
		e.Error(500, err, fmt.Sprintf("创建Videos失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Videos
// @Summary 修改Videos
// @Description 修改Videos
// @Tags Videos
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.VideosUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/videos/{id} [put]
// @Security Bearer
func (e Videos) Update(c *gin.Context) {
    req := dto.VideosUpdateReq{}
    s := service.Videos{}
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
		e.Error(500, err, fmt.Sprintf("修改Videos失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Videos
// @Summary 删除Videos
// @Description 删除Videos
// @Tags Videos
// @Param data body dto.VideosDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/videos [delete]
// @Security Bearer
func (e Videos) Delete(c *gin.Context) {
    s := service.Videos{}
    req := dto.VideosDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Videos失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
