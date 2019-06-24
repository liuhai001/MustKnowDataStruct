package main

import "fmt"

//把要删除的元素做标记，然后一次性删除排序(注意：此函数只适合正整数，是一种思想)
func removeDuplicates1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = -1 //标记
		}
	}

	newLen := int(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != -1 {
			nums[newLen] = nums[i]
			newLen++
		}
	}

	return newLen

}

//采用两个指针，选择排序的思想
func removeDuplicates2(nums []int) int {
	preIndex := int(0)
	afterIndex := int(1)

	for afterIndex < len(nums) {
		if nums[afterIndex] != nums[preIndex] {
			preIndex++
			nums[preIndex] = nums[afterIndex]
		}
		afterIndex++
	}
	return preIndex + 1

}

func main() {
	nums1 := []int{0, 0, 0, 1, 1, 1, 5, 5, 8, 9, 10, 12, 13, 13}
	fmt.Printf("origin nums1:%v\n", nums1)
	newLen := removeDuplicates2(nums1)
	nums2 := nums1[:newLen]
	fmt.Printf("after remove nums1:%v\n", nums2)
	fmt.Printf("newLen:%v\n", newLen)

}
