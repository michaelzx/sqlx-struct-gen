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
func ({{.ShortStructName}} *{{.StructName}}) getById(id int64) ({{.StructName}}, error) {
	var row {{.StructName}}
	if err := core.Db.Get(&row, "select * from {{.TableName}} where id = ?", id); err != nil {
		return row, err
	}
	return row, nil
}
`
