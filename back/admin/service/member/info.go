package member

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"gorm.io/gorm"
)

type IMemberInfoService interface {
	List(listReq req.MemberInfoListReq) (res req.MemberInfoResp, e error)
	Edit(editReq req.MemberInfoEditReq) (e error)
}

// NewResourceService 初始化
func NewInfoService(db *gorm.DB) IMemberInfoService {
	return &infoService{db: db}
}

// resourceService 溯源列服务实现类
type infoService struct {
	db *gorm.DB
}

func (srv infoService) List(listReq req.MemberInfoListReq) (res req.MemberInfoResp, e error) {
	model := srv.db.Model(&util.MemberInfo{})
	model = model.Where("create_date = ?", listReq.CreateDate)
	// 总数
	var count int64
	err := model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var obj util.MemberInfo
	err = model.Find(&obj).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resp := req.MemberInfoResp{}
	response.Copy(&resp, obj)
	return req.MemberInfoResp{
		ScanNumber:   resp.ScanNumber,
		CameraNumber: resp.CameraNumber,
		CreateDate:   resp.CreateDate,
	}, nil
}

func (hd infoService) Edit(editReq req.MemberInfoEditReq) (e error) {
	var obj util.MemberInfo
	err := hd.db.Where("create_date = ?", editReq.CreateDate).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		response.Copy(&obj, editReq)
		err := hd.db.Create(&obj).Error
		e = response.CheckErr(err, "Add Create err")
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	if editReq.CameraNumber == 0 {
		editReq.CameraNumber = obj.CameraNumber
	}
	if editReq.ScanNumber == 0 {
		editReq.ScanNumber = obj.ScanNumber
	}
	editReq.CameraNumber += 1
	editReq.ScanNumber += 1
	// 更新
	response.Copy(&obj, editReq)
	err = hd.db.Model(&obj).Where("create_date = ?", obj.CreateDate).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}
