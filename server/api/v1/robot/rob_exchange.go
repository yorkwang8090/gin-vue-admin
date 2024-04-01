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

type RobExchangeApi struct {
}

var robExchangeService = service.ServiceGroupApp.RobotServiceGroup.RobExchangeService


// CreateRobExchange 创建robExchange表
// @Tags RobExchange
// @Summary 创建robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobExchange true "创建robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robExchange/createRobExchange [post]
func (robExchangeApi *RobExchangeApi) CreateRobExchange(c *gin.Context) {
	var robExchange robot.RobExchange
	err := c.ShouldBindJSON(&robExchange)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robExchangeService.CreateRobExchange(&robExchange); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRobExchange 删除robExchange表
// @Tags RobExchange
// @Summary 删除robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobExchange true "删除robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robExchange/deleteRobExchange [delete]
func (robExchangeApi *RobExchangeApi) DeleteRobExchange(c *gin.Context) {
	id := c.Query("id")
	if err := robExchangeService.DeleteRobExchange(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRobExchangeByIds 批量删除robExchange表
// @Tags RobExchange
// @Summary 批量删除robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /robExchange/deleteRobExchangeByIds [delete]
func (robExchangeApi *RobExchangeApi) DeleteRobExchangeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := robExchangeService.DeleteRobExchangeByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRobExchange 更新robExchange表
// @Tags RobExchange
// @Summary 更新robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobExchange true "更新robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robExchange/updateRobExchange [put]
func (robExchangeApi *RobExchangeApi) UpdateRobExchange(c *gin.Context) {
	var robExchange robot.RobExchange
	err := c.ShouldBindJSON(&robExchange)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robExchangeService.UpdateRobExchange(robExchange); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRobExchange 用id查询robExchange表
// @Tags RobExchange
// @Summary 用id查询robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobExchange true "用id查询robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robExchange/findRobExchange [get]
func (robExchangeApi *RobExchangeApi) FindRobExchange(c *gin.Context) {
	id := c.Query("id")
	if rerobExchange, err := robExchangeService.GetRobExchange(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerobExchange": rerobExchange}, c)
	}
}

// GetRobExchangeList 分页获取robExchange表列表
// @Tags RobExchange
// @Summary 分页获取robExchange表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robotReq.RobExchangeSearch true "分页获取robExchange表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robExchange/getRobExchangeList [get]
func (robExchangeApi *RobExchangeApi) GetRobExchangeList(c *gin.Context) {
	var pageInfo robotReq.RobExchangeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := robExchangeService.GetRobExchangeInfoList(pageInfo); err != nil {
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
