import service from '@/utils/request'

// @Tags AuthcState
// @Summary 创建认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuthcState true "创建认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authcState/createAuthcState [post]
export const createAuthcState = (data) => {
  return service({
    url: '/authcState/createAuthcState',
    method: 'post',
    data
  })
}

// @Tags AuthcState
// @Summary 删除认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuthcState true "删除认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authcState/deleteAuthcState [delete]
export const deleteAuthcState = (params) => {
  return service({
    url: '/authcState/deleteAuthcState',
    method: 'delete',
    params
  })
}

// @Tags AuthcState
// @Summary 批量删除认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authcState/deleteAuthcState [delete]
export const deleteAuthcStateByIds = (params) => {
  return service({
    url: '/authcState/deleteAuthcStateByIds',
    method: 'delete',
    params
  })
}

// @Tags AuthcState
// @Summary 更新认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuthcState true "更新认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authcState/updateAuthcState [put]
export const updateAuthcState = (data) => {
  return service({
    url: '/authcState/updateAuthcState',
    method: 'put',
    data
  })
}

// @Tags AuthcState
// @Summary 用id查询认证状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AuthcState true "用id查询认证状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /authcState/findAuthcState [get]
export const findAuthcState = (params) => {
  return service({
    url: '/authcState/findAuthcState',
    method: 'get',
    params
  })
}

// @Tags AuthcState
// @Summary 分页获取认证状态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取认证状态列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authcState/getAuthcStateList [get]
export const getAuthcStateList = (params) => {
  return service({
    url: '/authcState/getAuthcStateList',
    method: 'get',
    params
  })
}
