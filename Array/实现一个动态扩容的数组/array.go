package array

import (
	"errors"
	"fmt"
	"sync"
)

/*
封装的一个动态数组模块
超过容量会自动扩容一半
数组是int32的数据
Author：LH
Date:2019-11-25
*/

type DynamicArray struct {
	totalNum   int32   //总容量
	currentNum int32   //当前数量
	data       []int32 //数据
	sync.Mutex
}

//数组结构体初始化
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		totalNum:   1,
		currentNum: 0,
		data:       make([]int32, 1),
	}
}

func (this *DynamicArray) String() string {
	return fmt.Sprintf("totalNum:%v,currentNum:%v,data:%v", this.totalNum, this.currentNum, this.data)
}

//1、增
//从尾部增加一个元素
func (this *DynamicArray) DataIncr(num int32) {
	this.Lock()
	defer this.Unlock()
	if this.currentNum == this.totalNum { //已满，扩容
		this.totalNum = this.totalNum * 2
		tmpData := make([]int32, this.totalNum)
		copy(tmpData, this.data)
		this.data = tmpData
		this.data[this.currentNum] = num
		this.currentNum++
	} else { //未满，直接添加
		this.data[this.currentNum] = num
		this.currentNum++
	}
}

//在某个位置上插入一个数据
func (this *DynamicArray) InsertData(index int32, num int32) error {
	if index < 0 || index >= this.currentNum {
		return errors.New("index is err!")
	}
	if this.currentNum == this.totalNum { //扩容
		this.totalNum = this.totalNum * 2
		tmpData := make([]int32, this.totalNum)
		for i := int32(0); i < index; i++ {
			tmpData[i] = this.data[i]
		}
		tmpData[index] = num
		for i := index; i < this.currentNum; i++ {
			tmpData[i+1] = this.data[i]
		}
		this.data = tmpData
		this.currentNum++
	} else {
		for i := this.currentNum; i > index; i-- {
			this.data[i] = this.data[i-1]
		}
		this.data[index] = num
		this.currentNum++
	}

	return nil
}

//删
//改
//查
