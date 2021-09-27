package main

import (
	"Kidd/utils"
	"fmt"
	"math"
	"strconv"
)

// 数字加密
// page 71

func numToSecret(a, b string) (res string) {
	numMap := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "J", 11: "Q", 12: "K"}

	flag := len(a) > len(b)

	diff := int(math.Abs(float64(len(a) - len(b))))
	for i := 0; i < diff; i++ {
		if flag {
			b = fmt.Sprintf("%s%s", "0", b)
		} else {
			a = fmt.Sprintf("%s%s", "0", a)
		}
	}

	for i := utils.Max(len(a), len(b)) - 1; i >= 0; i-- {
		aI, _ := strconv.Atoi(string(a[i]))
		bI, _ := strconv.Atoi(string(b[i]))
		if i%2 == 0 {
			res = fmt.Sprintf("%s%s", numMap[(aI+bI)%13], res)
		} else {
			tempDiff := bI - aI
			for tempDiff < 0 {
				tempDiff += 10
			}
			res = fmt.Sprintf("%s%s", numMap[tempDiff], res)
		}
	}

	return res
}
