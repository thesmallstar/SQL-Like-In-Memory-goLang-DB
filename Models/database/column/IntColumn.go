package column

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

//implements Column interface
type IntColumn struct {
	data []int
}

func (c IntColumn) GetData(index int) interface{} {
	return c.data[index]
}

func (c IntColumn) PrintData(index int) {
	fmt.Print(c.data[index])
	return
}

func (c *IntColumn) InsertData(val string) {
	//program exits if validate this doesn't work.
	c.ValidateValue(val)
	fval, _ := strconv.Atoi(val)
	c.data = append(c.data, fval)
}

//Errors are only logged at the moments, can be handled better.
func (c IntColumn) ValidateValue(val string) bool {
	if v, err := strconv.Atoi(val); err == nil && math.Abs(float64(v)) <= 1024 {
		return true
	} else {
		//for now program ends if a invalid input is given..
		print("Invalid value for INT column")
		os.Exit(1)
	}
	return false
}
