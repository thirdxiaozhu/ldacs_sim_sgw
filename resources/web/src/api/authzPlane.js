import service from '@/utils/request'

// @Tags AuthzPlane
// @Summary 创建飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuthzPlane true "创建飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authzPlane/createAuthzPlane [post]
export const createAuthzPlane = (data) => {
  return service({
    url: '/authzPlane/createAuthzPlane',
    method: 'post',
    data
  })
}


// @Tags AuthzPlane
// @Summary 删除飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuthzPlane true "删除飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authzPlane/deleteAuthzPlane [delete]
export const deleteAuthzPlane = (params) => {
  return service({
    url: '/authzPlane/deleteAuthzPlane',
    method: 'delete',
    params
  })
}

// @Tags AuthzPlane
// @Summary 批量删除飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authzPlane/deleteAuthzPlane [delete]
export const deleteAuthzPlaneByIds = (params) => {
  return service({
    url: '/authzPlane/deleteAuthzPlaneByIds',
    method: 'delete',
    params
  })
}

// @Tags AuthzPlane
// @Summary 更新飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuthzPlane true "更新飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authzPlane/updateAuthzPlane [put]
export const updateAuthzPlane = (data) => {
  return service({
    url: '/authzPlane/updateAuthzPlane',
    method: 'put',
    data
  })
}

// @Tags AuthzPlane
// @Summary 用id查询飞机业务授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AuthzPlane true "用id查询飞机业务授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /authzPlane/findAuthzPlane [get]
export const findAuthzPlane = (params) => {
  return service({
    url: '/authzPlane/findAuthzPlane',
    method: 'get',
    params
  })
}

// @Tags AuthzPlane
// @Summary 分页获取飞机业务授权列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取飞机业务授权列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authzPlane/getAuthzPlaneList [get]
export const getAuthzPlaneList = (params) => {
  return service({
    url: '/authzPlane/getAuthzPlaneList',
    method: 'get',
    params
  })
}

export const getOptions = (params) => {
  return service({
    url: '/authzPlane/getOptions',
    method: 'get',
    params
  })
}

export const setStateChange = (data) => {
  return service({
    url: '/authzPlane/setStateChange',
    method: 'put',
    data
  })
}

