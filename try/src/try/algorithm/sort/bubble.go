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

//注意：稳定性，在原序列中，r[i]=r[j]，且r[i]在r[j]之前，而在排序后的序列中，r[i]仍在r[j]之前，则称这种排序算法是稳定的；否则称为不稳定的。
//使用稳定性校验排序正确性

type Order int

const (
	ASC  Order = iota //升序
	DESC              //降序
)

type num struct {
	v    int
	name string
}

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	fmt.Println(reflect.TypeOf(s).Kind().String())

	BubbleSort(s, ASC)
	BubbleSort(s, DESC)

	//ASC
	n := []num{
		{v: 2, name: "index_1-2"},
		{v: 3, name: "index_2-3"},
		{v: 2, name: "index_3-2"},
		{v: 1, name: "index_4-1"},
	}
	for j := len(n); j > 0; j-- {
		for i := 0; i+1 < len(n); i++ {
			if n[i].v > n[i+1].v {
				fmt.Println(n[i].name, n[i+1].name)
				n[i], n[i+1] = n[i+1], n[i]
			}
		}
	}
	fmt.Println(n)

	//DESC
	n2 := []num{
		{v: 2, name: "index_1-2"},
		{v: 3, name: "index_2-3"},
		{v: 2, name: "index_3-2"},
		{v: 1, name: "index_4-1"},
	}

	for i := len(n2); i > 0; i-- {
		for j := len(n2) - 1; j-1 >= 0; j-- {
			if n2[j].v > n2[j-1].v {
				fmt.Println(n2[j], n2[j-1])
				n2[j], n2[j-1] = n2[j-1], n2[j]
			}
		}
	}
	fmt.Println(n2)

}

func BubbleSort(s []int, t Order) []int {

	for i := 0; i < len(s)-1; i++ { //计数器
		if t == ASC {
			for j := 0; j+1 < len(s); j++ {
				if s[j] > s[j+1] {
					s[j], s[j+1] = s[j+1], s[j]
				}
			}
		}
		if t == DESC {
			for j := len(s) - 1; j-1 >= 0; j-- {
				if s[j] > s[j-1] {
					s[j], s[j-1] = s[j-1], s[j]
				}
			}
		}
	}
	fmt.Println("s:", s)
	return s
}
