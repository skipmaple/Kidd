package main

import (
	"fmt"
	"math"
)

// 部分A+B

// 输入四个数a, b, da(一位数), db(一位数)；
// pa = a数中包含da个数组成的自然数，同理pb；
// 如 a=112233 da=1 则pa=11，b=2432512 db=2 则pb=222;
// 求pa+pb值
func main() {
	var a, b, da, db, pa, pb int64
	_, _ = fmt.Scanf("%d %d %d %d", &a, &da, &b, &db)
	pa = calc(da, a)
	pb = calc(db, b)
	fmt.Printf("%d", pa+pb)
}

func calc(target, num int64) (res int64) {
	for i:= 0; num > 0; {
		e := num%10
		num /= 10
		//fmt.Printf("calc num: %d, target: %d, element%d: %d\n", num, target, i, e)
		if e == target {
			res += target * int64(math.Pow10(i))
			i++
		}
	}
	//fmt.Println("res: ", res)
	return res
}
