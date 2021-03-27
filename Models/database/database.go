package database

import "InMemorySQLDB/Models/database/table"

type Database struct {
	name   string
	tables map[string]*table.Table
}

func (db *Database) InitDb(name string) {
	db.tables = make(map[string]*table.Table)
}

func (db *Database) AddNewTable(name string, columnName []string, columnType []table.DataType) {
	xTable := &table.Table{}
	xTable.ColumnNamesToindex = make(map[string]int)
	xTable.ConstructTable(columnName, columnType)
	db.tables[name] = xTable
}

func (db *Database) GetTable(name string) *table.Table {
	return db.tables[name]
}
