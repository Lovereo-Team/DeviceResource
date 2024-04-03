package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/admin/service/system"
	"DeviceResource/core"
	"DeviceResource/core/request"
	"DeviceResource/core/response"
	"DeviceResource/middleware"
	"DeviceResource/util"
	"github.com/gin-gonic/gin"
)

var ResourceGroup = core.Group("/", newResourceHandler, regResource, middleware.TokenAuth())

func newResourceHandler(srv system.IResourceService) *resourceHandler {
	return &resourceHandler{srv: srv}
}

func regResource(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *resourceHandler) {
		rg.GET("/resource/list", handle.list)
		rg.GET("/resource/detail", handle.detail)
		rg.POST("/resource/add", handle.add)
		rg.POST("/resource/edit", handle.edit)
		rg.POST("/resource/del", handle.del)
	})
}

type resourceHandler struct {
	srv system.IResourceService
}

// list resource列表
func (hd resourceHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.ResourceListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := hd.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail resource详情
func (hd resourceHandler) detail(c *gin.Context) {
	var detailReq req.ResourceDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := hd.srv.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// add resource新增
func (hd resourceHandler) add(c *gin.Context) {
	var addReq req.ResourceAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Add(addReq))
}

// edit resource编辑
func (hd resourceHandler) edit(c *gin.Context) {
	var editReq req.ResourceEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Edit(editReq))
}

// del resource删除
func (hd resourceHandler) del(c *gin.Context) {
	var delReq req.ResourceDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Del(delReq.Id))
}
