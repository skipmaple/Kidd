package ch4

import (
	"github.com/UncleMaple/Kidd/utils"
)

// 完美数列
// page 165

func perfectList(arr []int, p int) (res int) {

	for i := 1; i < len(arr); i++ {
		temp, j := arr[i], i
		for ; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}

	for i := 0; i < len(arr); i++ {
		temp := upperBound(i, len(arr), arr, arr[i]*p)
		//fmt.Println(i,temp)
		res = utils.Max(res, temp-i)
	}

	return res
}

func upperBound(left, right int, arr []int, min int) int {
	for left < right {
		mid := (left + right) / 2

		if mid >= len(arr) {
			return mid
		}

		if arr[mid] > min {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
