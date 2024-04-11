package robot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
	robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RobAdminLogApi struct {
}

var robAdminLogService = service.ServiceGroupApp.RobotServiceGroup.RobAdminLogService

// CreateRobAdminLog 创建robAdminLog表
// @Tags RobAdminLog
// @Summary 创建robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobAdminLog true "创建robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robAdminLog/createRobAdminLog [post]
func (robAdminLogApi *RobAdminLogApi) CreateRobAdminLog(c *gin.Context) {
	var robAdminLog robot.RobAdminLog
	// ShouldBindJSON 将JSON请求体绑定到提供的结构体指针上。
	// 参数:
	// - c: 路由上下文，提供绑定功能的环境。
	// - robAdminLog: 指向要绑定JSON数据的结构体的指针。
	// 返回值:
	// - 错误: 如果绑定过程中遇到任何错误，则返回错误信息。
	//
	// 这个函数尝试将HTTP请求中的JSON数据解析并绑定到robAdminLog结构体实例上。
	// 如果解析失败或数据格式不匹配，将返回相应的错误。
	err := c.ShouldBindJSON(&robAdminLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robAdminLogService.CreateRobAdminLog(&robAdminLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRobAdminLog 删除robAdminLog表
// @Tags RobAdminLog
// @Summary 删除robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobAdminLog true "删除robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robAdminLog/deleteRobAdminLog [delete]
func (robAdminLogApi *RobAdminLogApi) DeleteRobAdminLog(c *gin.Context) {
	id := c.Query("id")
	if err := robAdminLogService.DeleteRobAdminLog(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRobAdminLogByIds 批量删除robAdminLog表
// @Tags RobAdminLog
// @Summary 批量删除robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /robAdminLog/deleteRobAdminLogByIds [delete]
func (robAdminLogApi *RobAdminLogApi) DeleteRobAdminLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := robAdminLogService.DeleteRobAdminLogByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRobAdminLog 更新robAdminLog表
// @Tags RobAdminLog
// @Summary 更新robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobAdminLog true "更新robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robAdminLog/updateRobAdminLog [put]
func (robAdminLogApi *RobAdminLogApi) UpdateRobAdminLog(c *gin.Context) {
	var robAdminLog robot.RobAdminLog
	err := c.ShouldBindJSON(&robAdminLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robAdminLogService.UpdateRobAdminLog(robAdminLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRobAdminLog 用id查询robAdminLog表
// @Tags RobAdminLog
// @Summary 用id查询robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobAdminLog true "用id查询robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robAdminLog/findRobAdminLog [get]
func (robAdminLogApi *RobAdminLogApi) FindRobAdminLog(c *gin.Context) {
	id := c.Query("id")
	if rerobAdminLog, err := robAdminLogService.GetRobAdminLog(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerobAdminLog": rerobAdminLog}, c)
	}
}

// GetRobAdminLogList 分页获取robAdminLog表列表
// @Tags RobAdminLog
// @Summary 分页获取robAdminLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robotReq.RobAdminLogSearch true "分页获取robAdminLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robAdminLog/getRobAdminLogList [get]
func (robAdminLogApi *RobAdminLogApi) GetRobAdminLogList(c *gin.Context) {
	var pageInfo robotReq.RobAdminLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := robAdminLogService.GetRobAdminLogInfoList(pageInfo); err != nil {
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
