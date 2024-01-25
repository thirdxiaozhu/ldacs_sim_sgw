package request

import (
	"ldacs_sim_sgw/pkg/forward_module/model/common/request"
	"ldacs_sim_sgw/pkg/forward_module/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
