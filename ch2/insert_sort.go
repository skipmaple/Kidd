package ch2

// 插入排序
// main page 100

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ { // 进行 n-1 趟排序
		temp, j := arr[i], i // temp 临时存放arr[i], j从i开始往后枚举
		for ; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1] // arr[j-1]后移一位至arr[j]
		}
		arr[j] = temp
	}

	return arr
}
