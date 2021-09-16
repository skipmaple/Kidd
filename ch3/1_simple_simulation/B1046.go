package main

import "fmt"

// 划拳
// 输入一个划拳轮数count
// 输入count次a喊 a划 b喊 b划
// 如果一方喊出的数字等于两人划出数字之和，则为获胜，对方喝一杯; 同赢或同输则都不喝
// 分别给出count轮内a和b两人喝酒次数

func main() {
	var count, a1, a2, b1, b2, countA, countB int
	_, _ = fmt.Scanf("%d", &count)
	for count > 0 {
		count--
		_, _ = fmt.Scanf("%d %d %d %d", &a1, &a2, &b1, &b2)
		if a2 == a1+b1 && b2 != a1+b1 {
			// a猜中，b喝一杯
			countB++
		} else if a2 != a1+b1 && b2 == a1+b1 {
			// b猜中, a喝一杯
			countA++
		}
	}
	fmt.Printf("%d %d", countA, countB)
}
