// 自动生成模板RobPair
package robot

import (
	
	
	
)

// robPair表 结构体  RobPair
type RobPair struct {

      Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id字段 
      Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:所属用户;size:10;"`  //所属用户 
      Exid  *int `json:"exid" form:"exid" gorm:"column:exid;comment:所属交易所;size:10;"`  //所属交易所 
      Symbol  string `json:"symbol" form:"symbol" gorm:"column:symbol;comment:交易对   BTC_USDT;size:50;"`  //交易对   BTC_USDT 
      PricePrecision  *int `json:"pricePrecision" form:"pricePrecision" gorm:"column:price_precision;comment:价格小数位数;size:10;"`  //价格小数位数 
      VolumePrecision  *int `json:"volumePrecision" form:"volumePrecision" gorm:"column:volume_precision;comment:数量小数位数;size:10;"`  //数量小数位数 
      Ak  string `json:"ak" form:"ak" gorm:"column:ak;comment:;size:100;"`  //ak字段 
      Sk  string `json:"sk" form:"sk" gorm:"column:sk;comment:;size:100;"`  //sk字段 
      Beat  string `json:"beat" form:"beat" gorm:"column:beat;comment:对敲;"`  //对敲 
      Trend  string `json:"trend" form:"trend" gorm:"column:trend;comment:趋势;"`  //趋势 
      Follow  string `json:"follow" form:"follow" gorm:"column:follow;comment:跟随;"`  //跟随 
      Pankou  string `json:"pankou" form:"pankou" gorm:"column:pankou;comment:盘口自动挂单撤单设置;"`  //盘口自动挂单撤单设置 
      Addtime  *int `json:"addtime" form:"addtime" gorm:"column:addtime;comment:添加时间;size:10;"`  //添加时间 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:;"`  //status字段 
}


// TableName robPair表 RobPair自定义表名 rob_pair
func (RobPair) TableName() string {
  return "rob_pair"
}

