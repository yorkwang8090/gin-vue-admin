// 自动生成模板RobExchange
package robot

// robExchange表 结构体  RobExchange
type RobExchange struct {
	Id       *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                        //id字段
	Title    string `json:"title" form:"title" gorm:"column:title;comment:交易所名称;size:50;"`                     //交易所名称
	ApiUrl   string `json:"apiUrl" form:"apiUrl" gorm:"column:api_url;comment:api接口地址;size:100;"`              //api接口地址
	Extend   string `json:"extend" form:"extend" gorm:"column:extend;comment:扩展参数，如comexvip的uid：156;size:50;"` //扩展参数，如comexvip的uid：156
	Isfollow *bool  `json:"isfollow" form:"isfollow" gorm:"column:isfollow;comment:是否为跟随交易所;"`                 //是否为跟随交易所
	Addtime  *int   `json:"addtime" form:"addtime" gorm:"column:addtime;comment:添加时间;size:10;"`                //添加时间
	Status   *bool  `json:"status" form:"status" gorm:"column:status;comment:状态;"`                             //状态
}

// TableName robExchange表 RobExchange自定义表名 rob_exchange
func (RobExchange) TableName() string {
	return "rob_exchange"
}
