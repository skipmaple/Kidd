package __string

import "fmt"

// Are They Equal
// 用科学记数法比较在n位之内两个数字是否相等
// Page 241

type result struct {
	nStr string
	e    int // 指数的值
}

// areTheyEqual a,b 输入对比的两个数字, n 比较的有效位数
func areTheyEqual(a, b string, n int) (res string) {
	resA := processNum(a, n)
	resB := processNum(b, n)
	if resA == resB {
		res = fmt.Sprintf("YES 0.%s*10^%d", resA.nStr, resA.e)
	} else {
		res = fmt.Sprintf("NO 0.%s*10^%d 0.%s*10^%d", resA.nStr, resA.e, resB.nStr, resB.e)
	}

	return res
}

// processNum n 是有效位数
func processNum(nStr string, n int) (res result) {
	// k  nStr的下标
	k := 0

	// 去除前导零
	for len(nStr) > 0 && nStr[0] == '0' {
		nStr = nStr[1:]
	}

	// 若去掉前导零后是小数点, 则说明s是小于1的小数
	if nStr[0] == '.' {
		nStr = nStr[1:]
		for len(nStr) > 0 && nStr[0] == '0' {
			// 去掉小数点后非零位的所有零
			nStr = nStr[1:]
			// 每去掉一个0, 指数e减1
			res.e--
		}
	} else {
		//	若去掉前导零后不是小数点, 则找到后面的小数点删除
		for k < len(nStr) && nStr[k] != '.' { // 找到小数点
			k++
			// 只要不遇到小数点, 就让指数e++
			res.e++
		}

		//  循环结束后 k < len(nStr), 说明遇到小数点
		if k < len(nStr) {
			// 把小数点删除
			nStr = nStr[:k] + nStr[k+1:]
		}
	}

	if len(nStr) == 0 {
		res.e = 0
	}

	num := 0
	k = 0
	for num < n {
		if k < len(nStr) { // 只要精度还没有到n
			res.nStr += string(nStr[k]) // 只要还有数字, 就加到res末尾
			k++
		} else {
			res.nStr += "0" // 否则res末尾添加0
		}
		num++ // 精度加1
	}

	return res
}
