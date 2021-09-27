package utils

// 由大到小
func BubbleSortMax2Min(arr []int) (resArr []int) {
	resArr = make([]int, len(arr))
	copy(resArr, arr)
	for i := 0; i < len(resArr)-1; i++ {
		for j := 0; j < len(resArr)-1-i; j++ {
			if resArr[j] < resArr[j+1] {
				resArr[j], resArr[j+1] = resArr[j+1], resArr[j]
			}
		}
	}

	return resArr
}

// 由小到大
func BubbleSortMin2Max(arr []int) (resArr []int) {
	resArr = make([]int, len(arr))
	copy(resArr, arr)
	for i := 0; i < len(resArr)-1; i++ {
		for j := 0; j < len(resArr)-1-i; j++ {
			if resArr[j] > resArr[j+1] {
				resArr[j], resArr[j+1] = resArr[j+1], resArr[j]
			}
		}
	}

	return resArr
}
