package util

// Resource 溯源列实体
type Resource struct {
	Id         int    `gorm:"primarykey;comment:''"` //
	MemberId   int    `gorm:"comment:'员工编号'"`        // 员工编号
	DeviceCode string `gorm:"comment:'设备编号'"`        // 设备编号
	Date       string `gorm:"comment:'扫码日期'"`        // 扫码日期
	ImgTop     string `gorm:"comment:'图片_上'"`        // 图片_上
	ImgFront   string `gorm:"comment:'图片_前'"`        // 图片_前
	ImgBehind  string `gorm:"comment:'图片_后'"`        // 图片_后
	ImgLeft    string `gorm:"comment:'图片_左'"`        // 图片_左
	ImgRight   string `gorm:"comment:'图片_右'"`        // 图片_右
	//ImgS       string    `gorm:"comment:'图片集合'"`        // 图片集合
	VideoTop    string `gorm:"comment:'视频_上'"` // 视频
	VideoFront  string `gorm:"comment:'视频_前'"` // 视频
	VideoBehind string `gorm:"comment:'视频_后'"` // 视频
	VideoLeft   string `gorm:"comment:'视频_左'"` // 视频
	VideoRight  string `gorm:"comment:'视频_右'"` // 视频
	//Video      string `gorm:"comment:'视频'"`   // 视频
	CreateTime int64 `gorm:"comment:'创建时间'"` // 创建时间
}

type ImgResource struct {
	ImgTop    string `gorm:"comment:'图片_上'"` // 图片_上
	ImgFront  string `gorm:"comment:'图片_前'"` // 图片_前
	ImgBehind string `gorm:"comment:'图片_后'"` // 图片_后
	ImgLeft   string `gorm:"comment:'图片_左'"` // 图片_左
	ImgRight  string `gorm:"comment:'图片_右'"` // 图片_右
}
