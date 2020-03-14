/*
@Time : 2020-03-09 21:10
@Author : mengyueping
@File : shell
@Software: GoLand
*/
package main

import "fmt"

// 希尔排序(插入排序的一种改进版本，一种分组插入方法，缩小增量排序)：
// 1. 取增量整数d1（d1<n,一般取n/2，n是数据总量），分组：下标相差 d1 的分为一组，分为 n/2 组
// 2. 组内排序，排完形成新序列，然后减小增量（一般减半,d1/2），继续步骤1
// 3. 重复上面步骤，直到增量为1
//
// 数据序列有10条数据：
// 49，38，65，97，76，13，27，49，55，04
// 增量为：5，2，1
//
// 平均时间复杂度：O(n log n)
// 最好时间复杂度：O(n log平方 n)
// 最坏时间复杂度：O(n log平方 n)
// 空间复杂度：O(1)

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	ShellSort(s, OrderASC)
	ShellSort(s, OrderDESC)
}

func ShellSort(s []int, f func(s []int, i, add int)) (result []int) {
	result = make([]int, len(s))
	//copy 时是 copy 的 dst 和 src 最小的长度元素（如果 dst len==1,copy 到 src 1 个元素, 0 是 0 个元素。）
	copy(result, s)
	l := len(result)
	add := l / 2
	for add >= 1 {
		for i := 0; i+add < l; i++ {
			f(result, i, add)
		}
		add = add / 2
		fmt.Println(add)
	}
	fmt.Println(result)
	return result
}

//ASC 升序
func OrderASC(s []int, i, add int) {
	if s[i] > s[i+add] {
		s[i], s[i+add] = s[i+add], s[i]
	}
}

//DESC 降序
func OrderDESC(s []int, i, add int) {
	if s[i] < s[i+add] {
		s[i], s[i+add] = s[i+add], s[i]
	}
}
