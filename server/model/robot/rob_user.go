// 自动生成模板RobUser
package robot

// user表 结构体  RobUser
type RobUser struct {
	Id           *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;" binding:"required"`             //主键
	UserName     string `json:"userName" form:"userName" gorm:"column:user_name;comment:登陆账号;size:50;"`              //登陆账号
	UserNickName string `json:"userNickName" form:"userNickName" gorm:"column:user_nick_name;comment:英文昵称;size:50;"` //英文昵称
	UserPwd      string `json:"userPwd" form:"userPwd" gorm:"column:user_pwd;comment:登陆密码;size:100;"`                //登陆密码
	Mobile       string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:通知邮件;size:100;"`                    //通知邮件
	Addtime      *int   `json:"addtime" form:"addtime" gorm:"column:addtime;comment:添加时间;size:10;"`                  //添加时间
	Exptime      *int   `json:"exptime" form:"exptime" gorm:"column:exptime;comment:;size:10;"`                      //exptime字段
	Status       *bool  `json:"status" form:"status" gorm:"column:status;comment:;"`                                 //status字段
}

// TableName user表 RobUser自定义表名 rob_user
func (RobUser) TableName() string {
	return "rob_user"
}
