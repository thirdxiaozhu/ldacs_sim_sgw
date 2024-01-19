package service

type ServiceGroup struct {
	AccountPlaneService
	AccontFlightService
	AccountAuthzService
	AuthzPlaneService
	AccountAsService
	AuditAsRawService
}
