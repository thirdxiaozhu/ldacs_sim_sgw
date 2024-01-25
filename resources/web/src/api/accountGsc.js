import service from '@/utils/request'

// @Tags AccountGsc
// @Summary 创建地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGsc true "创建地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountGsc/createAccountGsc [post]
export const createAccountGsc = (data) => {
  return service({
    url: '/accountGsc/createAccountGsc',
    method: 'post',
    data
  })
}

// @Tags AccountGsc
// @Summary 删除地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGsc true "删除地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountGsc/deleteAccountGsc [delete]
export const deleteAccountGsc = (params) => {
  return service({
    url: '/accountGsc/deleteAccountGsc',
    method: 'delete',
    params
  })
}

// @Tags AccountGsc
// @Summary 批量删除地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountGsc/deleteAccountGsc [delete]
export const deleteAccountGscByIds = (params) => {
  return service({
    url: '/accountGsc/deleteAccountGscByIds',
    method: 'delete',
    params
  })
}

// @Tags AccountGsc
// @Summary 更新地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGsc true "更新地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountGsc/updateAccountGsc [put]
export const updateAccountGsc = (data) => {
  return service({
    url: '/accountGsc/updateAccountGsc',
    method: 'put',
    data
  })
}

// @Tags AccountGsc
// @Summary 用id查询地面控制站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccountGsc true "用id查询地面控制站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountGsc/findAccountGsc [get]
export const findAccountGsc = (params) => {
  return service({
    url: '/accountGsc/findAccountGsc',
    method: 'get',
    params
  })
}

// @Tags AccountGsc
// @Summary 分页获取地面控制站列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取地面控制站列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountGsc/getAccountGscList [get]
export const getAccountGscList = (params) => {
  return service({
    url: '/accountGsc/getAccountGscList',
    method: 'get',
    params
  })
}
