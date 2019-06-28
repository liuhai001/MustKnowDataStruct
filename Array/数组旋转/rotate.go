package main

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	k = k % length

	if k <= 0 {
		return
	}
	reserve(nums, 0, length-1)
	reserve(nums, 0, k-1)
	reserve(nums, k, length-1)

}
func reserve(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	data := []int{-1, -100, 3, 99}
	rotate(data, 5)
	fmt.Printf("data:%v\n", data)
}
