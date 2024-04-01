package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
    robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
)

type RobPairService struct {
}

// CreateRobPair 创建robPair表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robPairService *RobPairService) CreateRobPair(robPair *robot.RobPair) (err error) {
	err = global.GVA_DB.Create(robPair).Error
	return err
}

// DeleteRobPair 删除robPair表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robPairService *RobPairService)DeleteRobPair(id string) (err error) {
	err = global.GVA_DB.Delete(&robot.RobPair{},"id = ?",id).Error
	return err
}

// DeleteRobPairByIds 批量删除robPair表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robPairService *RobPairService)DeleteRobPairByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]robot.RobPair{},"id in ?",ids).Error
	return err
}

// UpdateRobPair 更新robPair表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robPairService *RobPairService)UpdateRobPair(robPair robot.RobPair) (err error) {
	err = global.GVA_DB.Save(&robPair).Error
	return err
}

// GetRobPair 根据id获取robPair表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robPairService *RobPairService)GetRobPair(id string) (robPair robot.RobPair, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&robPair).Error
	return
}

// GetRobPairInfoList 分页获取robPair表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robPairService *RobPairService)GetRobPairInfoList(info robotReq.RobPairSearch) (list []robot.RobPair, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&robot.RobPair{})
    var robPairs []robot.RobPair
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&robPairs).Error
	return  robPairs, total, err
}
