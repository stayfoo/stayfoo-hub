/*
@Time : 2020-03-09 17:11
@Author : mengyueping
@File : selection
@Software: GoLand
*/
package main

import "fmt"

// 选择排序：（以最大值为例，最小值类似）
// 1. 获取一组数中的最大值，记录最大值下标 index
// 2. 交换 index 值和第 i 值，i 从0开始。index 和 i 不相等
// 3. 交换成功，再从 i 以后的数中重复第1，2步

// 平均时间复杂度：O(n的平方)
// 最好时间复杂度：O(n的平方)
// 最坏时间复杂度：O(n的平方)
// 优点：交换移动数据次数比较少，性能上略优于冒泡
// 空间复杂度：O(1),只需要一个额外空间用于交换
func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	SelectSort(s, MaxNumIndex)
	SelectSort(s, MinNumIndex)
}

func SelectSort(s []int, f func(s []int) (index, num int)) []int {
	for i := 0; i < len(s); i++ {
		index, _ := f(s[i:])
		if i != i+index {
			s[i], s[i+index] = s[i+index], s[i]
		}
	}
	fmt.Println(s)
	return s
}

// DESC 降序
func MaxNumIndex(s []int) (index, maxNum int) {
	maxNum = s[index]
	for i := index + 1; i < len(s); i++ {
		if s[index] < s[i] {
			index = i
			maxNum = s[i]
		}
	}
	return
}

// ASC 升序
func MinNumIndex(s []int) (index, minNum int) {
	minNum = s[index]
	for i := index + 1; i < len(s); i++ {
		if s[index] > s[i] {
			index = i
			minNum = s[i]
		}
	}
	return
}
