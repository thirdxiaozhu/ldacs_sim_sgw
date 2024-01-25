import service from '@/utils/request'

// @Tags AccontFlight
// @Summary 创建航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccontFlight true "创建航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /accountFlight/createAccontFlight [post]
export const createAccontFlight = (data) => {
  return service({
    url: '/accountFlight/createAccontFlight',
    method: 'post',
    data
  })
}

// @Tags AccontFlight
// @Summary 删除航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccontFlight true "删除航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountFlight/deleteAccontFlight [delete]
export const deleteAccontFlight = (params) => {
  return service({
    url: '/accountFlight/deleteAccontFlight',
    method: 'delete',
    params
  })
}

// @Tags AccontFlight
// @Summary 批量删除航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountFlight/deleteAccontFlight [delete]
export const deleteAccontFlightByIds = (params) => {
  return service({
    url: '/accountFlight/deleteAccontFlightByIds',
    method: 'delete',
    params
  })
}

// @Tags AccontFlight
// @Summary 更新航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccontFlight true "更新航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountFlight/updateAccontFlight [put]
export const updateAccontFlight = (data) => {
  return service({
    url: '/accountFlight/updateAccontFlight',
    method: 'put',
    data
  })
}

// @Tags AccontFlight
// @Summary 用id查询航班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccontFlight true "用id查询航班"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountFlight/findAccontFlight [get]
export const findAccontFlight = (params) => {
  return service({
    url: '/accountFlight/findAccontFlight',
    method: 'get',
    params
  })
}

// @Tags AccontFlight
// @Summary 分页获取航班列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取航班列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountFlight/getAccontFlightList [get]
export const getAccontFlightList = (params) => {
  return service({
    url: '/accountFlight/getAccontFlightList',
    method: 'get',
    params
  })
}
