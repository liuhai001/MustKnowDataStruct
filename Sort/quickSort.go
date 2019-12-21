package main

import "fmt"

//快速排序：分治思想，递归编程技巧
//递推公式：quickSort(p..r) = quickSort(p...q-1) + quickSort(q+1...r)；p>=r 停止

func quickSort(a []int32, n int32) {
	quickSort_go(a, 0, n-1)
}

func quickSort_go(a []int32, p int32, r int32) {
	if p >= r {
		return
	}
	q := partition(a, p, r) //分区函数，分区点
	quickSort_go(a, p, q-1)
	quickSort_go(a, q+1, r)
}

//关键函数,有点技巧，类似于选择排序
func partition(a []int32, p, r int32) int32 {
	pivot := a[r]
	i := p
	for j := i; j < r; j++ {
		if a[j] > pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

func FindMaxK(a []int32, k int32) int32 {
	length := int32(len(a))
	if k > length {
		return 0
	}
	return findMaxK_go(a, 0, length-1, k)

}

//递归
func findMaxK_go(a []int32, p, r, k int32) int32 {
	pivot := partition(a, p, r)
	if k > pivot+1 {
		return findMaxK_go(a, pivot+1, r, k)
	} else if k < pivot+1 {
		return findMaxK_go(a, p, pivot-1, k)
	} else {
		return a[pivot]
	}
	return 0
}

func main() {
	a := []int32{4, 2, 5, 12, 3}
	fmt.Println(FindMaxK(a, 6))

}
