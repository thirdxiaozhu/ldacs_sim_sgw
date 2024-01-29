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

type AuthStateKind Constant

const (
	AUTH_STAGE_UNDEFINED AuthStateKind = iota /* 0 -- most likely accident */
	AUTH_STAGE_DELETING                       /* indicates state is dead but not yet freed */

	/* IKE states */
	AUTH_STAGE_A0
	AUTH_STAGE_G0
	AUTH_STAGE_A1
	AUTH_STAGE_G1
	AUTH_STAGE_A2
	AUTH_STAGE_G2

	AUTH_STAGE_OFF
)

type Bool Constant
const (
	FALSE Bool = iota
	TRUE
)

func (a AuthStateKind) String() string {
	return [...]string{
		"AUTH_STAGE_UNDEFINED",
		"AUTH_STAGE_DELETING",
		"AUTH_STAGE_A0",
		"AUTH_STAGE_G0",
		"AUTH_STAGE_A1",
		"AUTH_STAGE_G1",
		"AUTH_STAGE_A2",
		"AUTH_STAGE_G2",
		"AUTH_STAGE_OFF",
	}[a]
}

const (
	UA_AS_LEN  = 8
	UA_GS_LEN  = 4
	UA_GSC_LEN = 4
)
