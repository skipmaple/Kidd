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

// 十进制转换为Q进制
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

// Q进制转换为十进制
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

// give 12345, output []int64{1,2,3,4,5}

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
