package main

// 测试 B008.go

import (
	"fmt"
	//"io"
)

// Gcd 求 a 和 b 的最大公约数

func Gcd(a, b int) (res int) {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}

func arrLoopRight(arr []int, m int) []int {
	n := len(arr)
	if m%n == 0 {
		return arr
	}
	d := Gcd(n, m)
	// temp 临时变量, pos 存放当前处理位置, next 为下一个要处理的位置
	var temp, pos, next int

	for i := n - m; i < n-m+d; i++ {
		temp = arr[i]
		pos = i
		for {
			next = (pos - m + n) % n
			if next != i {
				arr[pos] = arr[next]
			} else {
				arr[pos] = temp
			}
			pos = next
			if pos == i {
				break
			}
		}
	}

	return arr
}

func main() {

	var n, m int
	_, err := fmt.Scanf("%d%d\n", &n, &m)
	if err != nil {
		fmt.Println(err)
	}
	arr := make([]int, n)
	//fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
		//if i == n-1 {
		//	fmt.Scanf("\n")
		//}
	}

	res := arrLoopRight(arr, m)
	//fmt.Printf("%d", res)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d", res[i])
		if i != len(res)-1 {
			fmt.Printf(" ")
		}
	}
}
