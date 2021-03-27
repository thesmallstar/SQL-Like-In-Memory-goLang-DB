# SQL-likeDB-golang
In memory SQL Like database implemented in golang.

The database has SQL like features, the entire database lives in memory and is cleaned after the program ends.

Features: 
1. Create Tables in Database. 
2. Tables has columns with "types" -> int, strings
3. Add new Rows to table
4. Have Constrains over columns.
5. Print Tables
6. Print any Columnn in a Table
7. Delete tables in database

Sample code to test the code:
```go

        //create Database
	db := &database.Database{}
	db.InitDb("School")

	//Set column names
	colNames := []string{"Name", "RollNumber", "Marks"}
	
	//set column types (this could have been a map)
	//column types are stored as ENUM
	colTypes := []column.DataType{column.STRING_TYPE, column.INT_TYPE, column.INT_TYPE}

        //Add new table to the database
	db.AddNewTable("Students", colNames, colTypes)

	val1 := []string{"Manthan", "40", "20"}
	val2 := []string{"Akash", "84", "0"}
	val3 := []string{"Manan", "80", "15"}
	val4 := []string{"Riya", "45", "20"}

 	//Add new Rows to a table
	db.GetTable("Students").AddRow(val1)
	db.GetTable("Students").AddRow(val2)
	db.GetTable("Students").AddRow(val3)
	db.GetTable("Students").AddRow(val4)

	//Print Entire Table
	db.GetTable("Students").PrintTable()
  
        //Print any column from a table
	db.GetTable("Students").PrintColumn("Name")

  	//Delete any table 
	db.DeleteTable("Students")
  
```





