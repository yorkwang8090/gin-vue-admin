package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
    robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
)

type RobExchangeService struct {
}

// CreateRobExchange 创建robExchange表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeService *RobExchangeService) CreateRobExchange(robExchange *robot.RobExchange) (err error) {
	err = global.GVA_DB.Create(robExchange).Error
	return err
}

// DeleteRobExchange 删除robExchange表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeService *RobExchangeService)DeleteRobExchange(id string) (err error) {
	err = global.GVA_DB.Delete(&robot.RobExchange{},"id = ?",id).Error
	return err
}

// DeleteRobExchangeByIds 批量删除robExchange表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeService *RobExchangeService)DeleteRobExchangeByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]robot.RobExchange{},"id in ?",ids).Error
	return err
}

// UpdateRobExchange 更新robExchange表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeService *RobExchangeService)UpdateRobExchange(robExchange robot.RobExchange) (err error) {
	err = global.GVA_DB.Save(&robExchange).Error
	return err
}

// GetRobExchange 根据id获取robExchange表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeService *RobExchangeService)GetRobExchange(id string) (robExchange robot.RobExchange, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&robExchange).Error
	return
}

// GetRobExchangeInfoList 分页获取robExchange表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeService *RobExchangeService)GetRobExchangeInfoList(info robotReq.RobExchangeSearch) (list []robot.RobExchange, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&robot.RobExchange{})
    var robExchanges []robot.RobExchange
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&robExchanges).Error
	return  robExchanges, total, err
}
