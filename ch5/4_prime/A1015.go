package __prime

import (
	"Kidd/utils"
	"strconv"
	"strings"
)

// Reversible Primes
// 反转素数
// page 213

type input struct {
	number int
	radix  int
}

func reversiblePrimes(inputs []input) (res []bool) {
	for _, e := range inputs {
		if isPrime(e.number) {
			numStr := utils.TtoQ(e.number, e.radix)
			numStrReverse := utils.Reverse(strings.Join(numStr, ""))
			numInt, _ := strconv.Atoi(numStrReverse)
			num := utils.QtoT(numInt, e.radix)
			if isPrime(num) {
				res = append(res, true)
			} else {
				res = append(res, false)
			}
		} else {
			res = append(res, false)
		}
	}

	return res
}
