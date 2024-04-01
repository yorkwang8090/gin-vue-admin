import service from '@/utils/request'

// @Tags RobCronLog
// @Summary 创建robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobCronLog true "创建robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robCronLog/createRobCronLog [post]
export const createRobCronLog = (data) => {
  return service({
    url: '/robCronLog/createRobCronLog',
    method: 'post',
    data
  })
}

// @Tags RobCronLog
// @Summary 删除robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobCronLog true "删除robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robCronLog/deleteRobCronLog [delete]
export const deleteRobCronLog = (params) => {
  return service({
    url: '/robCronLog/deleteRobCronLog',
    method: 'delete',
    params
  })
}

// @Tags RobCronLog
// @Summary 批量删除robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robCronLog/deleteRobCronLog [delete]
export const deleteRobCronLogByIds = (params) => {
  return service({
    url: '/robCronLog/deleteRobCronLogByIds',
    method: 'delete',
    params
  })
}

// @Tags RobCronLog
// @Summary 更新robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobCronLog true "更新robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robCronLog/updateRobCronLog [put]
export const updateRobCronLog = (data) => {
  return service({
    url: '/robCronLog/updateRobCronLog',
    method: 'put',
    data
  })
}

// @Tags RobCronLog
// @Summary 用id查询robCronLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RobCronLog true "用id查询robCronLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robCronLog/findRobCronLog [get]
export const findRobCronLog = (params) => {
  return service({
    url: '/robCronLog/findRobCronLog',
    method: 'get',
    params
  })
}

// @Tags RobCronLog
// @Summary 分页获取robCronLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取robCronLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robCronLog/getRobCronLogList [get]
export const getRobCronLogList = (params) => {
  return service({
    url: '/robCronLog/getRobCronLogList',
    method: 'get',
    params
  })
}
