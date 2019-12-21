package main

//冒泡排序：最好O(n),最坏O(n^2),平均O(n^2)
func BubbleSort1(a []int32, n int32) {
	flag := false
	for i := int32(0); i < n; i++ {
		flag = false
		for j := int32(0); j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

//插入排序：最好O(n),最坏O(n^2),平均O(n^2)
func InsertSort1(a []int32, n int32) {

	for i := int32(1); i < n; i++ {
		tmpValue := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > tmpValue {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = tmpValue
	}

}

//选择排序：最好O(n),最坏O(n^2),平均O(n^2)
func SelectSort1(a []int32, n int32) {
	for i := int32(0); i < n; i++ {
		minIndex := i
		for j := int32(i + 1); j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

func main() {

}
