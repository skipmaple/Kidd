package __big_int_calc

import (
	"Kidd/utils"
	"fmt"
)

// A除以B
// Page 223

func aDivB(strA string, b int) (res string) {
	a := utils.String2BigNum(strA)
	resStruct := bigNumDiv(a, b)
	res = fmt.Sprintf("%s %d", utils.BigNum2String(resStruct.q), resStruct.r)
	return res
}

type result struct {
	q []int
	r int
}

// bigNumDiv 大整数除法
func bigNumDiv(a []int, b int) (res result) {
	// rest 为余数
	rest := 0
	for i := len(a) - 1; i >= 0; i-- {
		tmp := rest*10 + a[i]
		if tmp < b {
			//	不够除
			res.q = append([]int{0}, res.q...)
			rest = tmp
		} else {
			res.q = append([]int{tmp / b}, res.q...)
			rest = tmp % b
		}
	}

	//去除尾部的0
	for len(res.q)-1 >= 1 && res.q[len(res.q)-1] == 0 {
		res.q = res.q[:len(res.q)-1]
	}

	res.r = rest

	return res
}
