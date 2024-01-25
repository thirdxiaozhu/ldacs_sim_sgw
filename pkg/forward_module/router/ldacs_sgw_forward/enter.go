package ldacs_sgw_forward

type RouterGroup struct {
	AccountPlaneRouter
	AccontFlightRouter
	AccountAuthzRouter
	AccountAsRouter
	AccountGsRouter
	AccountGscRouter
	AuthzPlaneRouter
	AuthcStateRouter
	AuditAsRawRouter
}
