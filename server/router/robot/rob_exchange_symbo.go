package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RobExchangeSymboRouter struct {
}

// InitRobExchangeSymboRouter 初始化 robExchangeSymbo表 路由信息
func (s *RobExchangeSymboRouter) InitRobExchangeSymboRouter(Router *gin.RouterGroup) {
	robExchangeSymboRouter := Router.Group("robExchangeSymbo").Use(middleware.OperationRecord())
	robExchangeSymboRouterWithoutRecord := Router.Group("robExchangeSymbo")
	var robExchangeSymboApi = v1.ApiGroupApp.RobotApiGroup.RobExchangeSymboApi
	{
		robExchangeSymboRouter.POST("createRobExchangeSymbo", robExchangeSymboApi.CreateRobExchangeSymbo)   // 新建robExchangeSymbo表
		robExchangeSymboRouter.DELETE("deleteRobExchangeSymbo", robExchangeSymboApi.DeleteRobExchangeSymbo) // 删除robExchangeSymbo表
		robExchangeSymboRouter.DELETE("deleteRobExchangeSymboByIds", robExchangeSymboApi.DeleteRobExchangeSymboByIds) // 批量删除robExchangeSymbo表
		robExchangeSymboRouter.PUT("updateRobExchangeSymbo", robExchangeSymboApi.UpdateRobExchangeSymbo)    // 更新robExchangeSymbo表
	}
	{
		robExchangeSymboRouterWithoutRecord.GET("findRobExchangeSymbo", robExchangeSymboApi.FindRobExchangeSymbo)        // 根据ID获取robExchangeSymbo表
		robExchangeSymboRouterWithoutRecord.GET("getRobExchangeSymboList", robExchangeSymboApi.GetRobExchangeSymboList)  // 获取robExchangeSymbo表列表
	}
}
