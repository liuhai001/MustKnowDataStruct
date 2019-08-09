package main

import (
	"errors"
	"fmt"
)

//实现一个动态扩容的数组几点要求：1、当前的数目；2、数组容量

type ArrayStruct struct {
	totalNums   int32
	currentNums int32
	array       []int32
}

//创建一个数组对象
func NewArray(totalNums int32) *ArrayStruct {
	arrayStruct := &ArrayStruct{
		totalNums:   totalNums,
		currentNums: 0,
		array:       make([]int32, totalNums),
	}
	return arrayStruct
}

//尾部添加一个元素
func (this *ArrayStruct) Add(num int32) {
	//不用扩容
	if this.currentNums < this.totalNums {
		this.array[this.currentNums] = num
		this.currentNums++
	} else { //扩容
		tmpArray := make([]int32, this.totalNums*2)
		for index, num := range this.array {
			tmpArray[index] = num
		}
		tmpArray[this.currentNums] = num
		this.currentNums++
		this.totalNums = this.totalNums * 2
		this.array = tmpArray
	}
}

//获取数组对象的元素个数
func (this *ArrayStruct) Len() int32 {
	return this.currentNums
}

//获取数组对象的容量
func (this *ArrayStruct) Cap() int32 {
	return this.totalNums
}

//插入一个元素
func (this *ArrayStruct) Insert(num int32, index int32) error {
	if index > this.currentNums-1 {
		return errors.New("index is outOfRange!")
	}
	//不用扩容
	if this.currentNums < this.totalNums {
		//往后移出一个空位
		for i := this.currentNums - 1; i >= index; i-- {
			this.array[i+1] = this.array[i]
		}
		this.array[index] = num
		this.currentNums++
	} else { //扩容
		tmpArray := make([]int32, this.totalNums*2)
		for i := int32(0); i < index; i++ {
			tmpArray[i] = this.array[i]
		}
		tmpArray[index] = num

		for i := index; i < this.currentNums; i++ {
			tmpArray[i+1] = this.array[i]
		}

		this.currentNums++
		this.totalNums = this.totalNums * 2
		this.array = tmpArray
	}

	return nil
}

//删除一个元素
func (this *ArrayStruct) Delete(index int32) (int32, error) {
	if index > this.currentNums-1 {
		return 0, errors.New("index is outOfRange!")
	}

	//往前移动填补index的位置
	value := this.array[index]
	for i := index; i < this.currentNums-1; i++ {
		this.array[i] = this.array[i+1]
	}
	this.currentNums--
	return value, nil
}

//获取一个元素
func (this *ArrayStruct) Get(index int32) (int32, error) {
	if index > this.currentNums-1 {
		return 0, errors.New("index is outOfRange!")
	}
	return this.array[index], nil
}

//打印数组对象
func (this *ArrayStruct) Print() {
	fmt.Printf("Cap:%v\n", this.Cap())
	fmt.Printf("Len:%v\n", this.Len())

	var format string
	for i := int32(0); i < this.Len(); i++ {
		format += fmt.Sprintf("%v:%+v|", i, this.array[i])
	}
	fmt.Println(format)

}

func main() {
	array := NewArray(10)
	for i := int32(0); i <= 11; i++ {
		array.Add(i)
	}
	array.Print()
	array.Insert(10,3)
	array.Print()
	array.Insert(10,12)
	array.Print()
	array.Delete(0)
	array.Print()
	array.Delete(12)
	array.Print()


}
