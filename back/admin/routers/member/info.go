package member

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/admin/service/member"
	"DeviceResource/core"
	"DeviceResource/core/response"
	"DeviceResource/middleware"
	"DeviceResource/util"
	"github.com/gin-gonic/gin"
)

var InfoGroup = core.Group("/member", newInfoHandler, regInfo, middleware.TokenAuth())

func newInfoHandler(srv member.IMemberInfoService) *infoHandler {
	return &infoHandler{srv: srv}
}

func regInfo(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *infoHandler) {
		rg.GET("/info", handle.list)
		rg.POST("/update", handle.update)
	})
}

type infoHandler struct {
	srv member.IMemberInfoService
}

func (hd infoHandler) list(c *gin.Context) {
	var listReq req.MemberInfoListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := hd.srv.List(listReq)
	response.CheckAndRespWithData(c, res, err)
}

func (hd infoHandler) update(c *gin.Context) {
	var editReq req.MemberInfoEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Edit(editReq))
}
