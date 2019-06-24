package main

import "fmt"

//1、O(n)方法，一次遍历
func maxSubArray(nums []int) ([]int, int) {
	maxRes := int(0)
	maxResNums := make([]int, 0)
	tmp := make([]int, 0)
	sum := int(0)
	for _, num := range nums {
		if sum > 0 {
			sum += num
			tmp = append(tmp, num)
		} else {
			tmp = make([]int, 0)
			sum = num
			tmp = append(tmp, num)
		}

		if maxRes < sum {
			maxResNums = make([]int, 0)
			maxRes = sum
			maxResNums = append(maxResNums, tmp...)
		}

	}
	return maxResNums, maxRes
}

//动态规划方法

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSlice, length := maxSubArray(nums)
	fmt.Printf("maxSubArray: %v, %v\n", maxSlice, length)

}
