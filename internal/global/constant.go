package global

type Constant uint32

type processStatus Constant

const (
	PROC_OK processStatus = iota
	PROC_NEW
	PROC_TERM
	PROC_FAIL
	PROC_JSON_FAIL
	PROC_INTERNAL_ERROR
	PROC_NOT_FOUND
)

type linkOrientation Constant

const (
	OriFl linkOrientation = iota
	OriRl
)

const (
	UA_AS_LEN  = 8
	UA_GS_LEN  = 4
	UA_GSC_LEN = 4
)
