package service

type ServiceGroup struct {
	AccountPlaneService
	AccontFlightService
	AccountAuthzService
	AccountAsService
	AccountGsService
	AccountGscService
	AuthzPlaneService
	AuthcStateService
	AuditAsRawService
}

var (
	AccountPlaneSer  AccountPlaneService
	AccountFlightSer AccontFlightService
	AccountAuthzSer  AccountAuthzService
	AccountAsSer     AccountAsService
	AccountGsSer     AccountGsService
	AccountGscSer    AccountGscService
	AuthzPlaneSer    AuthzPlaneService
	AuthcStateSer    AuthcStateService
	AuditAsRawSer    AuditAsRawService
)
