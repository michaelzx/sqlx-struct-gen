package dbmeta

// Constants for return types of golang
const (
	sqlNullString = "sql.NullString"
	sqlNullInt32  = "sql.NullInt32"
	sqlNullInt64  = "sql.NullInt64"
	sqlNullFloat  = "sql.NullFloat64"
	sqlNullBool   = "sql.NullBool"
	sqlNullTime   = "sql.NullTime"

	golangString = "string"
	golangInt32  = "int32"
	golangInt64  = "int64"
	golangFloat  = "float64"
	golangBool   = "bool"
	golangTime   = "time.Time"

	golangByteArray = "[]byte"
)

func sqlTypeToGoType(mysqlType string, nullable bool) string {
	switch mysqlType {
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		if nullable {
			return sqlNullString
		}
		return golangString
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			return sqlNullInt32
		}
		return golangInt32
	case "bigint":
		if nullable {
			return sqlNullInt64
		}
		return golangInt64
	case "float", "decimal", "double":
		if nullable {
			return sqlNullFloat
		}
		return golangFloat
	case "date", "datetime", "time", "timestamp":
		if nullable {
			return sqlNullTime
		}
		return golangTime
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	case "bit":
		return "bool"
	}

	return ""
}
