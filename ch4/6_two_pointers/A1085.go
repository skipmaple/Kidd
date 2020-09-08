package __two_pointers

import (
	"github.com/UncleMaple/Kidd/utils"
)

func perfect_arr(arr []int, p int) (res int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	//fmt.Println(arr)
	for i, j := 0, 0; i < len(arr) && j < len(arr); {
		for j < len(arr) && arr[i]*p >= arr[j] {
			res = utils.Max(res, j-i+1)
			j++
		}
		i++
	}

	return res
}
