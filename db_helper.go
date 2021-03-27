package main

import (
	"InMemorySQLDB/Models/database"
	"InMemorySQLDB/Models/database/table"
)

func startDB() {
	//create Database
	db := &database.Database{}
	db.InitDb("School")

	//Set column names
	colNames := []string{"Name", "RollNumber", "Marks"}
	//set column types (this could have been a map)
	colTypes := []table.DataType{table.STRING_TYPE, table.INT_TYPE, table.INT_TYPE}

	db.AddNewTable("Students", colNames, colTypes)

	val1 := []string{"Manthan", "40", "20"}
	val2 := []string{"Akash", "84", "0"}

	db.GetTable("Students").AddRow(val1)
	db.GetTable("Students").AddRow(val2)
	db.GetTable("Students").PrintTable()

}
