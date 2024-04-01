package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RobAdminLogRouter struct {
}

// InitRobAdminLogRouter 初始化 robAdminLog表 路由信息
func (s *RobAdminLogRouter) InitRobAdminLogRouter(Router *gin.RouterGroup) {
	robAdminLogRouter := Router.Group("robAdminLog").Use(middleware.OperationRecord())
	robAdminLogRouterWithoutRecord := Router.Group("robAdminLog")
	var robAdminLogApi = v1.ApiGroupApp.RobotApiGroup.RobAdminLogApi
	{
		robAdminLogRouter.POST("createRobAdminLog", robAdminLogApi.CreateRobAdminLog)   // 新建robAdminLog表
		robAdminLogRouter.DELETE("deleteRobAdminLog", robAdminLogApi.DeleteRobAdminLog) // 删除robAdminLog表
		robAdminLogRouter.DELETE("deleteRobAdminLogByIds", robAdminLogApi.DeleteRobAdminLogByIds) // 批量删除robAdminLog表
		robAdminLogRouter.PUT("updateRobAdminLog", robAdminLogApi.UpdateRobAdminLog)    // 更新robAdminLog表
	}
	{
		robAdminLogRouterWithoutRecord.GET("findRobAdminLog", robAdminLogApi.FindRobAdminLog)        // 根据ID获取robAdminLog表
		robAdminLogRouterWithoutRecord.GET("getRobAdminLogList", robAdminLogApi.GetRobAdminLogList)  // 获取robAdminLog表列表
	}
}
