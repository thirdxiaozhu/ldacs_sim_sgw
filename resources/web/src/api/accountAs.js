import service from '@/utils/request'

// @Tags AccountAs
// @Summary 创建飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountAs true "创建飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountAs/createAccountAs [post]
export const createAccountAs = (data) => {
  console.log(data)
  return service({
    url: '/accountAs/createAccountAs',
    method: 'post',
    data
  })
}

// @Tags AccountAs
// @Summary 删除飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountAs true "删除飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountAs/deleteAccountAs [delete]
export const deleteAccountAs = (params) => {
  return service({
    url: '/accountAs/deleteAccountAs',
    method: 'delete',
    params
  })
}

export const deprecateAccountAs = (params) => {
  return service({
    url: '/accountAs/deprecateAccountAs',
    method: 'delete',
    params
  })
}

// @Tags AccountAs
// @Summary 批量删除飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountAs/deleteAccountAs [delete]
export const deleteAccountAsByIds = (params) => {
  return service({
    url: '/accountAs/deleteAccountAsByIds',
    method: 'delete',
    params
  })
}

// @Tags AccountAs
// @Summary 更新飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountAs true "更新飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountAs/updateAccountAs [put]
export const updateAccountAs = (data) => {
  return service({
    url: '/accountAs/updateAccountAs',
    method: 'put',
    data
  })
}

// @Tags AccountAs
// @Summary 用id查询飞机站账户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccountAs true "用id查询飞机站账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountAs/findAccountAs [get]
export const findAccountAs = (params) => {
  return service({
    url: '/accountAs/findAccountAs',
    method: 'get',
    params
  })
}

// @Tags AccountAs
// @Summary 分页获取飞机站账户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取飞机站账户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountAs/getAccountAsList [get]
export const getAccountAsList = (params) => {
  return service({
    url: '/accountAs/getAccountAsList',
    method: 'get',
    params
  })
}
export const getOptions = (params) => {
  return service({
    url: '/accountAs/getOptions',
    method: 'get',
    params
  })
}

export const setStateChange = (data) => {
  return service({
    url: '/accountAs/setStateChange',
    method: 'put',
    data
  })
}

export const getAsByIdFlightApi = (params) => {
  return service({
    url: '/accountAs/getAsByIdFlight',
    method: 'get',
    params
  })
}
