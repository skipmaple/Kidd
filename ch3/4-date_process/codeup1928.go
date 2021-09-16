package main

import (
	"fmt"
	"io"
)

var date1, date2 int
var month = [13][2]int{{0, 0}, {31, 31}, {28, 29}, {31, 31}, {31, 31}, {31, 31}, {31, 31}, {31, 31}, {31, 31}, {31, 31}, {31, 31}, {31, 31}, {31, 31}}

// 输入两个日期 计算差值
func main() {
	res := 1
	for {
		_, err := fmt.Scanf("%d %d", &date1, &date2)
		if err == io.EOF {
			return
		} else {
			if date1 > date2 {
				date1, date2 = date2, date1
			}

			y1 := date1 / 10000
			m1 := date1 % 10000 / 100
			d1 := date1 % 100
			y2, m2, d2 := date2/10000, date2%10000/100, date2%100

			for y1 < y2 || m1 < m2 || d1 < d2 {
				d1++
				if d1 == month[m1][isLeap(y1)] {
					d1 = 1
					m1++
				}

				if m1 == 13 {
					m1 = 1
					y1++
				}

				res++
			}

			fmt.Printf("相隔%d天\n", res)
		}
	}

}

func isLeap(year int) int {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return 1
	} else {
		return 0
	}
}
