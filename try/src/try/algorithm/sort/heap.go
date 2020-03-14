/*
@Time : 2020-03-10 09:54
@Author : mengyueping
@File : heap
@Software: GoLand
*/
package main

import (
	"fmt"
)

// 堆排序（一种选择排序）：利用堆这种数据结构来设计的一种排序算法。
// 堆：一种近似完全二叉树的结构，子节点的键值和索引总是小于它的父节点（大顶堆，没有要求两个子节点做比较），或者总是大于它的父节点（小顶堆）。
//
// 堆，广度先序遍历，从根节点到子节点，标号从0开始，就可以把这些数据按标号放入数组切片中。
// 这个数组切片从逻辑上讲就是一个堆结构，可以使用简单的公司描述：
// 大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
// 小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]
//
// 时间复杂度：O(n log n)
// 空间复杂度：O(1)
// 步骤：
// 1. 构造大顶堆 (第一个非叶子节点，len(s)/2 - 1；两个子叶子节点没有要求做比较；)
// 2. 取出堆顶元素
// 3. 重复步骤 1，2

type Heap int

const (
	MinHeap Heap = iota
	MaxHeap
)

func main() {
	s := []int{10, 2, 5, 98, 24, 33}
	HeapSort(s, MinHeap)
	HeapSort(s, MaxHeap)
}

func HeapSort(s []int, h Heap) []int {
	tempS := make([]int, len(s))
	copy(tempS, s)
	sortSlice := make([]int, 0)
	for len(tempS) > 1 {
		if h == MinHeap {
			minHeap(tempS)
		}
		if h == MaxHeap {
			maxHeap(tempS)
		}

		sortSlice = append(sortSlice, tempS[0])
		if len(tempS) > 1 {
			tempS = tempS[1:]
		}
		if len(tempS) == 1 {
			sortSlice = append(sortSlice, tempS[0])
		}
	}

	fmt.Println(sortSlice)
	return sortSlice
}

// 构造大顶堆
func maxHeap(s []int) []int {
	//第一个非叶子节点
	index := len(s)/2 - 1
	for index >= 0 {
		subTreeIndex1 := 2*index + 1
		subTreeIndex2 := 2*index + 2
		if subTreeIndex1 < len(s) {
			if s[index] < s[subTreeIndex1] {
				s[index], s[subTreeIndex1] = s[subTreeIndex1], s[index]
			}
		}
		if subTreeIndex2 < len(s) {
			if s[index] < s[subTreeIndex2] {
				s[index], s[subTreeIndex2] = s[subTreeIndex2], s[index]
			}
		}
		index--
	}
	return s
}

// 构造小顶堆
func minHeap(s []int) []int {
	//第一个非叶子节点
	index := len(s)/2 - 1
	for index >= 0 {
		subTreeIndex1 := 2*index + 1
		subTreeIndex2 := 2*index + 2
		if subTreeIndex1 < len(s) {
			if s[index] > s[subTreeIndex1] {
				s[index], s[subTreeIndex1] = s[subTreeIndex1], s[index]
			}
		}
		if subTreeIndex2 < len(s) {
			if s[index] > s[subTreeIndex2] {
				s[index], s[subTreeIndex2] = s[subTreeIndex2], s[index]
			}
		}
		index--
	}
	return s
}
