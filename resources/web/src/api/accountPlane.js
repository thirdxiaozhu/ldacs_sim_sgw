import service from '@/utils/request'

// @Tags AccountPlane
// @Summary 创建飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPlane true "创建飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountplane/createAccountPlane [post]
export const createAccountPlane = (data) => {
  return service({
    url: '/accountplane/createAccountPlane',
    method: 'post',
    data
  })
}

// @Tags AccountPlane
// @Summary 删除飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPlane true "删除飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountplane/deleteAccountPlane [delete]
export const deleteAccountPlane = (params) => {
  return service({
    url: '/accountplane/deleteAccountPlane',
    method: 'delete',
    params
  })
}

// @Tags AccountPlane
// @Summary 批量删除飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountplane/deleteAccountPlane [delete]
export const deleteAccountPlaneByIds = (params) => {
  return service({
    url: '/accountplane/deleteAccountPlaneByIds',
    method: 'delete',
    params
  })
}

// @Tags AccountPlane
// @Summary 更新飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPlane true "更新飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountplane/updateAccountPlane [put]
export const updateAccountPlane = (data) => {
  return service({
    url: '/accountplane/updateAccountPlane',
    method: 'put',
    data
  })
}

// @Tags AccountPlane
// @Summary 用id查询飞机账户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccountPlane true "用id查询飞机账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountplane/findAccountPlane [get]
export const findAccountPlane = (params) => {
  return service({
    url: '/accountplane/findAccountPlane',
    method: 'get',
    params
  })
}

// @Tags AccountPlane
// @Summary 分页获取飞机账户管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取飞机账户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountplane/getAccountPlaneList [get]
export const getAccountPlaneList = (params) => {
  return service({
    url: '/accountplane/getAccountPlaneList',
    method: 'get',
    params
  })
}
