package main

import (
	"math"
	"strconv"
	"strings"
)

// 科学计数法
// page 69

func scienceCount(input string) (res float64) {
	arr := strings.Split(input, "E")
	cardinal, _ := strconv.ParseFloat(arr[0], 64) // 基数
	exp, _ := strconv.Atoi(arr[1])                // 指数值
	if exp < 0 {                                  // 区分正负是因为：基数 * 小数格式的指数值 造成精度不准确
		exp = -exp
		res = cardinal / math.Pow10(exp)
	} else {
		res = cardinal * math.Pow10(exp)
	}

	//fmt.Println("c = ", cardinal)
	//fmt.Println("exp = ", exp)
	//fmt.Println("pow = ", math.Pow10(exp))
	//fmt.Println("should res = ", cardinal / 10)
	return res
}

//func main() {
//	//input := "+1.23400E-03"
//	input := "-1.23400E-03"
//	arr := strings.Split(input, "E")
//	n1,_ := strconv.ParseFloat(arr[0],64)
//	n2,_ := strconv.ParseFloat(arr[1],64)
//	fmt.Println(n1, n2)
//}

//func main() {
//	fmt.Println(scienceCount("+3.1415926E-01"))
//}
