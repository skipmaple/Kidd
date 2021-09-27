package __simple_math

import (
	"Kidd/utils"
	"fmt"
)

func mathBlackHole(input int) (res []string) {
	num := input

	for {
		inputArr := utils.IntToSlice(num, []int{})
		maxArr, minArr := utils.BubbleSortMax2Min(inputArr), utils.BubbleSortMin2Max(inputArr)
		max, min := utils.SliceToInt(maxArr), utils.SliceToInt(minArr)
		num = max - min
		res = append(res, fmt.Sprintf("%04d-%04d=%04d\n", max, min, num))
		if num == 0 || num == 6174 {
			break
		}
	}

	return res
}
