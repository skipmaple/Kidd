package main

import (
	"fmt"
	"math"
)

// A+B and C(64bit)
// page 22

// !! 重点考察溢出问题

func main() {
	var count int
	_, _ = fmt.Scanf("%d", &count)
	res := make([]bool, count)
	for i := 0; i < count; i++ {
		var a, b, c int64
		//_, _ = fmt.Scanf("%v %v %v", &a, &b, &c)
		a = int64(math.Pow(-2, 64))
		b = int64(math.Pow(-2, 64))
		c = -1
		if a > 0 && b > 0 && a+b <= 0 { // ！！注意加 "=" 号，go里值溢出后将值设为0
			res[i] = true
		} else if a < 0 && b < 0 && a+b >= 0 {
			res[i] = false
		} else {
			//fmt.Printf("a+b: %v + %v = %v\n", a, b, a+b)
			//fmt.Printf("c: %v\n", c)
			if a+b > c {
				res[i] = true
			} else {
				res[i] = false
			}
		}
	}

	for i, k := range res {
		fmt.Printf("Case #%d: %t\n", i+1, k)
	}
}
