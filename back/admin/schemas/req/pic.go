package req

import "DeviceResource/core"

// AlbumListReq 相册管理列表参数
type AlbumListReq struct {
	Cid  int    `form:"cid"`  // 类目ID
	Aid  int    `form:"aid"`  // 管理员ID
	Uid  int    `form:"uid"`  // 用户ID
	Type int    `form:"type"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"name"` // 文件名称
	Uri  string `form:"uri"`  // 文件路径
	Ext  string `form:"ext"`  // 文件扩展
	Size int    `form:"size"` // 文件大小
}

// AlbumDetailReq 相册管理详情参数
type AlbumDetailReq struct {
	Id int `form:"id"` // 主键ID
}

// AlbumAdd 相册管理新增参数
type AlbumAdd struct {
	Cid  int    `form:"cid"`  // 类目ID
	Aid  int    `form:"aid"`  // 管理员ID
	Uid  int    `form:"uid"`  // 用户ID
	Type int    `form:"type"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"name"` // 文件名称
	Uri  string `form:"uri"`  // 文件路径
	Ext  string `form:"ext"`  // 文件扩展
	Size int    `form:"size"` // 文件大小
}

// AlbumEditReq 相册管理新增参数
type AlbumEditReq struct {
	Id   int    `form:"id"`   // 主键ID
	Cid  int    `form:"cid"`  // 类目ID
	Aid  int    `form:"aid"`  // 管理员ID
	Uid  int    `form:"uid"`  // 用户ID
	Type int    `form:"type"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"name"` // 文件名称
	Uri  string `form:"uri"`  // 文件路径
	Ext  string `form:"ext"`  // 文件扩展
	Size int    `form:"size"` // 文件大小
}

// AlbumDelReq 相册管理新增参数
type AlbumDelReq struct {
	Id int `form:"id"` // 主键ID
}

// AlbumResp 相册管理返回信息
type AlbumResp struct {
	Id         int         `json:"id" structs:"id"`                 // 主键ID
	Cid        int         `json:"cid" structs:"cid"`               // 类目ID
	Aid        int         `json:"aid" structs:"aid"`               // 管理员ID
	Uid        int         `json:"uid" structs:"uid"`               // 用户ID
	Type       int         `json:"type" structs:"type"`             // 文件类型: [10=图片, 20=视频]
	Name       string      `json:"name" structs:"name"`             // 文件名称
	Uri        string      `json:"uri" structs:"uri"`               // 文件路径
	Ext        string      `json:"ext" structs:"ext"`               // 文件扩展
	Size       int         `json:"size" structs:"size"`             // 文件大小
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

type AlbumAddReq struct {
	ImgBehindCameraNumber int      `json:"img_behind_camera_number"`
	ImgFrontCameraNumber  int      `json:"img_front_camera_number"`
	ImgLeftCameraNumber   int      `json:"img_left_camera_number"`
	ImgRightCameraNumber  int      `json:"img_right_camera_number"`
	ImgTopCameraNumber    int      `json:"img_top_camera_number"`
	ImgBehind             []string `json:"imgBehind"`
	ImgFront              []string `json:"imgFront"`
	ImgLeft               []string `json:"imgLeft"`
	ImgRight              []string `json:"imgRight"`
	ImgTop                []string `json:"imgTop"`
}
