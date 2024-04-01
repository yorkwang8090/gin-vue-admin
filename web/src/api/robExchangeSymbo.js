import service from '@/utils/request'

// @Tags RobExchangeSymbo
// @Summary 创建robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobExchangeSymbo true "创建robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robExchangeSymbo/createRobExchangeSymbo [post]
export const createRobExchangeSymbo = (data) => {
  return service({
    url: '/robExchangeSymbo/createRobExchangeSymbo',
    method: 'post',
    data
  })
}

// @Tags RobExchangeSymbo
// @Summary 删除robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobExchangeSymbo true "删除robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robExchangeSymbo/deleteRobExchangeSymbo [delete]
export const deleteRobExchangeSymbo = (params) => {
  return service({
    url: '/robExchangeSymbo/deleteRobExchangeSymbo',
    method: 'delete',
    params
  })
}

// @Tags RobExchangeSymbo
// @Summary 批量删除robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robExchangeSymbo/deleteRobExchangeSymbo [delete]
export const deleteRobExchangeSymboByIds = (params) => {
  return service({
    url: '/robExchangeSymbo/deleteRobExchangeSymboByIds',
    method: 'delete',
    params
  })
}

// @Tags RobExchangeSymbo
// @Summary 更新robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobExchangeSymbo true "更新robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robExchangeSymbo/updateRobExchangeSymbo [put]
export const updateRobExchangeSymbo = (data) => {
  return service({
    url: '/robExchangeSymbo/updateRobExchangeSymbo',
    method: 'put',
    data
  })
}

// @Tags RobExchangeSymbo
// @Summary 用id查询robExchangeSymbo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RobExchangeSymbo true "用id查询robExchangeSymbo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robExchangeSymbo/findRobExchangeSymbo [get]
export const findRobExchangeSymbo = (params) => {
  return service({
    url: '/robExchangeSymbo/findRobExchangeSymbo',
    method: 'get',
    params
  })
}

// @Tags RobExchangeSymbo
// @Summary 分页获取robExchangeSymbo表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取robExchangeSymbo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robExchangeSymbo/getRobExchangeSymboList [get]
export const getRobExchangeSymboList = (params) => {
  return service({
    url: '/robExchangeSymbo/getRobExchangeSymboList',
    method: 'get',
    params
  })
}
