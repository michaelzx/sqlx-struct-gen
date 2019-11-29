package utils


import (
	"github.com/serenize/snaker"
	"os"
	"strings"
	"text/template"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetTemplate(t string) (*template.Template, error) {
	var funcMap = template.FuncMap{
		// "pluralize":        inflection.Plural,
		"title":            strings.Title,
		"toLower":          strings.ToLower,
		"toLowerCamelCase": camelToLowerCamel,
		"toSnakeCase":      snaker.CamelToSnake,
	}
	tmpl, err := template.New("model").Funcs(funcMap).Parse(t)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func camelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])
	return strings.Join(ss, "")
}
