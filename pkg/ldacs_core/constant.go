package ldacscore

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

type authStateKind Constant

const (
	AUTH_STATE_UNDEFINED authStateKind = iota /* 0 -- most likely accident */
	AUTH_STATE_DELETING                       /* indicates state is dead but not yet freed */

	/* IKE states */
	AUTH_STATE_A0
	AUTH_STATE_G0
	AUTH_STATE_A1
	AUTH_STATE_G1
	AUTH_STATE_A2
	AUTH_STATE_G2

	AUTH_STATE_OFF
)

func (a authStateKind) String() string {
	return [...]string{
		"AUTH_STATE_UNDEFINED",
		"AUTH_STATE_DELETING",
		"AUTH_STATE_A0",
		"AUTH_STATE_G0",
		"AUTH_STATE_A1",
		"AUTH_STATE_G1",
		"AUTH_STATE_A2",
		"AUTH_STATE_G2",
		"AUTH_STATE_OFF",
	}[a]
}

type snpStateKind Constant

const (
	SNP_STATE_WAIT snpStateKind = iota
	SNP_STATE_CONNECTING
	SNP_STATE_CONNECTED
)

const (
	UA_AS_LEN  = 8
	UA_GS_LEN  = 4
	UA_GSC_LEN = 4
)

const (
	False = 0
	True  = 1
)
