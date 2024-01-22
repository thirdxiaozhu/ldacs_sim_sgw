package ldacscore

import (
	"ldacs_sim_sgw/pkg/ldacs_core/handle"
	"ldacs_sim_sgw/pkg/ldacs_core/model"
)

func InitCoreModule() {
	model.RegisterTables()
}
func MakeLdacsHandler() *handle.LdacsHandler {
	return &handle.LdacsHandler{}
}
