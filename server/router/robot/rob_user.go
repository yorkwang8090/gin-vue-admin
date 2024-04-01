package robot

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RobUserRouter struct {
}

// InitRobUserRouter 初始化 user表 路由信息
func (s *RobUserRouter) InitRobUserRouter(Router *gin.RouterGroup) {
	robuserRouter := Router.Group("robuser").Use(middleware.OperationRecord())
	robuserRouterWithoutRecord := Router.Group("robuser")
	var robuserApi = v1.ApiGroupApp.RobotApiGroup.RobUserApi
	{
		robuserRouter.POST("createRobUser", robuserApi.CreateRobUser)             // 新建user表
		robuserRouter.DELETE("deleteRobUser", robuserApi.DeleteRobUser)           // 删除user表
		robuserRouter.DELETE("deleteRobUserByIds", robuserApi.DeleteRobUserByIds) // 批量删除user表
		robuserRouter.PUT("updateRobUser", robuserApi.UpdateRobUser)              // 更新user表
	}
	{
		robuserRouterWithoutRecord.GET("findRobUser", robuserApi.FindRobUser)       // 根据ID获取user表
		robuserRouterWithoutRecord.GET("getRobUserList", robuserApi.GetRobUserList) // 获取user表列表
	}
}
