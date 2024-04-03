package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/admin/service/system"
	"DeviceResource/core"
	"DeviceResource/core/request"
	"DeviceResource/core/response"
	"DeviceResource/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"likeadmin/util"
)

var MemberGroup = core.Group("/", newMemberHandler, regMember, middleware.TokenAuth())

func newMemberHandler(srv system.IMemberService) *memberHandler {
	return &memberHandler{srv: srv}
}

func regMember(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *memberHandler) {
		rg.GET("/member/list", handle.list)
		rg.GET("/member/detail", handle.detail)
		rg.POST("/member/add", handle.add)
		rg.POST("/member/edit", handle.edit)
		rg.POST("/member/del", handle.del)
	})
}

type memberHandler struct {
	srv system.IMemberService
}

// list member列表
func (hd memberHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.MemberListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := hd.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail member详情
func (hd memberHandler) detail(c *gin.Context) {
	var detailReq req.MemberDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := hd.srv.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// add member新增
func (hd memberHandler) add(c *gin.Context) {
	var addReq req.MemberAddReq
	fmt.Println("aaa", addReq)
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Add(addReq))
}

// edit member编辑
func (hd memberHandler) edit(c *gin.Context) {
	var editReq req.MemberEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Edit(editReq))
}

// del member删除
func (hd memberHandler) del(c *gin.Context) {
	var delReq req.MemberDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Del(delReq.Id))
}
