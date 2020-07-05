package main

import "fmt"

// 判断a+b > c，count表示要输入测试用例的个数
func main() {
	var count, num int
	var a, b, c int64
	_, _ = fmt.Scanf("%d", &count)
	for ; count > 0; count-- {
		num++
		_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
		if a+b > c {
			fmt.Printf("Case #%d: true\n", num)
		} else {
			fmt.Printf("Case #%d: false\n", num)
		}
	}
}
