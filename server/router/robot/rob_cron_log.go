package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RobCronLogRouter struct {
}

// InitRobCronLogRouter 初始化 robCronLog表 路由信息
func (s *RobCronLogRouter) InitRobCronLogRouter(Router *gin.RouterGroup) {
	robCronLogRouter := Router.Group("robCronLog").Use(middleware.OperationRecord())
	robCronLogRouterWithoutRecord := Router.Group("robCronLog")
	var robCronLogApi = v1.ApiGroupApp.RobotApiGroup.RobCronLogApi
	{
		robCronLogRouter.POST("createRobCronLog", robCronLogApi.CreateRobCronLog)   // 新建robCronLog表
		robCronLogRouter.DELETE("deleteRobCronLog", robCronLogApi.DeleteRobCronLog) // 删除robCronLog表
		robCronLogRouter.DELETE("deleteRobCronLogByIds", robCronLogApi.DeleteRobCronLogByIds) // 批量删除robCronLog表
		robCronLogRouter.PUT("updateRobCronLog", robCronLogApi.UpdateRobCronLog)    // 更新robCronLog表
	}
	{
		robCronLogRouterWithoutRecord.GET("findRobCronLog", robCronLogApi.FindRobCronLog)        // 根据ID获取robCronLog表
		robCronLogRouterWithoutRecord.GET("getRobCronLogList", robCronLogApi.GetRobCronLogList)  // 获取robCronLog表列表
	}
}
