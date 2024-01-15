package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
