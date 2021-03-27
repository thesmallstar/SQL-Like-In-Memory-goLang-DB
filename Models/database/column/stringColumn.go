package column

import (
	"fmt"
	"os"
)

//implements Column interface
type StringColumn struct {
	data []string
}

func (c StringColumn) GetData(index int) interface{} {
	return c.data[index]
}

func (c StringColumn) PrintData(index int) {
	fmt.Print(c.data[index])
	return
}

func (c *StringColumn) InsertData(val string) {
	//program exits if validate this doesn't work.
	c.ValidateValue(val)
	c.data = append(c.data, val)
}

//Errors are only logged at the moments, can be handled better.
func (c StringColumn) ValidateValue(val string) bool {
	//this is as per the question I am trying to solve, feel free to ignore the constrain
	if len(val) <= 20 {
		return true
	} else {
		//for now program ends if a invalid input is given..
		print("Invalid value for String column")
		os.Exit(1)
	}
	return false
}
