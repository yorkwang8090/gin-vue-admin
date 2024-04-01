package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RobPairRouter struct {
}

// InitRobPairRouter 初始化 robPair表 路由信息
func (s *RobPairRouter) InitRobPairRouter(Router *gin.RouterGroup) {
	robPairRouter := Router.Group("robPair").Use(middleware.OperationRecord())
	robPairRouterWithoutRecord := Router.Group("robPair")
	var robPairApi = v1.ApiGroupApp.RobotApiGroup.RobPairApi
	{
		robPairRouter.POST("createRobPair", robPairApi.CreateRobPair)   // 新建robPair表
		robPairRouter.DELETE("deleteRobPair", robPairApi.DeleteRobPair) // 删除robPair表
		robPairRouter.DELETE("deleteRobPairByIds", robPairApi.DeleteRobPairByIds) // 批量删除robPair表
		robPairRouter.PUT("updateRobPair", robPairApi.UpdateRobPair)    // 更新robPair表
	}
	{
		robPairRouterWithoutRecord.GET("findRobPair", robPairApi.FindRobPair)        // 根据ID获取robPair表
		robPairRouterWithoutRecord.GET("getRobPairList", robPairApi.GetRobPairList)  // 获取robPair表列表
	}
}
