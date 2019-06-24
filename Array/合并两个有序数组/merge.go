package main

import "fmt"

func merge1(nums1 []int, m int, nums2 []int, n int) {
	//先写一个简单的，需要额外空间O(m+n)
	result := make([]int, 0, m+n)
	num1Index, num2Index := int(0), int(0)

	for num1Index < m && num2Index < n {
		if nums1[num1Index] <= nums2[num2Index] {
			result = append(result, nums1[num1Index])
			num1Index++
		} else {
			result = append(result, nums2[num2Index])
			num2Index++
		}
	}

	if num1Index >= m {
		result = append(result, nums2[num2Index:n]...)
	}

	if num2Index >= n {
		result = append(result, nums1[num1Index:m]...)
	}

	for idx, num := range result {
		nums1[idx] = num
	}

}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	//这个不需要额外空间的，逆向思维，从后往前填写nums1
	totalLen := m + n

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[totalLen-1] = nums1[m-1]
			m--
		} else {
			nums1[totalLen-1] = nums2[n-1]
			n--
		}
		totalLen--
	}

	for n > 0 {
		nums1[totalLen-1] = nums2[n-1]
		n--
		totalLen--
	}

}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}

	merge2(nums1, 1, nums2, 1)
	fmt.Printf("merge:%v\n", nums1)

}
