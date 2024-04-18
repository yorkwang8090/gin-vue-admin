package robot

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/robot"
	robotReq "github.com/flipped-aurora/gin-vue-admin/server/model/robot/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RobUserApi struct {
}

var robuserService = service.ServiceGroupApp.RobotServiceGroup.RobUserService

var store = captcha.NewDefaultRedisStore()

// CreateRobUser 创建user表
// @Tags RobUser
// @Summary 创建user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobUser true "创建user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robuser/createRobUser [post]
func (robuserApi *RobUserApi) CreateRobUser(c *gin.Context) {
	var robuser robot.RobUser
	err := c.ShouldBindJSON(&robuser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(robuser, utils.RobUserVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//获取当前时间戳单位秒
	temp := uint(time.Now().Unix())
	robuser.Addtime = &temp

	if err := robuserService.CreateRobUser(&robuser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRobUser 删除user表
// @Tags RobUser
// @Summary 删除user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobUser true "删除user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robuser/deleteRobUser [delete]
func (robuserApi *RobUserApi) DeleteRobUser(c *gin.Context) {
	id := c.Query("id")
	if err := robuserService.DeleteRobUser(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRobUserByIds 批量删除user表
// @Tags RobUser
// @Summary 批量删除user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /robuser/deleteRobUserByIds [delete]
func (robuserApi *RobUserApi) DeleteRobUserByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := robuserService.DeleteRobUserByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRobUser 更新user表
// @Tags RobUser
// @Summary 更新user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body robot.RobUser true "更新user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robuser/updateRobUser [put]
func (robuserApi *RobUserApi) UpdateRobUser(c *gin.Context) {
	var robuser robot.RobUser
	err := c.ShouldBindJSON(&robuser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := robuserService.UpdateRobUser(robuser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRobUser 用id查询user表
// @Tags RobUser
// @Summary 用id查询user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobUser true "用id查询user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robuser/findRobUser [get]
func (robuserApi *RobUserApi) FindRobUser(c *gin.Context) {
	id := c.Query("id")
	if rerobuser, err := robuserService.GetRobUser(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerobuser": rerobuser}, c)
	}
}

// GetRobUserList 分页获取user表列表
// @Tags RobUser
// @Summary 分页获取user表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robotReq.RobUserSearch true "分页获取user表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robuser/getRobUserList [get]
func (robuserApi *RobUserApi) GetRobUserList(c *gin.Context) {
	var pageInfo robotReq.RobUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := robuserService.GetRobUserInfoList(pageInfo); err != nil {
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

// login rob登录
// @Tags RobUser
// @Summary 登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query robot.RobUser true "登录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /robuser/login [post]
func (robuserApi *RobUserApi) Login(c *gin.Context) {
	var loginReq robotReq.UserLogin
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(loginReq, utils.UserLoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if loginReq.VerityId != "" && loginReq.Verity != "" && store.Verify(loginReq.VerityId, loginReq.Verity, true) {
		u := &robot.RobUser{UserName: loginReq.UserName, UserPwd: loginReq.Password}
		robuser, err := robuserService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登录失败! 用户名不存在或者密码错误", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if robuser.Status != nil && !*robuser.Status {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		// 生成token，并存入redis
		//token, err := utils.SetRedisJWT(*robuser)
		return
	}
	response.FailWithMessage("验证码错误", c)
}
