import service from '@/utils/request'

// @Tags KeyEntity
// @Summary 创建密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KeyEntity true "创建密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /km/createKeyEntity [post]
export const createKeyEntity = (data) => {
  return service({
    url: '/km/createKeyEntity',
    method: 'post',
    data
  })
}

// @Tags KeyEntity
// @Summary 删除密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KeyEntity true "删除密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /km/deleteKeyEntity [delete]
export const deleteKeyEntity = (params) => {
  return service({
    url: '/km/deleteKeyEntity',
    method: 'delete',
    params
  })
}

// @Tags KeyEntity
// @Summary 批量删除密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /km/deleteKeyEntity [delete]
export const deleteKeyEntityByIds = (params) => {
  return service({
    url: '/km/deleteKeyEntityByIds',
    method: 'delete',
    params
  })
}

// @Tags KeyEntity
// @Summary 更新密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KeyEntity true "更新密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /km/updateKeyEntity [put]
export const updateKeyEntity = (data) => {
  return service({
    url: '/km/updateKeyEntity',
    method: 'put',
    data
  })
}

// @Tags KeyEntity
// @Summary 用id查询密钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.KeyEntity true "用id查询密钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /km/findKeyEntity [get]
export const findKeyEntity = (params) => {
  return service({
    url: '/km/findKeyEntity',
    method: 'get',
    params
  })
}

// @Tags KeyEntity
// @Summary 分页获取密钥列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取密钥列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /km/getKeyEntityList [get]
export const getKeyEntityList = (params) => {
  return service({
    url: '/km/getKeyEntityList',
    method: 'get',
    params
  })
}
