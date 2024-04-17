package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/admin/service/system"
	"DeviceResource/core"
	"DeviceResource/core/response"
	"DeviceResource/middleware"
	"DeviceResource/util"
	"github.com/gin-gonic/gin"
)

var AlbumGroup = core.Group("/", newAlbumHandler, regAlbum, middleware.TokenAuth())

func newAlbumHandler(srv system.IAlbumService) *albumHandler {
	return &albumHandler{srv: srv}
}

func regAlbum(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *albumHandler) {
		rg.POST("/pic/add", handle.add)
	})
}

type albumHandler struct {
	srv system.IAlbumService
}

// add album新增
func (hd albumHandler) add(c *gin.Context) {
	var addReq req.AlbumAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Add(addReq))
}
