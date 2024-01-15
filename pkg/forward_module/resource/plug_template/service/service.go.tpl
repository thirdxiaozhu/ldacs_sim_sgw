package service

 {{- if .NeedModel }}
import (
   "ldacs_sim_sgw/pkg/forward_module/plugin/{{ .Snake}}/model"
)
{{ end }}

type {{ .PlugName}}Service struct{}

func (e *{{ .PlugName}}Service) PlugService({{- if .HasRequest }}req model.Request {{ end -}}) ({{- if .HasResponse }}res model.Response,{{ end -}} err error) {
    // 写你的业务逻辑
	return {{- if .HasResponse }} res,{{ end }} nil
}
