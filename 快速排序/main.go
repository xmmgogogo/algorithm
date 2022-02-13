package main

import "fmt"

// 时间复杂度：nlog2n(最差的情况跟选择排序一样n*n=n2)
func main() {
	item := []int{12, 11, 4, 21, 44, 99, 88, 77, 102}
	fmt.Println(quicksort(item))
}

// 快速排序算法
func quicksort(item []int) []int {
	if len(item) <= 1 {
		return item
	}

	key := item[0]
	var left, right []int

	for _, v := range item[1:] {
		if v <= key {
			left = append(left, v)
		}

		if v > key {
			right = append(right, v)
		}
	}

	// left + key + right
	// return quicksort(left) + key + quicksort(right)
	return append(append(quicksort(left), key), quicksort(right)...)
}
