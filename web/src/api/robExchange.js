import service from '@/utils/request'

// @Tags RobExchange
// @Summary 创建robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobExchange true "创建robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robExchange/createRobExchange [post]
export const createRobExchange = (data) => {
  return service({
    url: '/robExchange/createRobExchange',
    method: 'post',
    data
  })
}

// @Tags RobExchange
// @Summary 删除robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobExchange true "删除robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robExchange/deleteRobExchange [delete]
export const deleteRobExchange = (params) => {
  return service({
    url: '/robExchange/deleteRobExchange',
    method: 'delete',
    params
  })
}

// @Tags RobExchange
// @Summary 批量删除robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robExchange/deleteRobExchange [delete]
export const deleteRobExchangeByIds = (params) => {
  return service({
    url: '/robExchange/deleteRobExchangeByIds',
    method: 'delete',
    params
  })
}

// @Tags RobExchange
// @Summary 更新robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobExchange true "更新robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robExchange/updateRobExchange [put]
export const updateRobExchange = (data) => {
  return service({
    url: '/robExchange/updateRobExchange',
    method: 'put',
    data
  })
}

// @Tags RobExchange
// @Summary 用id查询robExchange表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RobExchange true "用id查询robExchange表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robExchange/findRobExchange [get]
export const findRobExchange = (params) => {
  return service({
    url: '/robExchange/findRobExchange',
    method: 'get',
    params
  })
}

// @Tags RobExchange
// @Summary 分页获取robExchange表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取robExchange表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robExchange/getRobExchangeList [get]
export const getRobExchangeList = (params) => {
  return service({
    url: '/robExchange/getRobExchangeList',
    method: 'get',
    params
  })
}
