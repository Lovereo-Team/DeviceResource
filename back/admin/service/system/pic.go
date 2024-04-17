package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type IAlbumService interface {
	Add(addReq req.AlbumAddReq) (e error)
}

// NewAlbumService 初始化
func NewAlbumService(db *gorm.DB) IAlbumService {
	return &albumService{db: db}
}

// albumService 相册管理服务实现类
type albumService struct {
	db *gorm.DB
}

// Add 相册管理新增
func (srv albumService) Add(addReq req.AlbumAddReq) (e error) {
	var add req.AlbumAdd
	if len(addReq.ImgTop) > 1 {
		for _, img := range addReq.ImgTop {
			add.Cid = addReq.ImgTopCameraNumber
			add.Type = 10
			add.Aid = 1
			add.Uid = 1
			// 找到最后一个 "/" 的位置
			index := strings.LastIndex(img, "/")
			if index != -1 {
				// 获取 "/" 后面的子串
				add.Name = img[index+1:]
			}
			add.Uri = "image/camera_0/" + add.Name
			add.Ext = ".jpg"
			var obj util.Album
			response.Copy(&obj, add)
			err := srv.db.Create(&obj).Error
			e = response.CheckErr(err, "Add Create err")
		}
	}
	var number int
	if len(addReq.ImgFront) > 0 {
		for _, img := range addReq.ImgFront {
			number += 1
			add.Cid = addReq.ImgFrontCameraNumber
			add.Type = 10
			add.Aid = 1
			add.Uid = 1
			// 找到最后一个 "/" 的位置
			index := strings.LastIndex(img, "/")
			if index != -1 {
				// 获取 "/" 后面的子串
				add.Name = img[index+1:]
			}
			add.Uri = "image/camera_1/" + add.Name
			add.Ext = ".jpg"
			var obj util.Album
			response.Copy(&obj, add)
			err := srv.db.Create(&obj).Error
			e = response.CheckErr(err, "Add Create err")
		}
	}
	fmt.Println("a", number)
	if len(addReq.ImgBehind) > 1 {
		for _, img := range addReq.ImgBehind {
			add.Cid = addReq.ImgBehindCameraNumber
			add.Type = 10
			add.Aid = 1
			add.Uid = 1
			// 找到最后一个 "/" 的位置
			index := strings.LastIndex(img, "/")
			if index != -1 {
				// 获取 "/" 后面的子串
				add.Name = img[index+1:]
			}
			add.Uri = "image/camera_0/" + add.Name
			add.Ext = ".jpg"
			var obj util.Album
			response.Copy(&obj, add)
			err := srv.db.Create(&obj).Error
			e = response.CheckErr(err, "Add Create err")
		}
	}
	if len(addReq.ImgLeft) > 1 {
		for _, img := range addReq.ImgLeft {
			add.Cid = addReq.ImgLeftCameraNumber
			add.Type = 10
			add.Aid = 1
			add.Uid = 1
			// 找到最后一个 "/" 的位置
			index := strings.LastIndex(img, "/")
			if index != -1 {
				// 获取 "/" 后面的子串
				add.Name = img[index+1:]
			}
			add.Uri = "image/camera_0/" + add.Name
			add.Ext = ".jpg"
			var obj util.Album
			response.Copy(&obj, add)
			err := srv.db.Create(&obj).Error
			e = response.CheckErr(err, "Add Create err")
		}
	}
	if len(addReq.ImgRight) > 1 {
		for _, img := range addReq.ImgRight {
			add.Cid = addReq.ImgRightCameraNumber
			add.Type = 10
			add.Aid = 1
			add.Uid = 1
			// 找到最后一个 "/" 的位置
			index := strings.LastIndex(img, "/")
			if index != -1 {
				// 获取 "/" 后面的子串
				add.Name = img[index+1:]
			}
			add.Uri = "image/camera_0/" + add.Name
			add.Ext = ".jpg"
			var obj util.Album
			response.Copy(&obj, add)
			err := srv.db.Create(&obj).Error
			e = response.CheckErr(err, "Add Create err")
		}
	}
	return
}
