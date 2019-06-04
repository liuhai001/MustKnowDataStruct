package main

import (
	"fmt"
	"time"
)

func firstMissingPositive(nums []int) int {
	insertSort(nums) // 先排序
	count := int(1)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			if nums[i] == count {
				count++
			} else {
				return count
			}
		}
	}
	return count

}

//写一个冒泡排序
func bubbleSort(nums []int) {
	length := len(nums)
	flag := false
	for i := 0; i < length-1; i++ { //只需要遍历n-1次
		flag = false
		for j := length - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

//写一个插入排序
func insertSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] < nums[i-1] { //把nums[i]往前插入
			temp := nums[i]
			k := i - 1
			for k >= 0 && nums[k] > temp {
				nums[k+1] = nums[k]
				k--
			}
			nums[k+1] = temp
		}
	}
}

func main() {
	//a := []int{2, 1, 5, 3, 5, 9, 4}
	//a = append(a, 3)
	//p := firstMissingPositive(a)
	pbTime := time.Now().Format("2006-01-02 15:04:05")
	pbDate := time.Now().Format("2006-01-02")
	fmt.Printf("pbTime: %v, pbDate:%v\n", pbTime, pbDate)
}
