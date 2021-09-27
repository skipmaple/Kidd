package __two_pointers

import (
	"Kidd/utils"
)

// 完美数列
// page 176

//func perfectArr(arr []int, p int) (res int) {
//	for i := 0; i < len(arr)-1; i++ {
//		for j := 0; j < len(arr)-i-1; j++ {
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//
//	//fmt.Println(arr)
//	for i, j := 0, 0; i < len(arr) && j < len(arr); {
//		for j < len(arr) && arr[i]*p >= arr[j] {
//			res = utils.Max(res, j-i+1)
//			j++
//		}
//		i++
//	}
//
//	return res
//}

func perfectArr(arr []int, p int) (res int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	l, r := 0, 0
	for l < len(arr) && r < len(arr) {
		for r < len(arr) && arr[r] <= arr[l]*p {
			res = utils.Max(res, r-l+1)
			r++
		}
		l++
	}

	return res
}
