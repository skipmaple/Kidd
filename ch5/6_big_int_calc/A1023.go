package __big_int_calc

import (
	"Kidd/utils"
	"fmt"
)

// Have Fun with Numbers
// Page 225

func haveFunWithNumbers(strN string) (res string) {
	n := utils.String2BigNum(strN)
	nMap := countNum(n)

	doubleN := utils.BigNumMul(n, 2)
	doubleNMap := countNum(doubleN)

	equalFlag := "No"
	if isSameMap(nMap, doubleNMap) {
		equalFlag = "Yes"
	}

	res = fmt.Sprintf("%s %s", equalFlag, utils.BigNum2String(doubleN))

	return res

}

func countNum(arr []int) (res map[int]int) {
	res = make(map[int]int, 10)
	for _, v := range arr {
		res[v]++
	}

	return res
}

func isSameMap(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}

	for ka, va := range a {
		if b[ka] != va {
			return false
		}
	}

	return true
}
