package dbmeta

import (
	"database/sql"
	"fmt"
	"github.com/jimsmart/schema"
	"strings"
)

type ModelInfo struct {
	PackageName     string
	StructName      string
	ShortStructName string
	TableName       string
	Fields          []string
	Imports         []string
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	// "API":   true,
	// "ASCII": true,
	// "CPU":   true,
	// "CSS":   true,
	// "DNS":   true,
	// "EOF":   true,
	// "GUID":  true,
	// "HTML":  true,
	// "HTTP":  true,
	// "HTTPS": true,
	// "ID":    true,
	// "IP":    true,
	// "JSON":  true,
	// "LHS":   true,
	// "QPS":   true,
	// "RAM":   true,
	// "RHS":   true,
	// "RPC":   true,
	// "SLA":   true,
	// "SMTP":  true,
	// "SSH":   true,
	// "TLS":   true,
	// "TTL":   true,
	// "UI":    true,
	// "UID":   true,
	// "UUID":  true,
	// "URI":   true,
	// "URL":   true,
	// "UTF8":  true,
	// "VM":    true,
	// "XML":   true,
}

var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// GenerateStruct generates a struct for the given table.
func GenerateStruct(db *sql.DB, tableName string, structName string, pkgName string) *ModelInfo {
	cols, _ := schema.Table(db, tableName)
	fields := generateFieldsTypes(db, cols, 0)

	// fields := generateMysqlTypes(db, columnTypes, 0, jsonAnnotation, gormAnnotation, gureguTypes)
	hasTimeField := false
	for _, field := range fields {
		if strings.Contains(field, "time.Time") {
			hasTimeField = true
			break
		}
	}
	hasSqlNullField := false
	for _, field := range fields {
		if strings.Contains(field, "sql.Null") {
			hasSqlNullField = true
			break
		}
	}
	var imports []string
	imports = append(imports, `"github.com/michaelzx/zky-server/core"`)
	if hasSqlNullField {
		imports = append(imports, `"database/sql"`)
	}
	if hasTimeField {
		imports = append(imports, `"time"`)
	}
	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructName:      structName,
		TableName:       tableName,
		ShortStructName: strings.ToLower(string(structName[0])),
		Fields:          fields,
		Imports:         imports,
	}

	return modelInfo
}

// Generate fields string
func generateFieldsTypes(db *sql.DB, columns []*sql.ColumnType, depth int) []string {

	// sort.Strings(keys)

	var fields []string
	var field = ""
	for i, c := range columns {
		nullable, _ := c.Nullable()
		key := c.Name()
		valueType := sqlTypeToGoType(strings.ToLower(c.DatabaseTypeName()), nullable)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := FmtFieldName(stringifyFirstChar(key))

		var annotations []string
		if i == 0 {
			// annotations = append(annotations, fmt.Sprintf("db:\"%s;primary_key\"", key))
			annotations = append(annotations, fmt.Sprintf("db:\"%s\"", key))
		} else {
			annotations = append(annotations, fmt.Sprintf("db:\"%s\"", key))
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			field = fmt.Sprintf("%s %s",
				fieldName,
				valueType)
		}

		fields = append(fields, field)
	}
	return fields
}
