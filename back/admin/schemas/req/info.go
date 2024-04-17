package req

// MemberInfoListReq 【请填写功能名称】列表参数
type MemberInfoListReq struct {
	ScanNumber   int    `form:"scanNumber"`   // 今日扫描次数
	CameraNumber int    `form:"cameraNumber"` // 今日拍摄台数
	CreateDate   string `form:"createDate"`   // 日期
}

// MemberInfoDetailReq 【请填写功能名称】详情参数
type MemberInfoDetailReq struct {
	Id int `form:"id"` //
}

// MemberInfoAddReq 【请填写功能名称】新增参数
type MemberInfoAddReq struct {
	ScanNumber   int    `form:"scanNumber"`   // 今日扫描次数
	CameraNumber int    `form:"cameraNumber"` // 今日拍摄台数
	CreateDate   string `form:"createDate"`   // 日期
}

// MemberInfoEditReq 【请填写功能名称】新增参数
type MemberInfoEditReq struct {
	Id           int    `form:"id"`           //
	ScanNumber   int    `form:"scanNumber"`   // 今日扫描次数
	CameraNumber int    `form:"cameraNumber"` // 今日拍摄台数
	CreateDate   string `form:"createDate"`   // 日期
}

// MemberInfoDelReq 【请填写功能名称】新增参数
type MemberInfoDelReq struct {
	Id int `form:"id"` //
}

// MemberInfoResp 【请填写功能名称】返回信息
type MemberInfoResp struct {
	ScanNumber   int    `json:"scanNumber" structs:"scanNumber"`     // 今日扫描次数
	CameraNumber int    `json:"cameraNumber" structs:"cameraNumber"` // 今日拍摄台数
	CreateDate   string `json:"createDate" structs:"createDate"`     // 日期
}
