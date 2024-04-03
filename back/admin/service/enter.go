package service

import (
	"DeviceResource/admin/service/common"
	"DeviceResource/admin/service/setting"
	"DeviceResource/admin/service/system"
)

var InitFunctions = []interface{}{
	// common
	common.NewAlbumService,
	common.NewIndexService,
	common.NewUploadService,
	// setting
	setting.NewSettingCopyrightService,
	setting.NewSettingDictDataService,
	setting.NewSettingDictTypeService,
	setting.NewSettingProtocolService,
	setting.NewSettingStorageService,
	setting.NewSettingWebsiteService,
	// system
	system.NewSystemAuthAdminService,
	system.NewSystemAuthDeptService,
	system.NewSystemAuthMenuService,
	system.NewSystemAuthPermService,
	system.NewSystemAuthPostService,
	system.NewSystemAuthRoleService,
	system.NewSystemLoginService,
	system.NewSystemLogsServer,
	system.NewMemberService,
	system.NewResourceService,
}
