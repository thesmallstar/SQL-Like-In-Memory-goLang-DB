package database

import (
	"InMemorySQLDB/Models/database/column"
	"InMemorySQLDB/Models/database/table"
)

type Database struct {
	name   string
	tables map[string]*table.Table
}

func (db *Database) InitDb(name string) {
	db.tables = make(map[string]*table.Table)
}

func (db *Database) AddNewTable(name string, columnName []string, columnType []column.DataType) {
	xTable := &table.Table{}
	xTable.ColumnNamesToindex = make(map[string]int)
	xTable.ConstructTable(columnName, columnType)
	db.tables[name] = xTable
}

func (db *Database) GetTable(name string) *table.Table {
	return db.tables[name]
}

func (db *Database) DeleteTable(name string) bool {
	delete(db.tables, name)
	if _, ok := db.tables[name]; !ok {
		//table not found
		print("Successfully Deleted Table: " + name)
		return true
	}
	return false
}
