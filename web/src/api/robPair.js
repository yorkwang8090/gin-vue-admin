import service from '@/utils/request'

// @Tags RobPair
// @Summary 创建robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobPair true "创建robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robPair/createRobPair [post]
export const createRobPair = (data) => {
  return service({
    url: '/robPair/createRobPair',
    method: 'post',
    data
  })
}

// @Tags RobPair
// @Summary 删除robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobPair true "删除robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robPair/deleteRobPair [delete]
export const deleteRobPair = (params) => {
  return service({
    url: '/robPair/deleteRobPair',
    method: 'delete',
    params
  })
}

// @Tags RobPair
// @Summary 批量删除robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robPair/deleteRobPair [delete]
export const deleteRobPairByIds = (params) => {
  return service({
    url: '/robPair/deleteRobPairByIds',
    method: 'delete',
    params
  })
}

// @Tags RobPair
// @Summary 更新robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobPair true "更新robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robPair/updateRobPair [put]
export const updateRobPair = (data) => {
  return service({
    url: '/robPair/updateRobPair',
    method: 'put',
    data
  })
}

// @Tags RobPair
// @Summary 用id查询robPair表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RobPair true "用id查询robPair表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robPair/findRobPair [get]
export const findRobPair = (params) => {
  return service({
    url: '/robPair/findRobPair',
    method: 'get',
    params
  })
}

// @Tags RobPair
// @Summary 分页获取robPair表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取robPair表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robPair/getRobPairList [get]
export const getRobPairList = (params) => {
  return service({
    url: '/robPair/getRobPairList',
    method: 'get',
    params
  })
}
