package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"ldacs_sim_sgw/internal/config"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	LOGGER *zap.Logger
	CONFIG config.SgwConfig
	VP     *viper.Viper
)
