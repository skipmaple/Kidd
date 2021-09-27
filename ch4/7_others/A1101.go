package __others

import (
	"github.com/UncleMaple/Kidd/utils"
)

// page 186
// 快速排序

func findOriginElement(arr []int) (res []int) {
	leftMax := make([]int, len(arr))
	rightMin := make([]int, len(arr))

	originArr := make([]int, len(arr))
	copy(originArr, arr)

	sortedArr := utils.BubbleSortMin2Max(originArr)
	min, max := sortedArr[0], sortedArr[len(arr)-1]
	leftMax[0], rightMin[len(arr)-1] = min, max

	for i := 1; i < len(arr); i++ {
		leftMax[i] = utils.Max(leftMax[i-1], arr[i])
	}

	for i := len(arr) - 2; i >= 0; i-- {
		rightMin[i] = utils.Min(rightMin[i+1], arr[i])
	}

	//fmt.Println(arr)
	//fmt.Println(leftMax)
	//fmt.Println(rightMin)

	for i := 0; i < len(arr); i++ {
		if arr[i] >= leftMax[i] && arr[i] <= rightMin[i] {
			res = append(res, arr[i])
		}
	}

	return res
}
