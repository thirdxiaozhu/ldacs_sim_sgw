package inside

import (
	"fmt"

	"gorm.io/gorm/logger"
	"ldacs_sim_sgw/pkg/forward_module/forward_global"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch forward_global.GVA_CONFIG.System.DbType {
	case "mysql":
		logZap = forward_global.GVA_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = forward_global.GVA_CONFIG.Pgsql.LogZap
	}
	if logZap {
		forward_global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
