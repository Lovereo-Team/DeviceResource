package util

// Album 相册管理实体
type Album struct {
	Id         int    `gorm:"primarykey;comment:'主键ID'"`      // 主键ID
	Cid        int    `gorm:"comment:'类目ID'"`                 // 类目ID
	Aid        int    `gorm:"comment:'管理员ID'"`                // 管理员ID
	Uid        int    `gorm:"comment:'用户ID'"`                 // 用户ID
	Type       int    `gorm:"comment:'文件类型: [10=图片, 20=视频]'"` // 文件类型: [10=图片, 20=视频]
	Name       string `gorm:"comment:'文件名称'"`                 // 文件名称
	Uri        string `gorm:"comment:'文件路径'"`                 // 文件路径
	Ext        string `gorm:"comment:'文件扩展'"`                 // 文件扩展
	Size       int    `gorm:"comment:'文件大小'"`                 // 文件大小
	IsDelete   int    `gorm:"comment:'是否删除: 0=否, 1=是'"`       // 是否删除: 0=否, 1=是
	CreateTime int64  `gorm:"comment:'创建时间'"`                 // 创建时间
	UpdateTime int64  `gorm:"comment:'更新时间'"`                 // 更新时间
	DeleteTime int64  `gorm:"comment:'删除时间'"`                 // 删除时间
}
