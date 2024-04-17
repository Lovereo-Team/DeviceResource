package req

import (
	"DeviceResource/core"
)

type MemberLoginReq struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
}

type MemberLogoutReq struct {
	Token string `header:"token" binding:"required"` // 令牌
}

// MemberListReq 员工列列表参数
type MemberListReq struct {
	Username  string `form:"username"`  // 用户名
	Password  string `form:"password"`  // 密码
	RealName  string `form:"realName"`  // 真实姓名
	Email     string `form:"email"`     // 邮箱
	Status    int    `form:"status"`    // 用户状态 1：正常 2：禁用
	LoginTime int    `form:"loginTime"` // 登录时间
	Token     string `form:"token"`     // 用户凭证
}

// MemberDetailReq 员工列详情参数
type MemberDetailReq struct {
	Id uint `form:"id"` //
}

// MemberAddReq 员工列新增参数
type MemberAddReq struct {
	Username   string `form:"username"`   // 用户名
	Password   string `form:"password"`   // 密码
	RealName   string `form:"realName"`   // 真实姓名
	Email      string `form:"email"`      // 邮箱
	Status     int    `form:"status"`     // 用户状态 1：正常 2：禁用
	CreateTime int    `form:"createTime"` // 创建时间
	Token      string `form:"token"`      // 用户凭证
}

// MemberEditReq 员工列新增参数
type MemberEditReq struct {
	Id         int    `form:"id"`         //
	Username   string `form:"username"`   // 用户名
	Password   string `form:"password"`   // 密码
	RealName   string `form:"realName"`   // 真实姓名
	Email      string `form:"email"`      // 邮箱
	Status     int    `form:"status"`     // 用户状态 1：正常 2：禁用
	UpdateTime int    `form:"updateTime"` // 登录时间
	Token      string `form:"token"`      // 用户凭证
}

// MemberDelReq 员工列新增参数
type MemberDelReq struct {
	Id uint `form:"id"` //
}

// MemberResp 员工列返回信息
type MemberResp struct {
	Id         int         `json:"id" structs:"id"`                 //
	Username   string      `json:"username" structs:"username"`     // 用户名
	Password   string      `json:"password" structs:"password"`     // 密码
	RealName   string      `json:"realName" structs:"realName"`     // 真实姓名
	Email      string      `json:"email" structs:"email"`           // 邮箱
	Status     int         `json:"status" structs:"status"`         // 用户状态 1：正常 2：禁用
	LoginTime  int         `json:"loginTime" structs:"loginTime"`   // 登录时间
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
	Token      string      `json:"token" structs:"token"`           // 用户凭证
}
