package service

type ServiceGroup struct {
	AccountPlaneService
	AccontFlightService
	AccountAuthzService
	AuthzPlaneService
	AccountAsService
	AccountGsService
	AuditAsRawService
}
