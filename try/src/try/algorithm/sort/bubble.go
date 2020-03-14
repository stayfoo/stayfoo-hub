/*
@Time : 2020-03-09 16:04
@Author : mengyueping
@File : bubble
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

//冒泡排序：比较相邻的数组元素，满足交换条件交换元素位置，直到 n-1 轮操作结束
//平均时间复杂度：O(n 的平方)
//最好时间复杂度：O(n) 数据中数据有序，遍历一次，不需要交换
//最坏时间复杂度：O(n 的平方)
//空间复杂度：O(1),只需要一个额外空间用于交换
//稳定性：稳定排序

type Order int

const (
	ASC  Order = iota //升序
	DESC              //降序
)

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	fmt.Println(reflect.TypeOf(s).Kind().String())

	BubbleSort(s, ASC)
	BubbleSort(s, DESC)
}

func BubbleSort(s []int, t Order) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if t == ASC {
				if s[i] > s[j] {
					s[i], s[j] = s[j], s[i]
				}
			}
			if t == DESC {
				if s[i] < s[j] {
					s[i], s[j] = s[j], s[i]
				}
			}

		}
	}
	fmt.Println("s:", s)
	return s
}
