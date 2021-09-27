package ch2

import "Kidd/utils"

// 归并排序

// 递归写法
// 将arr数组区间[left, right]进行归并排序
func mergeSort(arr []int, left, right int) []int {
	if left < right {
		mid := (left + right) / 2
		arr = mergeSort(arr, left, mid)
		arr = mergeSort(arr, mid+1, right)
		arr = merge(arr, left, mid, mid+1, right)
	}

	return arr
}

func merge(arr []int, l1 int, r1 int, l2 int, r2 int) []int {
	i, j := l1, l2
	tempArr := make([]int, len(arr))
	index := 0
	for i <= r1 && j <= r2 {
		if arr[i] <= arr[j] {
			tempArr[index] = arr[i]
			index++
			i++
		} else {
			tempArr[index] = arr[j]
			index++
			j++
		}
	}

	for i <= r1 {
		tempArr[index] = arr[i]
		index++
		i++
	}

	for j <= r2 {
		tempArr[index] = arr[j]
		index++
		j++
	}

	for i := 0; i < index; i++ {
		arr[l1+i] = tempArr[i]
	}

	return arr
}

// 非递归写法
func mergeSort2(arr []int) []int {
	// step 为组内元素个数，step/2为左子区间元素个数
	for step := 2; step/2 < len(arr); step *= 2 {
		// 每step个元素一组，组内前step/2 和 后step/2 个元素合并
		for i := 0; i < len(arr); i += step {
			mid := i + step/2 - 1
			if mid+1 < len(arr) {
				arr = merge(arr, i, mid, mid+1, utils.Min(i+step-1, len(arr)-1))
			}
		}
	}

	return arr
}
