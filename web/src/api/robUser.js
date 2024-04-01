import service from '@/utils/request'

// @Tags RobUser
// @Summary 创建user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobUser true "创建user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /robuser/createRobUser [post]
export const createRobUser = (data) => {
  return service({
    url: '/robuser/createRobUser',
    method: 'post',
    data
  })
}

// @Tags RobUser
// @Summary 删除user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobUser true "删除user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robuser/deleteRobUser [delete]
export const deleteRobUser = (params) => {
  return service({
    url: '/robuser/deleteRobUser',
    method: 'delete',
    params
  })
}

// @Tags RobUser
// @Summary 批量删除user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /robuser/deleteRobUser [delete]
export const deleteRobUserByIds = (params) => {
  return service({
    url: '/robuser/deleteRobUserByIds',
    method: 'delete',
    params
  })
}

// @Tags RobUser
// @Summary 更新user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RobUser true "更新user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /robuser/updateRobUser [put]
export const updateRobUser = (data) => {
  return service({
    url: '/robuser/updateRobUser',
    method: 'put',
    data
  })
}

// @Tags RobUser
// @Summary 用id查询user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RobUser true "用id查询user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /robuser/findRobUser [get]
export const findRobUser = (params) => {
  return service({
    url: '/robuser/findRobUser',
    method: 'get',
    params
  })
}

// @Tags RobUser
// @Summary 分页获取user表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取user表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /robuser/getRobUserList [get]
export const getRobUserList = (params) => {
  return service({
    url: '/robuser/getRobUserList',
    method: 'get',
    params
  })
}
