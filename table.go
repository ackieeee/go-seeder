package seeder

import "database/sql"

type Table struct {
	dbHelper  *DBHelper
	tableName string
}

func NewTable(db *sql.DB, tableName string) *Table {
	return &Table{
		dbHelper:  NewDBHelper(db),
		tableName: tableName,
	}
}

func (t *Table) Insert(columns map[string]any) error {
	return t.dbHelper.Insert(t.tableName, columns)
}
