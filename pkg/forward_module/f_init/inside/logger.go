package inside

import (
	"fmt"
	"ldacs_sim_sgw/internal/global"

	"gorm.io/gorm/logger"
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
	switch global.CONFIG.System.DbType {
	case "mysql":
		logZap = global.CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.LOGGER.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
