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
