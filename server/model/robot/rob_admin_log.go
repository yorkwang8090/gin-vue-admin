// 自动生成模板RobAdminLog
package robot

// robAdminLog表 结构体  RobAdminLog
type RobAdminLog struct {
	Id       *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                   //id字段
	UserName string `json:"userName" form:"userName" gorm:"column:user_name;comment:操作人;size:50;"`        //操作人
	OperPath string `json:"operPath" form:"operPath" gorm:"column:oper_path;comment:操作pathname;size:50;"` //操作pathname
	OperPost string `json:"operPost" form:"operPost" gorm:"column:oper_post;comment:提交数据;size:255;"`      //提交数据
	Content  string `json:"content" form:"content" gorm:"column:content;comment:操作说明;size:255;"`          //操作说明
	UserIp   string `json:"userIp" form:"userIp" gorm:"column:user_ip;comment:操作ip;size:60;"`             //操作ip
	OperTime *int   `json:"operTime" form:"operTime" gorm:"column:oper_time;comment:操作时间;size:10;"`       //操作时间
}

// TableName robAdminLog表 RobAdminLog自定义表名 rob_admin_log
func (RobAdminLog) TableName() string {
	return "rob_admin_log"
}
