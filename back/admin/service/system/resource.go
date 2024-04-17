package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/core/request"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"fmt"
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
	//fmt.Println("a", len(objs))
	fmt.Println("a", objs)
	//resps := []req.ResourceResp{}

	resps := make([]req.ResourceResp, len(objs))
	response.Copy(&resps, objs)
	for i := range objs {
		//truncatedDate := objs[i].Date.Truncate(24 * time.Hour) // 一天的时间戳
		// 将时间戳格式化为字符串，只保留年月日
		//date := objs[i].Date.Format("2006-01-02") // 格式为 "年-月-日"
		//fmt.Println("a", date)
		resps[i].Id = objs[i].Id
		resps[i].MemberId = objs[i].MemberId
		resps[i].DeviceCode = objs[i].DeviceCode
		resps[i].Date = objs[i].Date[:10]
		ImgTop := strings.Split(objs[i].ImgTop, ",")
		//fmt.Println("b", objs[i].ImgTop)
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
		//resps[i].ImgS = objs[i].ImgS
		//resps[i].Video = objs[i].Video
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
	response.Copy(&res, obj)
	return
}

// Add 溯源列新增
func (srv resourceService) Add(addReq req.ResourceAddReq) (e error) {
	var obj util.Resource
	model := srv.db.Model(&util.Resource{})
	model.Find(&obj)
	if obj.DeviceCode == addReq.DeviceCode {
		return
	}
	addReq.VideoTop = "http://127.0.0.1:8000/api/uploads/image/camera_0/" + addReq.DeviceCode + ".mp4"
	addReq.VideoFront = "http://127.0.0.1:8000/api/uploads/image/camera_1/" + addReq.DeviceCode + ".mp4"
	addReq.CreateTime = int(time.Now().Unix())
	response.Copy(&obj, addReq)
	fmt.Println("a", obj.Date)
	error := srv.db.Create(&obj).Error
	http.Get("http://127.0.0.1:9090/?code=" + obj.DeviceCode)
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
	fmt.Println("a", editReq)
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
