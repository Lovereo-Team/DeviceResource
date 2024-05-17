package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/core/request"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type IResourceService interface {
	List(page request.PageReq, listReq req.ResourceListReq) (res response.PageResp, e error)
	Detail(id uint) (res req.ResourceResp, e error)
	Add(addReq req.ResourceAddReq) (e error)
	Edit(editReq req.ResourceEditReq) (e error)
	Del(id uint) (e error)
}

// NewResourceService 初始化
func NewResourceService(db *gorm.DB) IResourceService {
	return &resourceService{db: db}
}

// resourceService 溯源列服务实现类
type resourceService struct {
	db *gorm.DB
}

// List 溯源列列表
func (srv resourceService) List(page request.PageReq, listReq req.ResourceListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	model := srv.db.Model(&util.Resource{})
	if listReq.MemberId != 0 {
		model = model.Where("member_id = ?", listReq.MemberId)
	}
	if listReq.DeviceCode != "" {
		model = model.Where("device_code = ?", listReq.DeviceCode)
	}
	if !listReq.Date.Equal(time.Time{}) {
		model = model.Where("date = ?", listReq.Date)
	}
	// 总数
	var count int64
	err := model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var objs []util.Resource
	err = model.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := make([]req.ResourceResp, len(objs))
	response.Copy(&resps, objs)
	for i := range objs {
		resps[i].Id = objs[i].Id
		resps[i].MemberId = objs[i].MemberId
		resps[i].DeviceCode = objs[i].DeviceCode
		resps[i].Date = objs[i].Date[:10]
		ImgTop := strings.Split(objs[i].ImgTop, ",")
		for _, url := range ImgTop {
			resps[i].ImgTop = url
		}
		ImgFront := strings.Split(objs[i].ImgFront, ",")
		for _, url := range ImgFront {
			resps[i].ImgFront = url
		}
		ImgBehind := strings.Split(objs[i].ImgBehind, ",")
		for _, url := range ImgBehind {
			resps[i].ImgBehind = url
		}
		ImgLeft := strings.Split(objs[i].ImgLeft, ",")
		for _, url := range ImgLeft {
			resps[i].ImgLeft = url
		}
		ImgRight := strings.Split(objs[i].ImgRight, ",")
		for _, url := range ImgRight {
			resps[i].ImgLeft = url
		}
	}

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// Detail 溯源列详情
func (srv resourceService) Detail(id uint) (res req.ResourceResp, e error) {
	var obj util.Resource
	err := srv.db.Where("id = ?", id).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	ImgTop := strings.Split(obj.ImgTop, ",")
	for _, url := range ImgTop {
		obj.ImgTop = url
	}
	ImgFront := strings.Split(obj.ImgFront, ",")
	for _, url := range ImgFront {
		obj.ImgFront = url
	}
	ImgBehind := strings.Split(obj.ImgBehind, ",")
	for _, url := range ImgBehind {
		obj.ImgBehind = url
	}
	ImgLeft := strings.Split(obj.ImgLeft, ",")
	for _, url := range ImgLeft {
		obj.ImgLeft = url
	}
	ImgRight := strings.Split(obj.ImgRight, ",")
	for _, url := range ImgRight {
		obj.ImgLeft = url
	}
	response.Copy(&res, obj)
	return
}

// Add 溯源列新增
func (srv resourceService) Add(addReq req.ResourceAddReq) (e error) {
	var objs []util.Resource
	model := srv.db.Model(&util.Resource{})
	model = model.Where("device_code = ?", addReq.DeviceCode)
	model.Find(&objs)
	if len(objs) > 0 {
		srv.db.Where("device_code = ?", addReq.DeviceCode)
		addReq.CreateTime = int(time.Now().Unix())
		response.Copy(&objs, addReq)
		http.Get("http://127.0.0.1:9090/?code=" + objs[0].DeviceCode)
		srv.db.Model(&objs).Updates(objs)
		return
	}
	var obj_1 util.Resource
	addReq.VideoTop = "http://127.0.0.1:8000/api/uploads/image/camera_0/" + addReq.DeviceCode + ".mp4"
	addReq.VideoFront = "http://127.0.0.1:8000/api/uploads/image/camera_1/" + addReq.DeviceCode + ".mp4"
	addReq.VideoFront = "http://127.0.0.1:8000/api/uploads/image/camera_2/" + addReq.DeviceCode + ".mp4"
	addReq.CreateTime = int(time.Now().Unix())
	response.Copy(&obj_1, addReq)
	error := srv.db.Create(&obj_1).Error
	http.Get("http://127.0.0.1:9090/?code=" + obj_1.DeviceCode)
	e = response.CheckErr(error, "Add Create err")
	return
}

// Edit 溯源列编辑
func (srv resourceService) Edit(editReq req.ResourceEditReq) (e error) {
	var obj util.Resource
	err := srv.db.Where("device_code = ?", editReq.Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = srv.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 溯源列删除
func (srv resourceService) Del(id uint) (e error) {
	var obj util.Resource
	err := srv.db.Where("id = ?", id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 删除
	err = srv.db.Delete(&obj).Error
	e = response.CheckErr(err, "Del Delete err")
	return
}
