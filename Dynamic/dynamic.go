package Dynamic

import (
	"fmt"
)

//找出数组中最长递增序列
func Backtrack(input []int) ([]int, int) {
	sliceLen := len(input)
	maxCount := 1
	maxSlice := make([]int, 0, sliceLen)

	for i := 0; i < sliceLen; i++ {
		for j := i + 1; j < sliceLen; j++ {
			m := i
			n := j
			count := 1
			temp := make([]int, 0, sliceLen)
			temp = append(temp, input[m])
			for n < sliceLen {
				if input[m] < input[n] {
					count++
					temp = append(temp, input[n])
					m = n
				}
				n = n + 1
			}

			if count > maxCount {
				maxCount = count
				maxSlice = make([]int, 0, sliceLen)
				maxSlice = append(maxSlice, temp...)

			}
		}
	}

	return maxSlice, maxCount

}

func main() {
	a := []int{1}
	//sort.Ints(a)
	b := a[1:]
	fmt.Printf("%p,%p\n", &a, &b)
	return
}
