package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
    robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
)

type RobAdminLogService struct {
}

// CreateRobAdminLog 创建robAdminLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robAdminLogService *RobAdminLogService) CreateRobAdminLog(robAdminLog *robot.RobAdminLog) (err error) {
	err = global.GVA_DB.Create(robAdminLog).Error
	return err
}

// DeleteRobAdminLog 删除robAdminLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robAdminLogService *RobAdminLogService)DeleteRobAdminLog(id string) (err error) {
	err = global.GVA_DB.Delete(&robot.RobAdminLog{},"id = ?",id).Error
	return err
}

// DeleteRobAdminLogByIds 批量删除robAdminLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robAdminLogService *RobAdminLogService)DeleteRobAdminLogByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]robot.RobAdminLog{},"id in ?",ids).Error
	return err
}

// UpdateRobAdminLog 更新robAdminLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robAdminLogService *RobAdminLogService)UpdateRobAdminLog(robAdminLog robot.RobAdminLog) (err error) {
	err = global.GVA_DB.Save(&robAdminLog).Error
	return err
}

// GetRobAdminLog 根据id获取robAdminLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robAdminLogService *RobAdminLogService)GetRobAdminLog(id string) (robAdminLog robot.RobAdminLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&robAdminLog).Error
	return
}

// GetRobAdminLogInfoList 分页获取robAdminLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robAdminLogService *RobAdminLogService)GetRobAdminLogInfoList(info robotReq.RobAdminLogSearch) (list []robot.RobAdminLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&robot.RobAdminLog{})
    var robAdminLogs []robot.RobAdminLog
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&robAdminLogs).Error
	return  robAdminLogs, total, err
}
