package utils

import (
	"math"
	"strconv"
)

type StudentInfo struct {
	Sid     int64
	TestNum int
	RealNum int
	Name    string
	Score   int
}

// Some math methods
//
//fmt.Println(math.Abs(float64(i)))         //取到绝对值
//
//fmt.Println(math.Ceil(3.8))               //向上取整
//
//fmt.Println(math.Floor(3.6))              //向下取整
//
//fmt.Println(math.Mod(11,3))               //取余数 11%3 效果一样
//
//fmt.Println(math.Modf(3.22))              //取整数跟小数
//
//fmt.Println(math.Pow(3,2))             	//X 的 Y次方  乘方
//
//fmt.Println(math.Pow10(3))            	//10的N次方 乘方
//
//fmt.Println(math.Sqrt(9))             	//开平方  3
//
//fmt.Println(math.Cbrt(8))             	//开立方  2
//
//fmt.Println(math.Pi)                     	//π
//
//fmt.Println(math.Round(4.2))          	//四舍五入
//
//fmt.Println(math.IsNaN(3.4))           	//false   报告f是否表示一个NaN（Not A Number）值。
//
//fmt.Println(math.Trunc(1.999999))      	//1    返回整数部分（的浮点值）。
//
//fmt.Println(math.Max(-1.3, 0))     		//0   返回x和y中最大值
//
//fmt.Println(math.Min(-1.3, 0))    		//-1.3  返回x和y中最小值
//
//fmt.Println(math.Dim(-12, -19))   		//7 函数返回x-y和0中的最大值
//
//fmt.Println(math.Dim(-12, 19))    		//0 函数返回x-y和0中的最大值
//
//fmt.Println(math.Cbrt(8))             	//2  返回x的三次方根
//
//fmt.Println(math.Hypot(3, 4))     		//5  返回Sqrt(p*p + q*q)
//
//fmt.Println(math.Pow(2, 8))       		//256  返回x^y

// TtoQ 十进制转换为Q进制
func TtoQ(num, q int) (res []string) {
	if num == 0 {
		return []string{"0"}
	}

	for num != 0 {
		res = append(res, strconv.Itoa(num%q)) //先插入数组的是末位
		num /= q
	}

	// 交换顺序
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// QtoT Q进制转换为十进制
func QtoT(num int, q int) (res int) {
	for i := 0; num != 0; i++ {
		temp := num % 10
		exp := int(math.Pow(float64(q), float64(i)))
		//fmt.Println("i: ", i, "c= ", temp, "exp= ", exp)
		res += temp * exp
		//fmt.Println("res = ", res)
		//res += num % q * int(math.Pow(float64(q), float64(i)))
		num /= 10
	}

	return res
}

// IntToSlice give 12345, output []int64{1,2,3,4,5}
func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func SliceToInt(arr []int) int {
	//s, _ := json.Marshal(arr)
	//fmt.Println(s)
	//fmt.Println(strings.Trim(string(s), "[]"))
	str := ""
	for _, v := range arr {
		str += strconv.Itoa(v)
	}

	res, _ := strconv.Atoi(str)
	return res
}

// Gcd 求 a 和 b 的最大公约数
func Gcd(a, b int) (res int) {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}
