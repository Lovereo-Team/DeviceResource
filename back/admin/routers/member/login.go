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

var LoginGroup = core.Group("/member", newLoginHandler, regLogin, middleware.TokenAuth())

func newLoginHandler(srv member.IMemberLoginService) *loginHandler {
	return &loginHandler{srv: srv}
}

func regLogin(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *loginHandler) {
		rg.POST("/login", handle.login)
		rg.POST("/logout", handle.logout)
	})
}

type loginHandler struct {
	srv member.IMemberLoginService
}

// login 登录系统
func (lh loginHandler) login(c *gin.Context) {
	var loginReq req.MemberLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := lh.srv.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

// logout 登录退出
func (lh loginHandler) logout(c *gin.Context) {
	var logoutReq req.MemberLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndResp(c, lh.srv.Logout(&logoutReq))
}
