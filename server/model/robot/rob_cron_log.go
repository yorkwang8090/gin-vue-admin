// 自动生成模板RobCronLog
package robot

// robCronLog表 结构体  RobCronLog
type RobCronLog struct {
	Id      *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`         //id字段
	Uid     *int   `json:"uid" form:"uid" gorm:"column:uid;comment:所属用户;size:10;"`             //所属用户
	Exid    *int   `json:"exid" form:"exid" gorm:"column:exid;comment:交易所;size:10;"`           //交易所
	Pid     *int   `json:"pid" form:"pid" gorm:"column:pid;comment:交易对;size:10;"`              //交易对
	Symbol  string `json:"symbol" form:"symbol" gorm:"column:symbol;comment:交易对;size:50;"`     //交易对
	Type    string `json:"type" form:"type" gorm:"column:type;comment:类型 对敲  趋势  跟随;size:20;"` //类型 对敲  趋势  跟随
	Content string `json:"content" form:"content" gorm:"column:content;comment:内容;size:1000;"` //内容
	Addtime *int   `json:"addtime" form:"addtime" gorm:"column:addtime;comment:;size:10;"`     //addtime字段
}

// TableName robCronLog表 RobCronLog自定义表名 rob_cron_log
func (RobCronLog) TableName() string {
	return "rob_cron_log"
}
