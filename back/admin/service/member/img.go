package member

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"gorm.io/gorm"
	"strings"
)

type IImgService interface {
	List(listReq req.ResourceListReq) (res req.ResourceRespRes, e error)
}

// NewImgService 初始化
func NewImgService(db *gorm.DB) IImgService {
	return &imgService{db: db}
}

// imgService 溯源列服务实现类
type imgService struct {
	db *gorm.DB
}

func (srv imgService) List(listReq req.ResourceListReq) (res req.ResourceRespRes, e error) {
	// 查询
	model := srv.db.Model(&util.Resource{})
	model = model.Where("device_code = ?", listReq.DeviceCode)
	// 总数
	var count int64
	err := model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var objs []util.Resource
	err = model.Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}

	var resps []req.Img
	for _, obj := range objs {
		img := req.Img{
			ImgTop:    strings.Split(obj.ImgTop, ","),
			ImgFront:  strings.Split(obj.ImgFront, ","),
			ImgBehind: strings.Split(obj.ImgBehind, ","),
			ImgLeft:   strings.Split(obj.ImgLeft, ","),
			ImgRight:  strings.Split(obj.ImgRight, ","),
		}
		resps = append(resps, img)
	}

	return req.ResourceRespRes{
		Lists: resps,
	}, nil
}
