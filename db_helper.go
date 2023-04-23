package seeder

import (
	"database/sql"
	"fmt"
	"strconv"
)

type DBHelper struct {
	db *sql.DB
}

func NewDBHelper(db *sql.DB) *DBHelper {
	return &DBHelper{db}
}

func (dh *DBHelper) Insert(tableName string, columns map[string]any) error {
	var columnStr, placeholderStr string

	for column, placeholder := range columns {
		if columnStr == "" {
			columnStr = column
		} else {
			columnStr += "," + column
		}

		if placeholderStr == "" {
			placeholderStr = convertToStr(placeholder)
		} else {
			placeholderStr += "," + convertToStr(placeholder)
		}
	}

	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", tableName, columnStr, placeholderStr)

	_, err := dh.db.Exec(query)
	return err
}

func convertToStr(target any) string {
	switch target.(type) {
	case string:
		return `"` + target.(string) + `"`
	case int, int8, int16, int32, int64:
		return strconv.Itoa(target.(int))
	default:
		return ""
	}
}
