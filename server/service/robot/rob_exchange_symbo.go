package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
    robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
)

type RobExchangeSymboService struct {
}

// CreateRobExchangeSymbo 创建robExchangeSymbo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeSymboService *RobExchangeSymboService) CreateRobExchangeSymbo(robExchangeSymbo *robot.RobExchangeSymbo) (err error) {
	err = global.GVA_DB.Create(robExchangeSymbo).Error
	return err
}

// DeleteRobExchangeSymbo 删除robExchangeSymbo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeSymboService *RobExchangeSymboService)DeleteRobExchangeSymbo(id string) (err error) {
	err = global.GVA_DB.Delete(&robot.RobExchangeSymbo{},"id = ?",id).Error
	return err
}

// DeleteRobExchangeSymboByIds 批量删除robExchangeSymbo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeSymboService *RobExchangeSymboService)DeleteRobExchangeSymboByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]robot.RobExchangeSymbo{},"id in ?",ids).Error
	return err
}

// UpdateRobExchangeSymbo 更新robExchangeSymbo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeSymboService *RobExchangeSymboService)UpdateRobExchangeSymbo(robExchangeSymbo robot.RobExchangeSymbo) (err error) {
	err = global.GVA_DB.Save(&robExchangeSymbo).Error
	return err
}

// GetRobExchangeSymbo 根据id获取robExchangeSymbo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeSymboService *RobExchangeSymboService)GetRobExchangeSymbo(id string) (robExchangeSymbo robot.RobExchangeSymbo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&robExchangeSymbo).Error
	return
}

// GetRobExchangeSymboInfoList 分页获取robExchangeSymbo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (robExchangeSymboService *RobExchangeSymboService)GetRobExchangeSymboInfoList(info robotReq.RobExchangeSymboSearch) (list []robot.RobExchangeSymbo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&robot.RobExchangeSymbo{})
    var robExchangeSymbos []robot.RobExchangeSymbo
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&robExchangeSymbos).Error
	return  robExchangeSymbos, total, err
}
