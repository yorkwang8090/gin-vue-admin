package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RobExchangeRouter struct {
}

// InitRobExchangeRouter 初始化 robExchange表 路由信息
func (s *RobExchangeRouter) InitRobExchangeRouter(Router *gin.RouterGroup) {
	robExchangeRouter := Router.Group("robExchange").Use(middleware.OperationRecord())
	robExchangeRouterWithoutRecord := Router.Group("robExchange")
	var robExchangeApi = v1.ApiGroupApp.RobotApiGroup.RobExchangeApi
	{
		robExchangeRouter.POST("createRobExchange", robExchangeApi.CreateRobExchange)   // 新建robExchange表
		robExchangeRouter.DELETE("deleteRobExchange", robExchangeApi.DeleteRobExchange) // 删除robExchange表
		robExchangeRouter.DELETE("deleteRobExchangeByIds", robExchangeApi.DeleteRobExchangeByIds) // 批量删除robExchange表
		robExchangeRouter.PUT("updateRobExchange", robExchangeApi.UpdateRobExchange)    // 更新robExchange表
	}
	{
		robExchangeRouterWithoutRecord.GET("findRobExchange", robExchangeApi.FindRobExchange)        // 根据ID获取robExchange表
		robExchangeRouterWithoutRecord.GET("getRobExchangeList", robExchangeApi.GetRobExchangeList)  // 获取robExchange表列表
	}
}
