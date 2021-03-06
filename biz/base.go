
package biz

import (
	"github.com/xormplus/core"
	"html/template"
	"strings"
)

func Tags3(columnName string) template.HTML {
	return template.HTML("\"" + columnName + "\"")
}
func ExportLabel(columnName string) template.HTML {
	return template.HTML("\"" + columnName + "\"")
}

func Tags1(columnName string) template.HTML {
	return template.HTML("`json:\"" + columnName + "\"`")
}

func Unescaped(x string) interface{} {
	return template.HTML(x)
}

func IsUUID(str string) bool {
	return "uuid" == str
}

func FirstCharLower(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func FirstCharUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func ExportColumn(columnName string) string {
	columnItems := strings.Split(columnName, "_")
	columnItems[0] = FirstCharUpper(columnItems[0])
	for i := 0; i < len(columnItems); i++ {
		item := strings.Title(columnItems[i])

		if strings.ToUpper(item) == "ID" {
			item = "ID"
		}

		columnItems[i] = item
	}

	return strings.Join(columnItems, "")

}

func TypeConvert(st core.SQLType) string {

	switch st.Name {
	case "smallint", "tinyint":
		return "int8"

	case "varchar", "text", "longtext", "char":
		return "string"

	case "date":
		return "string"

	case "int":
		return "int"

	case "timestamp", "datetime":
		return "time.Time"

	case "bigint":
		return "int" // "int64"

	case "float", "double":
		return "float64"

	case "decimal":
		return "decimal.Decimal"

	case "uuid":
		return "gocql.UUID"

	default:
		return st.Name
	}
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func ColumnAndType(table_schema []*core.Column) string {
	result := make([]string, 0, len(table_schema))
	for _, t := range table_schema {
		result = append(result, t.Name+" "+TypeConvert(t.SQLType))
		//t.SQLType.DefaultLength
	}
	return strings.Join(result, ",")
}

func ColumnWithPostfix(columns []string, Postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, t := range columns {
		result = append(result, t+Postfix)
	}
	return strings.Join(result, sep)
}

func MakeQuestionMarkList(num int) string {
	a := strings.Repeat("?,", num)
	return a[:len(a)-1]
}

func HumpStructName(tableName string) string {
	humps := strings.Split(tableName, "_")
	value := ""
	for _, item := range humps {
		value = value + strings.Title(item)
	}

	if value[len(value)-1:] == "s" {
		value = value[:len(value)-1]
	}

	return value
}
