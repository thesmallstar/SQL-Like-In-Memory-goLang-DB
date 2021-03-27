package table

import (
	"InMemorySQLDB/Models/database/column"
	"fmt"
)

type Table struct {
	numRows            int            //number of rows in this table
	numColumns         int            //number of columns in this table
	ColumnNamesToindex map[string]int //get on which index any column "X" is stored
	ColumnNames        []string
	columns            []column.Column   //this stores entire data - look column class
	columnTypes        []column.DataType //**not used for this question**, entire thing works without it.
}

func (tab *Table) ConstructTable(columnName []string, columnType []column.DataType) {
	tab.numColumns = len(columnName)
	tab.ColumnNames = make([]string, len(columnName))
	copy(tab.ColumnNames, columnName)
	for index, _ := range columnName {
		tab.ColumnNamesToindex[columnName[index]] = index
		if columnType[index] == column.INT_TYPE {
			intCol := &column.IntColumn{}
			tab.columns = append(tab.columns, intCol)
		} else if columnType[index] == column.STRING_TYPE {
			strCol := &column.StringColumn{}
			tab.columns = append(tab.columns, strCol)
		}
	}
}

func (tab *Table) PrintTable() {
	for k, _ := range tab.ColumnNames {
		fmt.Print(tab.ColumnNames[k] + " ")
	}
	fmt.Println()
	for i := 0; i < tab.numRows; i++ {
		for j := 0; j < tab.numColumns; j++ {
			tab.columns[j].PrintData(i)
			fmt.Print("  ")
		}
		fmt.Println()
	}
}

func (tab *Table) PrintColumn(name string) {
	fmt.Println(name + ":")
	val := tab.ColumnNamesToindex[name]
	for i := 0; i < tab.numRows; i++ {
		tab.columns[val].PrintData(i)
		fmt.Print(" ")
	}
}

func (tab *Table) AddRow(values []string) {
	if len(values) != tab.numColumns {
		fmt.Print("Number of columns you are trying to insert is incorrect!")
		return
	}

	for i := 0; i < tab.numColumns; i++ {
		tab.columns[i].InsertData(values[i])
	}
	tab.numRows++
}
