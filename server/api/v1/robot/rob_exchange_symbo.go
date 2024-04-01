package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/robot"
    robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type RobExchangeSymboApi struct {
}

var robExchangeSymboService = service.ServiceGroupApp.RobotServiceGroup.RobExchangeSymboService


// CreateRobExchangeSymbo 创建robExchangeSymbo表
// @Tags RobExchangeSymbo
// @Summary 创建robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobExchangeSymbo true "创建robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robExchangeSymbo/createRobExchangeSymbo [post]
func (robExchangeSymboApi *RobExchangeSymboApi) CreateRobExchangeSymbo(c *gin.Context) {
	var robExchangeSymbo robot.RobExchangeSymbo
	err := c.ShouldBindJSON(&robExchangeSymbo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robExchangeSymboService.CreateRobExchangeSymbo(&robExchangeSymbo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRobExchangeSymbo 删除robExchangeSymbo表
// @Tags RobExchangeSymbo
// @Summary 删除robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobExchangeSymbo true "删除robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robExchangeSymbo/deleteRobExchangeSymbo [delete]
func (robExchangeSymboApi *RobExchangeSymboApi) DeleteRobExchangeSymbo(c *gin.Context) {
	id := c.Query("id")
	if err := robExchangeSymboService.DeleteRobExchangeSymbo(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRobExchangeSymboByIds 批量删除robExchangeSymbo表
// @Tags RobExchangeSymbo
// @Summary 批量删除robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /robExchangeSymbo/deleteRobExchangeSymboByIds [delete]
func (robExchangeSymboApi *RobExchangeSymboApi) DeleteRobExchangeSymboByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := robExchangeSymboService.DeleteRobExchangeSymboByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRobExchangeSymbo 更新robExchangeSymbo表
// @Tags RobExchangeSymbo
// @Summary 更新robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobExchangeSymbo true "更新robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robExchangeSymbo/updateRobExchangeSymbo [put]
func (robExchangeSymboApi *RobExchangeSymboApi) UpdateRobExchangeSymbo(c *gin.Context) {
	var robExchangeSymbo robot.RobExchangeSymbo
	err := c.ShouldBindJSON(&robExchangeSymbo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robExchangeSymboService.UpdateRobExchangeSymbo(robExchangeSymbo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRobExchangeSymbo 用id查询robExchangeSymbo表
// @Tags RobExchangeSymbo
// @Summary 用id查询robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobExchangeSymbo true "用id查询robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robExchangeSymbo/findRobExchangeSymbo [get]
func (robExchangeSymboApi *RobExchangeSymboApi) FindRobExchangeSymbo(c *gin.Context) {
	id := c.Query("id")
	if rerobExchangeSymbo, err := robExchangeSymboService.GetRobExchangeSymbo(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerobExchangeSymbo": rerobExchangeSymbo}, c)
	}
}

// GetRobExchangeSymboList 分页获取robExchangeSymbo表列表
// @Tags RobExchangeSymbo
// @Summary 分页获取robExchangeSymbo表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robotReq.RobExchangeSymboSearch true "分页获取robExchangeSymbo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robExchangeSymbo/getRobExchangeSymboList [get]
func (robExchangeSymboApi *RobExchangeSymboApi) GetRobExchangeSymboList(c *gin.Context) {
	var pageInfo robotReq.RobExchangeSymboSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := robExchangeSymboService.GetRobExchangeSymboInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
