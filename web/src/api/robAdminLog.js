import service from '@/utils/request'

// @Tags RobAdminLog
// @Summary 创建robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobAdminLog true "创建robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robAdminLog/createRobAdminLog [post]
export const createRobAdminLog = (data) => {
  return service({
    url: '/robAdminLog/createRobAdminLog',
    method: 'post',
    data
  })
}

// @Tags RobAdminLog
// @Summary 删除robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobAdminLog true "删除robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robAdminLog/deleteRobAdminLog [delete]
export const deleteRobAdminLog = (params) => {
  return service({
    url: '/robAdminLog/deleteRobAdminLog',
    method: 'delete',
    params
  })
}

// @Tags RobAdminLog
// @Summary 批量删除robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robAdminLog/deleteRobAdminLog [delete]
export const deleteRobAdminLogByIds = (params) => {
  return service({
    url: '/robAdminLog/deleteRobAdminLogByIds',
    method: 'delete',
    params
  })
}

// @Tags RobAdminLog
// @Summary 更新robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobAdminLog true "更新robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robAdminLog/updateRobAdminLog [put]
export const updateRobAdminLog = (data) => {
  return service({
    url: '/robAdminLog/updateRobAdminLog',
    method: 'put',
    data
  })
}

// @Tags RobAdminLog
// @Summary 用id查询robAdminLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RobAdminLog true "用id查询robAdminLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robAdminLog/findRobAdminLog [get]
export const findRobAdminLog = (params) => {
  return service({
    url: '/robAdminLog/findRobAdminLog',
    method: 'get',
    params
  })
}

// @Tags RobAdminLog
// @Summary 分页获取robAdminLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取robAdminLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robAdminLog/getRobAdminLogList [get]
export const getRobAdminLogList = (params) => {
  return service({
    url: '/robAdminLog/getRobAdminLogList',
    method: 'get',
    params
  })
}
