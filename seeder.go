package seeder

import "database/sql"

type Seeder struct {
	db *sql.DB
}

func NewSeeder(db *sql.DB) *Seeder {
	return &Seeder{db}
}

func (s *Seeder) Table(tableName string) *Table {
	return NewTable(s.db, tableName)
}
