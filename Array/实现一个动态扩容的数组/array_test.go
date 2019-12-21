package array

import (
	"fmt"
	"testing"
)

func TestDynamicArray_DataIncr(t *testing.T) {
	array := NewDynamicArray()
	array.DataIncr(1)
	array.DataIncr(2)
	array.DataIncr(3)
	array.DataIncr(4)
	array.DataIncr(5)
	array.InsertData(1, 6)
	array.InsertData(1, 4)
	array.InsertData(1, 4)
	array.InsertData(1, 4)
	fmt.Printf("%v\n", array.String())
}
