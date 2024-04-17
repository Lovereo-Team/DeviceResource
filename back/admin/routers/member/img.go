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

var ImgGroup = core.Group("/member", newImgHandler, regImg, middleware.TokenAuth())

func newImgHandler(srv member.IImgService) *imgHandler {
	return &imgHandler{srv: srv}
}

func regImg(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *imgHandler) {
		rg.GET("/img", handle.list)
		rg.GET("/getVideo", handle.getVideo)
	})
}

type imgHandler struct {
	srv member.IImgService
}

func (hd imgHandler) list(c *gin.Context) {
	var listReq req.ResourceListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := hd.srv.List(listReq)
	response.CheckAndRespWithData(c, res, err)
}

func (hd imgHandler) getVideo(c *gin.Context) {
	var getReq req.ResourceListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &getReq)) {
		return
	}
	res, err := hd.srv.GetVideo(getReq)
	response.CheckAndRespWithData(c, res, err)
}
