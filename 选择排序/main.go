package main

import "fmt"

// 选择排序时间复杂度：n*n=n2
func main() {
	items := []int{3, 4, 11, 9, 5, 20, 9}

	fmt.Println("选择排序", selectSort(items))
}

// 选择排序
func selectSort(items []int) []int {

	var result []int

	for i := 0; i < len(items); i++ {
		smallIndex := findSmall(items)
		result = append(result, items[smallIndex])
		items = append(items[:smallIndex], items[smallIndex+1:]...)
	}

	return result
}

// 找到最小的值
func findSmall(items []int) int {
	smallKeyValue := items[0]
	smallKeyIndex := 0

	for i := 1; i < len(items); i++ {
		if items[i] < smallKeyValue {
			smallKeyValue = items[i]
			smallKeyIndex = i
		}
	}

	return smallKeyIndex
}
