package column

//complicated design, but had to return interface{} to write the code that works for all types
//better design idea? please make an issue :)

type DataType int

//Enum to store datatypes supported
const (
	INT_TYPE DataType = iota
	STRING_TYPE
)

type Column interface {
	GetData(index int) interface{}   //returns the data stored at ith index, or ith row.
	ValidateValue(value string) bool //each column type has it's own validate methods
	InsertData(s string)             //insert data takes string, then validates the data by parsing it
	PrintData(index int)             //prints data at ith index, or ith row
}
