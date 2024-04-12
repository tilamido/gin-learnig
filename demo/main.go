package main

import (
	"fmt"
	"strconv"
)

// merge 用于合并数组的两个部分
func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// 创建临时切片
	L := make([]int, n1)
	R := make([]int, n2)

	// 数据拷贝到临时切片
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// 合并临时切片到原切片
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// 拷贝L的剩余
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// 拷贝R的剩余
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}

	// 打印每次合并后的数组
	fmt.Printf("Merged %v: %v\n", strconv.Itoa(l)+"-"+strconv.Itoa(r), arr)
}

// mergeSort 归并排序函数，递归排序并合并
func mergeSort(arr []int, l, r int) {
	if l < r {
		// 找到中点
		m := l + (r-l)/2

		// 对前半部分进行归并排序
		mergeSort(arr, l, m)
		// 对后半部分进行归并排序
		mergeSort(arr, m+1, r)

		// 合并两部分
		merge(arr, l, m, r)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Given array is:", arr)

	mergeSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array is:", arr)
}
