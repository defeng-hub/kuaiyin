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

type KyComment struct {
	api.Api
}

// GetPage 获取KyComment列表
// @Summary 获取KyComment列表
// @Description 获取KyComment列表
// @Tags KyComment
// @Param videoId query int64 false "视频id"
// @Param userId query int64 false "用户id"
// @Param msg query string false "评论内容"
// @Param state query int64 false "评论状态（01正常 2有风险）"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.KyComment}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-comment [get]
// @Security Bearer
func (e KyComment) GetPage(c *gin.Context) {
    req := dto.KyCommentGetPageReq{}
    s := service.KyComment{}
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
	list := make([]models.KyComment, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyComment失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取KyComment
// @Summary 获取KyComment
// @Description 获取KyComment
// @Tags KyComment
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.KyComment} "{"code": 200, "data": [...]}"
// @Router /api/v1/ky-comment/{id} [get]
// @Security Bearer
func (e KyComment) Get(c *gin.Context) {
	req := dto.KyCommentGetReq{}
	s := service.KyComment{}
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
	var object models.KyComment

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取KyComment失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建KyComment
// @Summary 创建KyComment
// @Description 创建KyComment
// @Tags KyComment
// @Accept application/json
// @Product application/json
// @Param data body dto.KyCommentInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ky-comment [post]
// @Security Bearer
func (e KyComment) Insert(c *gin.Context) {
    req := dto.KyCommentInsertReq{}
    s := service.KyComment{}
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
		e.Error(500, err, fmt.Sprintf("创建KyComment失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改KyComment
// @Summary 修改KyComment
// @Description 修改KyComment
// @Tags KyComment
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.KyCommentUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ky-comment/{id} [put]
// @Security Bearer
func (e KyComment) Update(c *gin.Context) {
    req := dto.KyCommentUpdateReq{}
    s := service.KyComment{}
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
		e.Error(500, err, fmt.Sprintf("修改KyComment失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除KyComment
// @Summary 删除KyComment
// @Description 删除KyComment
// @Tags KyComment
// @Param data body dto.KyCommentDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ky-comment [delete]
// @Security Bearer
func (e KyComment) Delete(c *gin.Context) {
    s := service.KyComment{}
    req := dto.KyCommentDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除KyComment失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
