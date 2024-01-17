package config

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                   // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	ForwardAddr  int    `mapstructure:"forward-addr" json:"forward-addr" yaml:"forward-addr"`    // 端口值
	BackwardAddr int    `mapstructure:"backward-addr" json:"backward-addr" yaml:"backward-addr"` // 端口值
}
