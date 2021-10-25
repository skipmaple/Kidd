package __factors

import (
	"fmt"
	"math"
)

// Consecutive Factors
// 连续质因子
// Page 218

type result struct {
	count  int
	resStr string
}

func consecutiveFactors(n int) (res result) {
	// sqr 为根号n, ansLen为最长连续整数, ansI为对应的第一个整数
	sqr := math.Sqrt(float64(n))
	ansI, ansLen := 0, 0
	for i := 2; i <= int(sqr); i++ {
		// temp 为当前连续整数的乘积
		temp := 1
		j := i
		for {
			temp *= j
			if n%temp != 0 {
				break
			}
			if j-i+1 > ansLen {
				// 更新第一个整数
				ansI = i
				//	更新最长长度
				ansLen = j - i + 1
			}
			j++
		}
	}

	// 最大长度为0, 说明根号n范围内没有解
	if ansLen == 0 {
		res.count = 1
		res.resStr = fmt.Sprintf("%d", n)
	} else {
		res.count = ansLen
		for i := 0; i < ansLen; i++ {
			res.resStr += fmt.Sprintf("%d", ansI+i)
			if i < ansLen-1 {
				res.resStr += " * "
			}
		}
	}

	return res
}
