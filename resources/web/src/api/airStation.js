import service from '@/utils/request'

// @Tags AirStation
// @Summary 创建飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AirStation true "创建飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /airStation/createAirStation [post]
export const createAirStation = (data) => {
  return service({
    url: '/airStation/createAirStation',
    method: 'post',
    data
  })
}

// @Tags AirStation
// @Summary 删除飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AirStation true "删除飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /airStation/deleteAirStation [delete]
export const deleteAirStation = (params) => {
  return service({
    url: '/airStation/deleteAirStation',
    method: 'delete',
    params
  })
}

// @Tags AirStation
// @Summary 批量删除飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /airStation/deleteAirStation [delete]
export const deleteAirStationByIds = (params) => {
  return service({
    url: '/airStation/deleteAirStationByIds',
    method: 'delete',
    params
  })
}

// @Tags AirStation
// @Summary 更新飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AirStation true "更新飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /airStation/updateAirStation [put]
export const updateAirStation = (data) => {
  return service({
    url: '/airStation/updateAirStation',
    method: 'put',
    data
  })
}

// @Tags AirStation
// @Summary 用id查询飞机站
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AirStation true "用id查询飞机站"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /airStation/findAirStation [get]
export const findAirStation = (params) => {
  return service({
    url: '/airStation/findAirStation',
    method: 'get',
    params
  })
}

// @Tags AirStation
// @Summary 分页获取飞机站列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取飞机站列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /airStation/getAirStationList [get]
export const getAirStationList = (params) => {
  return service({
    url: '/airStation/getAirStationList',
    method: 'get',
    params
  })
}
