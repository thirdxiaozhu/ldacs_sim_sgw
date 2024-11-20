package global

type LdEnum interface {
	GetString() string
	CheckValid() bool
}

type AuthStateKind uint16

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

func (f AuthStateKind) GetString() string {
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
	}[f]
}

func (f AuthStateKind) CheckValid() bool {
	return f <= AUTH_STAGE_OFF
}
