package req

import (
	"DeviceResource/core"
	"time"
)

// ResourceListReq 溯源列列表参数
type ResourceListReq struct {
	MemberId    int       `form:"memberId"`    // 员工编号
	DeviceCode  string    `form:"deviceCode"`  // 设备编号
	Date        time.Time `form:"date"`        // 扫码日期
	ImgTop      string    `form:"imgTop"`      // 图片_上
	ImgFront    string    `form:"imgFront"`    // 图片_前
	ImgBehind   string    `form:"imgBehind"`   // 图片_后
	ImgLeft     string    `form:"imgLeft"`     // 图片_左
	ImgRight    string    `form:"imgRight"`    // 图片_右
	ImgS        string    `form:"imgS"`        // 图片集合
	VideoTop    string    `form:"videoTop"`    // 视频
	VideoFront  string    `form:"videoFront"`  // 视频
	VideoBehind string    `form:"videoBehind"` // 视频
	VideoLeft   string    `form:"videoLeft"`   // 视频
	VideoRight  string    `form:"videoRight"`  // 视频
}

// ResourceDetailReq 溯源列详情参数
type ResourceDetailReq struct {
	Id uint `form:"id"` //
}

// ResourceAddReq 溯源列新增参数
type ResourceAddReq struct {
	MemberId    int    `form:"memberId"`    // 员工编号
	DeviceCode  string `form:"deviceCode"`  // 设备编号
	Date        string `form:"date"`        // 扫码日期
	ImgTop      string `form:"imgTop"`      // 图片_上
	ImgFront    string `form:"imgFront"`    // 图片_前
	ImgBehind   string `form:"imgBehind"`   // 图片_后
	ImgLeft     string `form:"imgLeft"`     // 图片_左
	ImgRight    string `form:"imgRight"`    // 图片_右
	ImgS        string `form:"imgS"`        // 图片集合
	VideoTop    string `form:"videoTop"`    // 视频
	VideoFront  string `form:"videoFront"`  // 视频
	VideoBehind string `form:"videoBehind"` // 视频
	VideoLeft   string `form:"videoLeft"`   // 视频
	VideoRight  string `form:"videoRight"`  // 视频
	CreateTime  int    `form:"createTime"`  // 创建时间
}

// ResourceEditReq 溯源列新增参数
type ResourceEditReq struct {
	Id          int         `form:"id"`          //
	MemberId    int         `form:"memberId"`    // 员工编号
	DeviceCode  string      `form:"deviceCode"`  // 设备编号
	Date        core.TsTime `form:"date"`        // 扫码日期
	ImgTop      string      `form:"imgTop"`      // 图片_上
	ImgFront    string      `form:"imgFront"`    // 图片_前
	ImgBehind   string      `form:"imgBehind"`   // 图片_后
	ImgLeft     string      `form:"imgLeft"`     // 图片_左
	ImgRight    string      `form:"imgRight"`    // 图片_右
	ImgS        string      `form:"imgS"`        // 图片集合
	VideoTop    string      `form:"videoTop"`    // 视频
	VideoFront  string      `form:"videoFront"`  // 视频
	VideoBehind string      `form:"videoBehind"` // 视频
	VideoLeft   string      `form:"videoLeft"`   // 视频
	VideoRight  string      `form:"videoRight"`  // 视频
}

// ResourceDelReq 溯源列新增参数
type ResourceDelReq struct {
	Id uint `form:"id"` //
}

// ResourceResp 溯源列返回信息
type ResourceResp struct {
	Id          int         `json:"id" structs:"id"`                 //
	MemberId    int         `json:"memberId" structs:"memberId"`     // 员工编号
	DeviceCode  string      `json:"deviceCode" structs:"deviceCode"` // 设备编号
	Date        string      `json:"date" structs:"date"`             // 扫码日期
	ImgTop      string      `json:"imgTop" structs:"imgTop"`         // 图片_上
	ImgFront    string      `json:"imgFront" structs:"imgFront"`     // 图片_前
	ImgBehind   string      `json:"imgBehind" structs:"imgBehind"`   // 图片_后
	ImgLeft     string      `json:"imgLeft" structs:"imgLeft"`       // 图片_左
	ImgRight    string      `json:"imgRight" structs:"imgRight"`     // 图片_右
	ImgS        string      `json:"imgS" structs:"imgS"`             // 图片集合
	VideoTop    string      `form:"videoTop"`                        // 视频
	VideoFront  string      `form:"videoFront"`                      // 视频
	VideoBehind string      `form:"videoBehind"`                     // 视频
	VideoLeft   string      `form:"videoLeft"`                       // 视频
	VideoRight  string      `form:"videoRight"`                      // 视频
	CreateTime  core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
}

type Img struct {
	ImgTop    []string `json:"imgTop" structs:"imgTop"`       // 图片_上
	ImgFront  []string `json:"imgFront" structs:"imgFront"`   // 图片_前
	ImgBehind []string `json:"imgBehind" structs:"imgBehind"` // 图片_后
	ImgLeft   []string `json:"imgLeft" structs:"imgLeft"`     // 图片_左
	ImgRight  []string `json:"imgRight" structs:"imgRight"`   // 图片_右
}

type ResourceRespRes struct {
	Lists interface{} `json:"lists"` // 数据
}

type VideoRespRes struct {
	VideoTop    string `form:"videoTop"`    // 视频
	VideoFront  string `form:"videoFront"`  // 视频
	VideoBehind string `form:"videoBehind"` // 视频
	VideoLeft   string `form:"videoLeft"`   // 视频
	VideoRight  string `form:"videoRight"`  // 视频
}
