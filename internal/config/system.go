package config

import (
	"bytes"
	"os/exec"
)

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                   // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	ForwardAddr  int    `mapstructure:"forward-addr" json:"forward-addr" yaml:"forward-addr"`    // 端口值
	BackwardAddr int    `mapstructure:"backward-addr" json:"backward-addr" yaml:"backward-addr"` // 端口值
	SgwUA        uint   `mapstructure:"sgw-ua" json:"sgw-ua" yaml:"sgw-ua"`                      // 端口值
	ConnectMode  string `mapstructure:"connect-mode" json:"connect-mode" yaml:"connect-mode"`    // 端口值
}

func GetLinuxDistroCommand() string {
	cmd := exec.Command("lsb_release", "-i", "-s")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "Unknown Release"
	}
	return out.String()
}
