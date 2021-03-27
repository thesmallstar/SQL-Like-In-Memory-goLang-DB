package main

import (
	"InMemorySQLDB/Models/database"
	"InMemorySQLDB/Models/database/column"
	"fmt"
)

func startDB() {
	//create Database
	db := &database.Database{}
	db.InitDb("School")

	//Set column names
	colNames := []string{"Name", "RollNumber", "Marks"}
	//set column types (this could have been a map)
	colTypes := []column.DataType{column.STRING_TYPE, column.INT_TYPE, column.INT_TYPE}

	db.AddNewTable("Students", colNames, colTypes)

	val1 := []string{"Manthan", "40", "20"}
	val2 := []string{"Akash", "84", "0"}
	val3 := []string{"Manan", "80", "15"}
	val4 := []string{"Riya", "45", "20"}

	db.GetTable("Students").AddRow(val1)
	db.GetTable("Students").AddRow(val2)
	db.GetTable("Students").AddRow(val3)
	db.GetTable("Students").AddRow(val4)

	fmt.Println("-------Test table Print---------\n")
	db.GetTable("Students").PrintTable()

	fmt.Println()

	fmt.Println("-------Test Column Print---------\n")
	db.GetTable("Students").PrintColumn("Name")

	fmt.Println("\n")
	fmt.Println("-------Test Table Delete---------\n")
	db.DeleteTable("Students")
}
