package ch2

// 选择排序
// main page 99

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ { // 进行 n 趟操作
		k := i
		for j := i; j < len(arr); j++ { // 选出[ i, arr[len(arr)-1] ]中最小的元素，下标为k
			if arr[j] < arr[k] {
				k = j
			}
		}

		arr[i], arr[k] = arr[k], arr[i]
	}

	return arr
}
