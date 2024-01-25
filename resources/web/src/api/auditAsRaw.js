import service from '@/utils/request'

// @Tags AuditAsRaw
// @Summary 创建AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuditAsRaw true "创建AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /auditAsRaw/createAuditAsRaw [post]
export const createAuditAsRaw = (data) => {
  return service({
    url: '/auditAsRaw/createAuditAsRaw',
    method: 'post',
    data
  })
}

// @Tags AuditAsRaw
// @Summary 删除AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuditAsRaw true "删除AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /auditAsRaw/deleteAuditAsRaw [delete]
export const deleteAuditAsRaw = (params) => {
  return service({
    url: '/auditAsRaw/deleteAuditAsRaw',
    method: 'delete',
    params
  })
}

// @Tags AuditAsRaw
// @Summary 批量删除AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /auditAsRaw/deleteAuditAsRaw [delete]
export const deleteAuditAsRawByIds = (params) => {
  return service({
    url: '/auditAsRaw/deleteAuditAsRawByIds',
    method: 'delete',
    params
  })
}

// @Tags AuditAsRaw
// @Summary 更新AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AuditAsRaw true "更新AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /auditAsRaw/updateAuditAsRaw [put]
export const updateAuditAsRaw = (data) => {
  return service({
    url: '/auditAsRaw/updateAuditAsRaw',
    method: 'put',
    data
  })
}

// @Tags AuditAsRaw
// @Summary 用id查询AS报文
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AuditAsRaw true "用id查询AS报文"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /auditAsRaw/findAuditAsRaw [get]
export const findAuditAsRaw = (params) => {
  return service({
    url: '/auditAsRaw/findAuditAsRaw',
    method: 'get',
    params
  })
}

// @Tags AuditAsRaw
// @Summary 分页获取AS报文列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AS报文列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /auditAsRaw/getAuditAsRawList [get]
export const getAuditAsRawList = (params) => {
  return service({
    url: '/auditAsRaw/getAuditAsRawList',
    method: 'get',
    params
  })
}
