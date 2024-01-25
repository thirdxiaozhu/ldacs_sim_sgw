import service from '@/utils/request'

// @Tags AccountAuthz
// @Summary 创建业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountAuthz true "创建业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountAuthz/createAccountAuthz [post]
export const createAccountAuthz = (data) => {
  return service({
    url: '/accountAuthz/createAccountAuthz',
    method: 'post',
    data
  })
}

// @Tags AccountAuthz
// @Summary 删除业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountAuthz true "删除业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountAuthz/deleteAccountAuthz [delete]
export const deleteAccountAuthz = (params) => {
  return service({
    url: '/accountAuthz/deleteAccountAuthz',
    method: 'delete',
    params
  })
}

// @Tags AccountAuthz
// @Summary 批量删除业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountAuthz/deleteAccountAuthz [delete]
export const deleteAccountAuthzByIds = (params) => {
  return service({
    url: '/accountAuthz/deleteAccountAuthzByIds',
    method: 'delete',
    params
  })
}

// @Tags AccountAuthz
// @Summary 更新业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountAuthz true "更新业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountAuthz/updateAccountAuthz [put]
export const updateAccountAuthz = (data) => {
  return service({
    url: '/accountAuthz/updateAccountAuthz',
    method: 'put',
    data
  })
}

// @Tags AccountAuthz
// @Summary 用id查询业务权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccountAuthz true "用id查询业务权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountAuthz/findAccountAuthz [get]
export const findAccountAuthz = (params) => {
  return service({
    url: '/accountAuthz/findAccountAuthz',
    method: 'get',
    params
  })
}

// @Tags AccountAuthz
// @Summary 分页获取业务权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取业务权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountAuthz/getAccountAuthzList [get]
export const getAccountAuthzList = (params) => {
  return service({
    url: '/accountAuthz/getAccountAuthzList',
    method: 'get',
    params
  })
}
