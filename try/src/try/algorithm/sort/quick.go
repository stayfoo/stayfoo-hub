/*
@Time : 2020-03-10 11:29
@Author : mengyueping
@File : quick
@Software: GoLand
*/
package main

import "fmt"

// 快速排序：对冒泡排序的改进
// 1. 选取参考值（一般选首个元素，s[0]），数据元素与参考值比较，划分为两组数据（大于参考值，小于参考值）
// 2. 递归：对两组数据分别重复步骤1

// 平均时间复杂度：O(n log n)
// 最好时间复杂度：O(n log n)
// 最坏时间复杂度：O(n平方)
// 空间复杂度：O(log n)

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	QuickSort(s)
}

func QuickSort(s []int) []int {

	index := quickCore(s, 0)

	fmt.Println(s)
	fmt.Println(index)

	return s
}

func quickCore(s []int, index int) int {
	if index == 0 {
		for i := 1; i < len(s); i++ {
			if s[index] > s[i] {
				s[index], s[i] = s[i], s[index]
				index = i
			}
		}
	}

	if index > 0 {
		quickCore(s[:index], 0)
		quickCore(s[index+1:], 0)
	}

	return index
}
