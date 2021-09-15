package main

import "fmt"

func main() {
	res := callatz(3)
	if res != 5 {
		fmt.Println("wrong!!")
	} else {
		fmt.Println("right")
	}
}

// 卡拉兹猜想
func callatz(n int) int {
	count := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (3*n + 1) / 2
		}
		count++
	}
	return count
}
