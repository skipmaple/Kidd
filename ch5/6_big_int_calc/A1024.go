package __big_int_calc

import (
	"Kidd/utils"
	"fmt"
)

// Palindromic Number
// Page 228

func palindromicNumber(nStr string, step int) (res string) {
	n := utils.String2BigNum(nStr)
	for i := 0; i < step; i++ {
		if isPalindromic(n) {
			return fmt.Sprintf("%s %d", utils.BigNum2String(n), i)
		}

		reverseN := reverseArr(n)
		//fmt.Println(n, reverseN)
		n = utils.BigNumAdd(n, reverseN)
	}

	return fmt.Sprintf("%s %d", utils.BigNum2String(n), step)
}

func isPalindromic(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

func reverseArr(arr []int) (res []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}

	return res
}
