package main

import (
	"fmt"
)

var findNum int

// 二分查询，必须是已经排好序的集合，算法复杂率是log2n
// 查找指定的值
func main() {
	items := []int{}

	for i := 0; i < 100_000_000; i++ {
		items = append(items, i)
	}
	word := 6

	find(items, word)
	fmt.Println("二分查找到的值：", word)
	fmt.Println("二分查找到的次数：", findNum)
}

// 0 1 2 3 4 5 => index_2
// 0 1 2 3 4 => index_1
func find(items []int, key int) int {
	findNum++
	// 这里统计的是索引
	midIndex := int(len(items) / 2)
	// fmt.Println("查询过程：", items, midIndex)

	if key == items[midIndex] {
		return midIndex
	} else if key < items[midIndex] {
		// 大于
		return find(items[:midIndex], key)
	} else {
		// 小于
		return find(items[midIndex+1:], key)
	}
}
