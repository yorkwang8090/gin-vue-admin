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

type RobCronLogApi struct {
}

var robCronLogService = service.ServiceGroupApp.RobotServiceGroup.RobCronLogService


// CreateRobCronLog 创建robCronLog表
// @Tags RobCronLog
// @Summary 创建robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobCronLog true "创建robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robCronLog/createRobCronLog [post]
func (robCronLogApi *RobCronLogApi) CreateRobCronLog(c *gin.Context) {
	var robCronLog robot.RobCronLog
	err := c.ShouldBindJSON(&robCronLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robCronLogService.CreateRobCronLog(&robCronLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRobCronLog 删除robCronLog表
// @Tags RobCronLog
// @Summary 删除robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobCronLog true "删除robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robCronLog/deleteRobCronLog [delete]
func (robCronLogApi *RobCronLogApi) DeleteRobCronLog(c *gin.Context) {
	id := c.Query("id")
	if err := robCronLogService.DeleteRobCronLog(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRobCronLogByIds 批量删除robCronLog表
// @Tags RobCronLog
// @Summary 批量删除robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /robCronLog/deleteRobCronLogByIds [delete]
func (robCronLogApi *RobCronLogApi) DeleteRobCronLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := robCronLogService.DeleteRobCronLogByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRobCronLog 更新robCronLog表
// @Tags RobCronLog
// @Summary 更新robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobCronLog true "更新robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robCronLog/updateRobCronLog [put]
func (robCronLogApi *RobCronLogApi) UpdateRobCronLog(c *gin.Context) {
	var robCronLog robot.RobCronLog
	err := c.ShouldBindJSON(&robCronLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robCronLogService.UpdateRobCronLog(robCronLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRobCronLog 用id查询robCronLog表
// @Tags RobCronLog
// @Summary 用id查询robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobCronLog true "用id查询robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robCronLog/findRobCronLog [get]
func (robCronLogApi *RobCronLogApi) FindRobCronLog(c *gin.Context) {
	id := c.Query("id")
	if rerobCronLog, err := robCronLogService.GetRobCronLog(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerobCronLog": rerobCronLog}, c)
	}
}

// GetRobCronLogList 分页获取robCronLog表列表
// @Tags RobCronLog
// @Summary 分页获取robCronLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robotReq.RobCronLogSearch true "分页获取robCronLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robCronLog/getRobCronLogList [get]
func (robCronLogApi *RobCronLogApi) GetRobCronLogList(c *gin.Context) {
	var pageInfo robotReq.RobCronLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := robCronLogService.GetRobCronLogInfoList(pageInfo); err != nil {
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
