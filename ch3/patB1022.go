package main

import (
	"fmt"
	"io"
)

// D进制的A+B
func main() {
	// 测试QtoTen()
	//ideal := 84
	//res := QtoTen(11, 77)
	//if ideal == res {
	//	fmt.Println("yes")
	//} else {
	//	fmt.Println("no")
	//}

	// 测试TentoQ()
	//fmt.Println(TentoQ(8, 11))

	var a, b, q int
	for {
		_, err := fmt.Scanf("%d %d %d", &a, &b, &q)
		if err == io.EOF {
			break
		}
		r := TentoQ(q, a+b)
		for i := len(r)-1; i>=0; i--{
			fmt.Printf("%d", r[i])
			if i == 0 {
				fmt.Printf("\n")
			}
		}
	}
}

// Q进制转换为十进制
func QtoTen(q int, num int) (res int) {
	product := 1
	for num != 0 {
		res += (num % 10) * product
		num /= 10
		product *= q
	}
	return res
}

// 十进制转换为Q进制
func TentoQ(q int, num int) (res []int) {
	if num == 0 {
		res = append(res, 0)
		return res
	}

	i := 0
	for num != 0 {
		res = append(res, num%q)
		i++
		num /= q
	}
	return res
}
