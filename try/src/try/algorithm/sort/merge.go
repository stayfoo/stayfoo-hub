/*
@Time : 2020-03-10 10:26
@Author : mengyueping
@File : merge
@Software: GoLand
*/
package main

import (
	"fmt"
)

// 归并排序：采用分治法（Divide and Conquer）
// 将两个有序表合并成一个有序表，称为二路归并。
// 1. 把一个数据序列分成两个（长度n，每个长度n/2）
// 2. 递归：对每个子序列重复步骤1，直到每个子序列元素最多只有2个，每个子序列元素比较，得到的每个子序列都是有序的
// 3. 递归：子序列两两归并
//
// 平均时间复杂度：O(n log n)
// 最好时间复杂度：O(n log n)
// 最坏时间复杂度：O(n log n)
// 空间复杂度：O(n)
//

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	MergeSort(s)
}

func MergeSort(s []int) []int {
	result := divide(s)
	fmt.Println(result)
	fmt.Println(s)
	return result
}

func divide(s []int) []int {
	l := len(s)
	if l < 2 {
		return s
	}
	index := len(s) / 2

	left := divide(s[:index])
	right := divide(s[index:])
	result := merge(left, right)

	return result
}

func merge(left, right []int) (result []int) {
	leftIndex, rightIndex := 0, 0
	leftLen, rightLen := len(left), len(right)
	for leftIndex < leftLen && rightIndex < rightLen {
		if left[leftIndex] > right[rightIndex] {
			result = append(result, right[rightIndex])
			rightIndex++
			continue
		}
		result = append(result, left[leftIndex])
		leftIndex++
	}

	result = append(result, right[rightIndex:]...)
	result = append(result, left[leftIndex:]...)

	return
}
