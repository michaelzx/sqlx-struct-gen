package tpl

var ModelTmpl = `package {{.PackageName}}
{{ $length := len .Imports }}
{{if ne $length 0}}
import (
    {{range .Imports}}{{.}}
    {{end}}
)
{{end}}

type {{.StructName}} struct {
    {{range .Fields}}{{.}}
    {{end}}
}

func ({{.ShortStructName}} *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
