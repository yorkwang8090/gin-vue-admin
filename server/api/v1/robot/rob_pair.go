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

type RobPairApi struct {
}

var robPairService = service.ServiceGroupApp.RobotServiceGroup.RobPairService


// CreateRobPair 创建robPair表
// @Tags RobPair
// @Summary 创建robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobPair true "创建robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robPair/createRobPair [post]
func (robPairApi *RobPairApi) CreateRobPair(c *gin.Context) {
	var robPair robot.RobPair
	err := c.ShouldBindJSON(&robPair)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robPairService.CreateRobPair(&robPair); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRobPair 删除robPair表
// @Tags RobPair
// @Summary 删除robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobPair true "删除robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robPair/deleteRobPair [delete]
func (robPairApi *RobPairApi) DeleteRobPair(c *gin.Context) {
	id := c.Query("id")
	if err := robPairService.DeleteRobPair(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRobPairByIds 批量删除robPair表
// @Tags RobPair
// @Summary 批量删除robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /robPair/deleteRobPairByIds [delete]
func (robPairApi *RobPairApi) DeleteRobPairByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := robPairService.DeleteRobPairByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRobPair 更新robPair表
// @Tags RobPair
// @Summary 更新robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobPair true "更新robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robPair/updateRobPair [put]
func (robPairApi *RobPairApi) UpdateRobPair(c *gin.Context) {
	var robPair robot.RobPair
	err := c.ShouldBindJSON(&robPair)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robPairService.UpdateRobPair(robPair); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRobPair 用id查询robPair表
// @Tags RobPair
// @Summary 用id查询robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobPair true "用id查询robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robPair/findRobPair [get]
func (robPairApi *RobPairApi) FindRobPair(c *gin.Context) {
	id := c.Query("id")
	if rerobPair, err := robPairService.GetRobPair(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerobPair": rerobPair}, c)
	}
}

// GetRobPairList 分页获取robPair表列表
// @Tags RobPair
// @Summary 分页获取robPair表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robotReq.RobPairSearch true "分页获取robPair表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robPair/getRobPairList [get]
func (robPairApi *RobPairApi) GetRobPairList(c *gin.Context) {
	var pageInfo robotReq.RobPairSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := robPairService.GetRobPairInfoList(pageInfo); err != nil {
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
