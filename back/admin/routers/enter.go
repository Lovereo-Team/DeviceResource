package routers

import (
	"DeviceResource/admin/routers/common"
	"DeviceResource/admin/routers/monitor"
	"DeviceResource/admin/routers/setting"
	"DeviceResource/admin/routers/system"
	"DeviceResource/core"
)

var InitRouters = []*core.GroupBase{
	// common
	common.AlbumGroup,
	common.IndexGroup,
	common.UploadGroup,
	// monitor
	monitor.MonitorGroup,
	// setting
	setting.CopyrightGroup,
	setting.DictDataGroup,
	setting.DictTypeGroup,
	setting.ProtocolGroup,
	setting.StorageGroup,
	setting.WebsiteGroup,
	// system
	system.AdminGroup,
	system.DeptGroup,
	system.LogGroup,
	system.LoginGroup,
	system.MenuGroup,
	system.PostGroup,
	system.RoleGroup,
	system.MemberGroup,
	system.ResourceGroup,
}
