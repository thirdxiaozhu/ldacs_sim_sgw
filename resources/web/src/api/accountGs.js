import service from '@/utils/request'

// @Tags AccountGs
// @Summary 创建地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGs true "创建地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountGs/createAccountGs [post]
export const createAccountGs = (data) => {
  return service({
    url: '/accountGs/createAccountGs',
    method: 'post',
    data
  })
}

// @Tags AccountGs
// @Summary 删除地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGs true "删除地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountGs/deleteAccountGs [delete]
export const deleteAccountGs = (params) => {
  return service({
    url: '/accountGs/deleteAccountGs',
    method: 'delete',
    params
  })
}

// @Tags AccountGs
// @Summary 批量删除地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountGs/deleteAccountGs [delete]
export const deleteAccountGsByIds = (params) => {
  return service({
    url: '/accountGs/deleteAccountGsByIds',
    method: 'delete',
    params
  })
}

// @Tags AccountGs
// @Summary 更新地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGs true "更新地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountGs/updateAccountGs [put]
export const updateAccountGs = (data) => {
  return service({
    url: '/accountGs/updateAccountGs',
    method: 'put',
    data
  })
}

// @Tags AccountGs
// @Summary 用id查询地面站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccountGs true "用id查询地面站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountGs/findAccountGs [get]
export const findAccountGs = (params) => {
  return service({
    url: '/accountGs/findAccountGs',
    method: 'get',
    params
  })
}

// @Tags AccountGs
// @Summary 分页获取地面站列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取地面站列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountGs/getAccountGsList [get]
export const getAccountGsList = (params) => {
  return service({
    url: '/accountGs/getAccountGsList',
    method: 'get',
    params
  })
}
