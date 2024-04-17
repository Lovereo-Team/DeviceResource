package member

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/admin/schemas/resp"
	"DeviceResource/config"
	"DeviceResource/core"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"runtime/debug"
	"strconv"
	"time"
)

type IMemberLoginService interface {
	Login(c *gin.Context, req *req.MemberLoginReq) (res resp.MemberLoginResp, e error)
	Logout(req *req.MemberLogoutReq) (e error)
	RecordLoginLog(c *gin.Context, adminId int, username string, errStr string) (e error)
}

// NewSystemLoginService 初始化
func NewMemberLoginService(db *gorm.DB, memberSrv IMemberAuthAdminService) IMemberLoginService {
	return &memberLoginService{db: db, memberSrv: memberSrv}
}

// systemLoginService 系统登录服务实现类
type memberLoginService struct {
	db        *gorm.DB
	memberSrv IMemberAuthAdminService
}

// Login 登录
func (loginSrv memberLoginService) Login(c *gin.Context, req *req.MemberLoginReq) (res resp.MemberLoginResp, e error) {
	member, err := loginSrv.memberSrv.FindByUsername(req.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	} else if err != nil {
		core.Logger.Errorf("Login FindByUsername err: err=[%+v]", err)
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.Failed.Msg()); e != nil {
			return
		}
		e = response.Failed
		return
	}
	if member.Status == 2 {
		if e = loginSrv.RecordLoginLog(c, member.Id, req.Username, response.LoginDisableError.Msg()); e != nil {
			return
		}
		e = response.LoginDisableError
		return
	}
	md5Pwd := util.ToolsUtil.MakeMd5(req.Password)
	if member.Password != md5Pwd {
		if e = loginSrv.RecordLoginLog(c, member.Id, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			// 自定义类型
			case response.RespType:
				panic(r)
			// 其他类型
			default:
				core.Logger.Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
				loginSrv.RecordLoginLog(c, member.Id, req.Username, response.Failed.Msg())
				panic(response.Failed)
			}
		}
	}()
	token := util.ToolsUtil.MakeToken()
	memberIdStr := strconv.FormatUint(uint64(member.Id), 10)

	// 缓存登录信息
	util.RedisUtil.Set(config.MemberConfig.BackstageTokenKey+token, memberIdStr, 7200)
	loginSrv.memberSrv.CacheAdminUserByUid(member.Id)

	// 更新登录信息
	err = loginSrv.db.Model(&member).Updates(
		util.Member{UpdateTime: time.Now().Unix()}).Error
	if err != nil {
		if e = loginSrv.RecordLoginLog(c, member.Id, req.Username, response.SystemError.Msg()); e != nil {
			return
		}
		if e = response.CheckErr(err, "Login Updates err"); e != nil {
			return
		}
	}
	// 记录登录日志
	if e = loginSrv.RecordLoginLog(c, member.Id, req.Username, ""); e != nil {
		return
	}
	// 返回登录信息
	return resp.MemberLoginResp{Token: token}, nil
}

// Logout 退出
func (loginSrv memberLoginService) Logout(req *req.MemberLogoutReq) (e error) {
	util.RedisUtil.Del(config.MemberConfig.BackstageTokenKey + req.Token)
	return
}

// RecordLoginLog 记录登录日志
func (loginSrv memberLoginService) RecordLoginLog(c *gin.Context, adminId int, username string, errStr string) (e error) {
	ua := core.UAParser.Parse(c.GetHeader("user-agent"))
	var status uint8
	if errStr == "" {
		status = 1
	}
	err := loginSrv.db.Create(&util.MemberLogLogin{
		MemberId: adminId, Username: username, Ip: c.ClientIP(), Os: ua.Os.Family,
		Browser: ua.UserAgent.Family, Status: status}).Error
	e = response.CheckErr(err, "RecordLoginLog Create err")
	return
}
