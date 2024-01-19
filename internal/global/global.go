package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"ldacs_sim_sgw/internal/config"
	"sync"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	LOGGER *zap.Logger
	CONFIG config.SgwConfig
	VP     *viper.Viper

	lock sync.RWMutex
)

const (
	UA_LEN  = 28
	SAC_LEN = 12
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
