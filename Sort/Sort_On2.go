package main

//数组的冒泡排序,升序，时间复杂度：最好O(n),最坏O(n^2),平均O(n^2),稳定排序
func BubbleSort(array []int32, num int32) {
	flag := false
	for i := int32(0); i < num; i++ {
		flag = false
		for j := int32(0); j < num-i-1; j++ {
			if array[j+1] < array[j] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}
		if !flag { //数组已经有序，直接退出
			break
		}
	}
}

//数组的插入排序,升序，时间复杂度：最好O(n),最坏O(n^2),平均O(n^2)，稳定排序
func InsertSort(array []int32, num int32) {
	for i := int32(1); i < num; i++ {
		value := array[i] //保存下要插入的值
		j := i - 1
		for ; j >= 0; j-- { //查找插入的位子
			if array[j] > value {
				array[j+1] = array[j] // 搬移数据
			} else {
				break //找到了插入位子
			}
		}
		array[j+1] = value
	}
}

//数组的选择排序,升序，时间复杂度：最好O(n^2),最坏O(n^2),平均O(n^2)，不稳定排序
func SelectSort(array []int32, num int32) {
	for i := int32(0); i < num; i++ {
		minIndex := i
		for j := i + 1; j < num; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

