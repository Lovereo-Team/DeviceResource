package member

import (
	"DeviceResource/config"
	"DeviceResource/util"
	"gorm.io/gorm"
	"strconv"
)

type IMemberAuthAdminService interface {
	FindByUsername(username string) (member util.Member, err error)
	CacheAdminUserByUid(id int) (err error)
}

func NewMemberAuthAdminService(db *gorm.DB) IMemberAuthAdminService {
	return &memberAuthAdminService{db: db}
}

type memberAuthAdminService struct {
	db *gorm.DB
}

func (memberSrv memberAuthAdminService) FindByUsername(username string) (member util.Member, err error) {
	err = memberSrv.db.Where("username = ?", username).Limit(1).First(&member).Error
	return
}

func (memberSrv memberAuthAdminService) CacheAdminUserByUid(id int) (err error) {
	var member util.Member
	err = memberSrv.db.Where("id = ?", id).Limit(1).First(&member).Error
	if err != nil {
		return
	}
	str, err := util.ToolsUtil.ObjToJson(&member)
	if err != nil {
		return
	}
	util.RedisUtil.HSet(config.AdminConfig.BackstageManageKey, strconv.FormatUint(uint64(member.Id), 10), str, 0)
	return nil
}
