package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
    robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
)

type RobCronLogService struct {
}

// CreateRobCronLog 创建robCronLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robCronLogService *RobCronLogService) CreateRobCronLog(robCronLog *robot.RobCronLog) (err error) {
	err = global.GVA_DB.Create(robCronLog).Error
	return err
}

// DeleteRobCronLog 删除robCronLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robCronLogService *RobCronLogService)DeleteRobCronLog(id string) (err error) {
	err = global.GVA_DB.Delete(&robot.RobCronLog{},"id = ?",id).Error
	return err
}

// DeleteRobCronLogByIds 批量删除robCronLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robCronLogService *RobCronLogService)DeleteRobCronLogByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]robot.RobCronLog{},"id in ?",ids).Error
	return err
}

// UpdateRobCronLog 更新robCronLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robCronLogService *RobCronLogService)UpdateRobCronLog(robCronLog robot.RobCronLog) (err error) {
	err = global.GVA_DB.Save(&robCronLog).Error
	return err
}

// GetRobCronLog 根据id获取robCronLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robCronLogService *RobCronLogService)GetRobCronLog(id string) (robCronLog robot.RobCronLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&robCronLog).Error
	return
}

// GetRobCronLogInfoList 分页获取robCronLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robCronLogService *RobCronLogService)GetRobCronLogInfoList(info robotReq.RobCronLogSearch) (list []robot.RobCronLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&robot.RobCronLog{})
    var robCronLogs []robot.RobCronLog
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&robCronLogs).Error
	return  robCronLogs, total, err
}
