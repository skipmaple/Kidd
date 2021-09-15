package main

import (
	"fmt"
	"strings"
)

// Read Number in Chinese
// page 83

var numMap = map[int]string{0: "零", 1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 7: "七", 8: "八", 9: "九", 10: "十", 100: "百", 1000: "千", 10000: "万", 100000000: "亿", -1: "负"}

func readChineseNum(n int) (res string) {
	if n < 0 {
		res = fmt.Sprintf("%s", numMap[-1])
		n = -n
	}

	yi := n / 100000000
	if yi > 0 {
		res = fmt.Sprintf("%s%s%s", res, numMap[yi], numMap[100000000])
		n %= 100000000
	}

	wan := n / 10000
	if wan > 0 {
		res = fmt.Sprintf("%s%s%s", res, handleNum(wan, yi > 0), numMap[10000])
		n %= 10000
	}

	if n > 0 {
		res = fmt.Sprintf("%s%s", res, handleNum(n, wan%10 == 0))
	}

	return strings.TrimSpace(res)
}

// flag 代表万位有数与否
// ex: 120,000 : flag=true
//     100     : flag=false
func handleNum(num int, flag bool) (res string) {
	qian := num / 1000
	if qian > 0 {
		res = fmt.Sprintf("%s%s", numMap[qian], numMap[1000])
		num %= 1000
	}

	bai := num / 100
	if bai > 0 {
		if qian == 0 && flag {
			res = fmt.Sprintf("%s", numMap[0])
		}
		res = fmt.Sprintf("%s%s%s", res, numMap[bai], numMap[100])
		num %= 100
	}

	shi := num / 10
	if shi > 0 {
		// qian == 0 bai == 0 : +f
		// qian == 0 bai != 0 : -
		// qian != 0 bai == 0 : +
		// qian != 0 bai != 0 : -
		if bai == 0 && (qian != 0 || (qian == 0 && flag)) {
			res = fmt.Sprintf("%s%s", res, numMap[0])
		}
		res = fmt.Sprintf("%s%s%s", res, numMap[shi], numMap[10])
		num %= 10
	}

	ge := num
	if ge > 0 {
		// qian bai shi flag
		// shi == 0
		// qian != 0 || bai != 0 || (qian == 0 && bai == 0 && flag)
		if shi == 0 && (qian != 0 || bai != 0 || (qian == 0 && bai == 0 && flag)) {
			res = fmt.Sprintf("%s%s", res, numMap[0])
		}
		res = fmt.Sprintf("%s%s", res, numMap[ge])
	}

	return res
}
