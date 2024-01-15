package global

{{- if .HasGlobal }}

import "ldacs_sim_sgw/pkg/forward_module/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}