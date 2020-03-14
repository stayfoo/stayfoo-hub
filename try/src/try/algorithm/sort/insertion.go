/*
@Time : 2020-03-09 17:56
@Author : mengyueping
@File : insertion
@Software: GoLand
*/
package main

import (
	"fmt"
)

// 插入排序：
// 1. 获取一段有序数据：一个元素组成的切片（s[0]）
// 2. 有序序列之后的元素（第i个：s[1]），比较找到在有序序列中的位置（从i倒序遍历），插入（交换）
// 3. 继续 2 步骤 (第i个：s[2])，直到最后结束
// 平均时间复杂度：O(n平方)
// 最好时间复杂度：O(n)
// 最坏时间复杂度：O(n平方)
// 空间复杂度：O(1),只需要一个额外空间用于交换

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	fmt.Printf("%p\n", s)

	InsertionSort(s)
	fmt.Println(s)
}
func InsertionSort(s []int) []int {
	fmt.Printf("%p\n", s)
	l := len(s)
	if l < 2 {
		return s
	}
	for i := 1; i < l; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			fmt.Println("i:", s[i], " j:", s[j])
			s[j], s[j-1] = s[j-1], s[j]
		}
	}

	fmt.Println(s)
	fmt.Printf("%p\n", s)
	return s
}
