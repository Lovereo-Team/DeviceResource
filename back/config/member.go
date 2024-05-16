package config

var MemberConfig = memberConfig{
	// 管理缓存键
	BackstageManageKey: "backstage:manage",
	// 角色缓存键
	BackstageRolesKey: "backstage:roles",
	// 令牌缓存键
	BackstageTokenKey: "backstage:token:",
	// 令牌的集合
	BackstageTokenSet: "backstage:token:set:",

	// 免登录验证
	NotLoginUri: []string{
		"member:login",        // 登录接口
		"common:index:config", // 配置接口
	},

	// 免权限验证
	NotAuthUri: []string{
		"member:logout",         // 退出登录
		"member:menu:menus",     // 系统菜单
		"member:menu:route",     // 菜单路由
		"member:admin:upInfo",   // 管理员更新
		"member:admin:self",     // 管理员信息
		"member:role:all",       // 所有角色
		"member:post:all",       // 所有岗位
		"member:dept:list",      // 所有部门
		"setting:dict:type:all", // 所有字典类型
		"setting:dict:data:all", // 所有字典数据
		"article:cate:all",      // 所有文章分类
	},

	// 演示模式白名单
	ShowWhitelistUri: []string{
		"member:login",  // 登录接口
		"member:logout", // 退出登录
	},

	// 请求临时数据
	SuperAdminId:   1,
	ReqAdminIdKey:  "member_id",
	ReqRoleIdKey:   "role",
	ReqUsernameKey: "username",
	ReqNicknameKey: "real_name",
}

type memberConfig struct {
	BackstageManageKey string
	BackstageRolesKey  string
	BackstageTokenKey  string
	BackstageTokenSet  string
	NotLoginUri        []string
	NotAuthUri         []string
	ShowWhitelistUri   []string
	SuperAdminId       uint
	ReqAdminIdKey      string
	ReqRoleIdKey       string
	ReqUsernameKey     string
	ReqNicknameKey     string
}
