package system

import (
	"DeviceResource/admin/schemas/req"
	"DeviceResource/core/request"
	"DeviceResource/core/response"
	"DeviceResource/util"
	"gorm.io/gorm"
	"time"
)

type IMemberService interface {
	List(page request.PageReq, listReq req.MemberListReq) (res response.PageResp, e error)
	Detail(id uint) (res req.MemberResp, e error)
	Add(addReq req.MemberAddReq) (e error)
	Edit(editReq req.MemberEditReq) (e error)
	Del(id uint) (e error)
}

// NewMemberService 初始化
func NewMemberService(db *gorm.DB) IMemberService {
	return &memberService{db: db}
}

// memberService 员工列服务实现类
type memberService struct {
	db *gorm.DB
}

// List 员工列列表
func (srv memberService) List(page request.PageReq, listReq req.MemberListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	model := srv.db.Model(&util.Member{})
	if listReq.Username != "" {
		model = model.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Password != "" {
		model = model.Where("password = ?", listReq.Password)
	}
	if listReq.RealName != "" {
		model = model.Where("real_name like ?", "%"+listReq.RealName+"%")
	}
	if listReq.Email != "" {
		model = model.Where("email = ?", listReq.Email)
	}
	if listReq.Status != 0 {
		model = model.Where("status = ?", listReq.Status)
	}
	//if listReq.LoginTime >= 0 {
	//	model = model.Where("login_time = ?", listReq.LoginTime)
	//}
	if listReq.Token != "" {
		model = model.Where("token = ?", listReq.Token)
	}
	// 总数
	var count int64
	err := model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []util.Member
	err = model.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []req.MemberResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// Detail 员工列详情
func (srv memberService) Detail(id uint) (res req.MemberResp, e error) {
	var obj util.Member
	err := srv.db.Where("id = ?", id).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	obj.Password = ""
	response.Copy(&res, obj)
	return
}

// Add 员工列新增
func (srv memberService) Add(addReq req.MemberAddReq) (e error) {
	var obj util.Member
	addReq.Password = util.ToolsUtil.MakeMd5(addReq.Password)
	addReq.Token = util.ToolsUtil.MakeToken()
	addReq.CreateTime = int(time.Now().Unix())
	response.Copy(&obj, addReq)
	err := srv.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 员工列编辑
func (srv memberService) Edit(editReq req.MemberEditReq) (e error) {
	var obj util.Member
	err := srv.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	editReq.Password = util.ToolsUtil.MakeMd5(editReq.Password)
	editReq.UpdateTime = int(time.Now().Unix())
	response.Copy(&obj, editReq)
	err = srv.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 员工列删除
func (srv memberService) Del(id uint) (e error) {
	var obj util.Member
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
