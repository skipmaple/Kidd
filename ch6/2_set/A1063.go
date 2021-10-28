package __set

import (
	"fmt"
)

// Set Similarity
// Page 238

func setSimilarity(a, b []int) (res string) {
	// 交集
	sameNum := make([]int, 0, len(a)+len(b))
	// 并集
	diffNum := make([]int, 0, len(a)+len(b))

	for _, ia := range a {
		if contains(b, ia) {
			sameNum = append(sameNum, ia)
		}
	}

	diffNum = append(diffNum, a...)

	for _, ib := range b {
		if !contains(sameNum, ib) && !contains(diffNum, ib) {
			diffNum = append(diffNum, ib)
		}
	}

	//fmt.Println(sameNum, diffNum)

	res = fmt.Sprintf("%.1f", float64(len(sameNum))/float64(len(diffNum))*100) + "%"

	return res
}

func contains(arr []int, e int) bool {
	for _, item := range arr {
		if item == e {
			return true
		}
	}

	return false
}
