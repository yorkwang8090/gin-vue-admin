package robot

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
	robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type RobUserService struct {
}

// Login
func (robuserService *RobUserService) Login(robuser *robot.RobUser) (userInfor *robot.RobUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user robot.RobUser
	err = global.GVA_DB.Where("user_name = ?", robuser.UserName).First(&user).Error
	if err == nil {
		//判断密码是否正确
		if ok := utils.BcryptCheck(robuser.UserPwd, user.UserPwd); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

// CreateRobUser 创建user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robuserService *RobUserService) CreateRobUser(robuser *robot.RobUser) (err error) {
	err = global.GVA_DB.Create(robuser).Error
	return err
}

// DeleteRobUser 删除user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robuserService *RobUserService) DeleteRobUser(id string) (err error) {
	err = global.GVA_DB.Delete(&robot.RobUser{}, "id = ?", id).Error
	return err
}

// DeleteRobUserByIds 批量删除user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robuserService *RobUserService) DeleteRobUserByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]robot.RobUser{}, "id in ?", ids).Error
	return err
}

// UpdateRobUser 更新user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robuserService *RobUserService) UpdateRobUser(robuser robot.RobUser) (err error) {
	err = global.GVA_DB.Save(&robuser).Error
	return err
}

// GetRobUser 根据id获取user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robuserService *RobUserService) GetRobUser(id string) (robuser robot.RobUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&robuser).Error
	return
}

// GetRobUserInfoList 分页获取user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robuserService *RobUserService) GetRobUserInfoList(info robotReq.RobUserSearch) (list []robot.RobUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&robot.RobUser{})
	var robusers []robot.RobUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	//添加排序
	err = db.Order("id desc").Find(&robusers).Error
	return robusers, total, err
}
