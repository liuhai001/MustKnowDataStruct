package Sort

import (
	"errors"
)

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

//简单的二分查找实现,存在返回数组中的下标，不存在返回错误
func BinarySearch(array []int32, n int32, value int32) (int32, error) {
	low := int32(0)
	high := n - 1
	for high >= low {
		//mid := (low + high) / 2//可能会溢出
		//mid := low + (high-low)/2
		mid := low + ((high - low) >> 1)
		if array[mid] > value {
			high = mid - 1
		} else if array[mid] < value {
			low = mid + 1
		} else { //找到数据
			return mid, nil
		}
	}
	return -1, errors.New("value is not exist!")
}

//思考题，求一个数的平方根，要求精确到小数点后6位
func BinarySqrt(value float32) float32 {
	if value < 0 {
		return float32(-1)
	}
	low := float32(0)
	high := value
	if value > 0 && value < 1 {
		low = value
		high = float32(1)
	}
	mid := low + (high-low)/2
	for high-low >= 0.000001 {
		if mid*mid > value {
			high = mid
		} else if mid*mid < value {
			low = mid
		} else {
			return mid
		}
		mid = low + (high-low)/2
	}
	return mid
}

//二分查找的变形问题，查找第一次值等于给定值的元素
func BinarySearch_1(array []int32, n int32, value int32) (int32, error) {
	low := int32(0)
	high := n - 1
	for high >= low {
		mid := low + ((high - low) >> 1)
		if array[mid] > value {
			high = mid - 1
		} else if array[mid] < value {
			low = mid + 1
		} else if array[mid] == value {
			if mid == 0 || array[mid-1] != value {
				return mid, nil
			} else {
				high = mid - 1
			}
		}

	}
	return -1, errors.New("value is not exist!")
}

//二分查找的变形问题，查找最后一次值等于给定值的元素
func BinarySearch_2(array []int32, n int32, value int32) (int32, error) {
	low := int32(0)
	high := n - 1
	for high >= low {
		mid := low + ((high - low) >> 1)
		if array[mid] > value {
			high = mid - 1
		} else if array[mid] < value {
			low = mid + 1
		} else if array[mid] == value {
			if mid == n-1 || array[mid+1] != value {
				return mid, nil
			} else {
				low = mid + 1
			}
		}

	}
	return -1, errors.New("value is not exist!")
}

//循环有序数组中二分查找值等于给定值的元素  456123
func BinarySearch_3(array []int32, n int32, value int32) (int32, error) {
	low := int32(0)
	high := n - 1
	for high >= low {
		mid := low + ((high - low) >> 1)
		if array[mid] == value {
			return mid, nil
		}
		if array[mid] >= array[low] { //左半边有序
			if value >= array[low] && value < array[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { //右半边有序
			if value > array[mid] && value <= array[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
			low = mid + 1
		}

	}
	return -1, errors.New("value is not exist!")
}

//
//func main() {
//	array := []int32{4, 5, 6, 1, 2, 3}
//	index, err := BinarySearch_3(array, 6, 1)
//	fmt.Printf("元素下标：%v，err:%v\n", index, err)
//
//}
