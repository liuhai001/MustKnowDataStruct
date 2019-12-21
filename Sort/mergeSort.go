package main

import "fmt"

//归并排序：是一种分治思想，应用递归编程技巧来实现
//递推公式：mergeSort(p...r) = merge(mergeSort(p..q),mergeSort(q...r)),p>=r 停止

func mergeSort(a []int32, n int32) {
	mergeSort_go(a, 0, n-1)
}

func mergeSort_go(a []int32, p int32, r int32) {
	if p >= r {
		return
	}

	q := (p + r) / 2
	mergeSort_go(a, p, q)
	mergeSort_go(a, q+1, r)
	merge(a, p, q, r)
}

func merge(a []int32, p, q, r int32) {

	i := p
	j := q + 1
	k := 0
	tmp := make([]int32, r-p+1, r-p+1)

	for i <= q && j <= r {
		if a[i] <= a[j] { //加上等于号就是稳定排序
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	begin := i
	end := q
	if j <= r {
		begin = j
		end = r
	}
	for begin <= end {
		tmp[k] = a[begin]
		k++
		begin++
	}

	//复制合并后的数据
	for x := int32(0); x < r-p+1; x++ {
		a[p+x] = tmp[x]
	}

}

func main() {
	a := []int32{5, 4, 65}
	mergeSort(a, int32(len(a)))
	fmt.Println(a)
}
