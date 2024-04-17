package util

// Member 员工列实体
type Member struct {
	Id         int    `gorm:"primarykey;comment:''"`    //
	Username   string `gorm:"comment:'用户名'"`            // 用户名
	Password   string `gorm:"comment:'密码'"`             // 密码
	RealName   string `gorm:"comment:'真实姓名'"`           // 真实姓名
	Email      string `gorm:"comment:'邮箱'"`             // 邮箱
	Status     int    `gorm:"comment:'用户状态 1：正常 2：禁用'"` // 用户状态 1：正常 2：禁用
	LoginTime  int    `gorm:"comment:'登录时间'"`           // 登录时间
	CreateTime int64  `gorm:"comment:'创建时间'"`           // 创建时间
	UpdateTime int64  `gorm:"comment:'更新时间'"`           // 更新时间
	Token      string `gorm:"comment:'用户凭证'"`           // 用户凭证
}

type MemberLogLogin struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	MemberId   int    `gorm:"not null;default:0;comment:'管理员ID'"`
	Username   string `gorm:"not null;default:'';comment:'登录账号'"`
	Ip         string `gorm:"not null;default:'';comment:'登录地址'"`
	Os         string `gorm:"not null;default:'';comment:'操作系统'"`
	Browser    string `gorm:"not null;default:'';comment:'浏览器'"`
	Status     uint8  `gorm:"not null;default:0;comment:'操作状态: 1=成功, 0=失败'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
}

// MemberInfo 【请填写功能名称】实体
type MemberInfo struct {
	Id           int    `gorm:"primarykey;comment:''"` //
	ScanNumber   int    `gorm:"comment:'今日扫描次数'"`      // 今日扫描次数
	CameraNumber int    `gorm:"comment:'今日拍摄台数'"`      // 今日拍摄台数
	CreateDate   string `gorm:"comment:'日期'"`          // 日期
}
